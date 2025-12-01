# Concurrency Control

Consider the given scenario, you've a ticket booking platform which allows multiple users to book
ticket at same time. In one particular case, two users decides to book tickets for same seat at same time.
Ideally the querying logic for this would be to

- Check the availability of seat in DB -> `SELECT is_booked FROM seats WHERE id ='xyz'`
- If available book the seat and update booked to `true` -> `UPDATE seats SET is_booked=1 WHERE id = 'xyz'`

This logic in our case would book the tickets and send the confirmation to both users which shouldn't be happening.
Such scenarios exist all over applications which need to check the inventory and reserve the item for end users.
You could prevent this mishap if you'd control over who could access the selected seat at start of any transaction.
Such features are provided under Concurrency control category and locking is one of the most commonly
used implementations. 

## Locks
Locks are logical constructs which could be placed over a resource to prevents access of certain operations over it
from other parties. Based on access pattern, locks can be divided into 2 major category:

- **Exclusive Locks**: Can be acquired by 1 transaction exclusively allowing it to read-write over the data while 
  preventing read-writes from others. These locks are primarily used to allow a transaction modify data without external
  conflict.  
- **Shared Locks**: Can be shared between multiple transaction for reading but prevents writes from any. This allows
  multiple transaction to read same data concurrently while maintaining the integrity of data. 

Usage of locks should be managed with care as it could easily lead to issues like **Dead Lock** where 2 transactions
are waiting indefinitely to acquire lock obtained by each other. Since neither of transaction will release the 
lock before committing the changes, they’ll both keep waiting forever. Such scenarios must be handled by the
database and depending on implementation rollback and fail one of the transactions. One way to prevent Dead locks
is by using Two Phase Locking.

**Two Phase Locking** ensures transaction are ordered during execution by controlling how they
acquire and release locks. It works in two phases

1. Growing phase where locks are acquired.
2. Shrinking phase where locks are released.

For example, in Double Booking Problem discussed at the start, if we’d acquired the exclusive lock during the
reading phase we could avoid the double booking situation as other transactions can’t read this value currently.
Only after the first transaction commits does this lock gets released upon which it can see that the seat 
is already booked, so it’ll return immediately.

### Locking in Postgres

While the concept of locks on top remains as discussed above, their implementation could vary over different DB systems,
to provide more granular control. For example, we'll discuss the kinds of locks provide by Postgres. For more information,
can view the original postgres docs Reference:
[Postgres Docs](https://www.postgresql.org/docs/current/explicit-locking.html){target=_blank}


Postgres categorizes its locks into 5 categories

#### Tabel Level Locks
Postgres provides 8 different type of table locks and transactions can have multiple locks on same table.
Some of these lock can conflict, others don’t.

- `ACCESS EXCLUSIVE`: conflicts will all other table locks and as such completely locks the table for other transactions.
- `ACCESS SHARE`: Generally acquired by queries which only reads from table like `SELECT`.
- `EXCLUSIVE`: similar to `ACCESS EXCLUSIVE` except it doesn’t conflict `ACCESS SHARE` locks for reading.
  It’s only used by `REFRESH MATERIALIZED VIEW CONCURRENTLY` command, which seems it was added so that users
  can refresh their materialized view while also reading from it.
- `ROW SHARE`: designed for `SELECT FOR...` commands like `SELECT FOR UPDATE`, `SELECT FOR SHARE` which works on
  row level. These commands obtain two kinds of locks — a Row lock and ROW SHARE table lock.
- `ROW EXCLUSIVE`: this lock mainly impacts write latency in system since it's used by write commands like `UPDATE`,
  `DELETE`, `INSERT`, `MERGE`, `COPY FROM`.
- `SHARE ROW EXCLUSIVE`: prevents table against concurrent data changes, and is self-exclusive so that only one session 
  can hold it at a time. Acquired by some `ALTER TABLE` command and `CREATE TRIGGER`.
- `SHARE`: protects a table against concurrent data changes but isn't self exclusive. Used by `CREATE INDEX` so that 
  you can create multiple indexes concurrently but the data isn’t allowed to change.
- `SHARE UPDATE EXCLUSIVE`: allows concurrent writes and reads but prevents schema changes and VACCUM runs.

Below table will summarize which locks would conflict with one another.

![Postgres Table Lock Conflicts](../../static/pg_table_locks_conflict.png)

#### Row Level Locks

Row locks are critical to prevent lost updates. New tuples don’t require locks as they’re only visible to current
transaction in which they were created. That’s why postgres doesn’t support READ UNCOMMITTED isolation. 
These locks are limited to `DELETE` , `UPDATE (no key)` , `UPDATE (key)`, and all `SELECT FOR`s.
(key/no key refers if the column has unique index on that column or not) 

- `FOR UPDATE`: highest row lock, you can’t delete, update on the row when this lock is acquired by other transaction. 
  It’s self conflicting, so you can’t use two `FOR UPDATE` concurrently. You can still read it through normal `SELECT`.
  It’s obtained by `DELETE`, `UPDATE (key)`, `SELECT` commands.
- `FOR NO KEY UPDATE`: acquired for updates to column without a unique index. It’s weaker `FOR UPDATE` as it allows
  `SELECT FOR KEY SHARE`.
- `FOR SHARE`: true shared lock, it can be acquired by multiple transactions. When acquired, it blocks modification to row
- `FOR KEY SHARE`: like `FOR SHARE` but allows update to column without unique index.

To view conflict among above locks, refer to below table
![Postgres Row Lock Conflicts](../../static/pg_row_lock_conflict.png)

!!! note ""
    Postgres stores table locks in memory because they’re coarse but row locks are stored alongside table
    in `xmax` system field, which saves memory but costs disk write. 

#### Page Level Locks

Postgres page are of size 8KB and stores tuples for table and indexes.
Since these pages loaded in shared buffer pool and Postgres being process based backend, multiple process 
can access these pages in shared buffer memory which could lead to inconsistent read or conflicting writes.
To avoid such cases, postgres provides page-level share/exclusive locks to control read/write access to table pages 
in the shared buffer pool.

#### Dead Locks

Using explicit locking can increase the likelihood of deadlocks among transactions in DB.
Postgres detects such conditions and kills one of the transaction to avoid blocking forever.
So long as no deadlock situation is detected, a transaction seeking either a table-level or row-level lock
will wait indefinitely for conflicting locks to be released.
This means it’s a bad idea for applications to hold transactions open for long periods of time

#### Advisory Locks

Sometimes application requirements aren’t satisfied by postgres built in locks. 
To help with this, postgres provides application based locks which are managed by application.
However, these locks still live in databases. These are of two types:
- session lock: obtained with `pg_advisory_lock()` , are kept for the length of session
- transaction lock: obtained with `pg_advisory_xact_lock()`, are kept for length of current running transaction


## Optimistic Concurrency Control

Locking as discussed above is categorized under Pessimistic Concurrency Control, as in you don’t trust others
(Pessimistic) while updating a row, that you lock it and when anyone comes to mess with it while you’re updating 
the row, you can tell them whatever you want to.

Optimistic Concurrency Control in contrast allow transactions to freely (Optimistic) operate during transaction but
validates conflicts during COMMIT. If there's any conflict found, it'll abort the transaction and return
an error to user. This kind of concurrency control is mostly used in read heavy workload where you get
conflicts rarely as data isn't modified as frequently.

## Multi-Version Concurrency Control (MVCC)

Instead of locking rows for reads, the DB keeps multiple versions of rows. Readers see a consistent snapshot
without blocking writers. Postgres uses MVCC which you can read more about in their 
[docs](https://www.postgresql.org/docs/7.1/mvcc.html){target=_blank}