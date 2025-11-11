ACID, which stands for Atomicity, Consistency, Isolation and Durability - are fundamental properties desirable across
all database systems.

To understand these properties individually, you should know about *Transactions*.

??? note annotate "Transactions briefly"
    A transaction is a collection of DML(1) queries treated as one unit of work at application logic. 
    
    For Example, money transfer between accounts requires multiple operations (check balance, debit one account, credit another)
    to succeed or fail together. However, this is performed using a bunch DML queries and failure in any single of them
    could result in bugs like money being debited even if there's no sufficient balance, money being debited without
    crediting, etc. That's is why all these operations should be wrapped in an **Transaction** which ensures they
    either run all or none.
    
    Transaction lifecycle involves keywords to start (`BEGIN`) , save changes (`COMMIT`) and
    discard changes (`ROLLBACK`). Each of these commands are implemented differently across different DBMS,
    like COMMIT either flushes changes made in memory to disk in one go, or it saves individual changes separately.
    This is due to the tradeoffs involved with such decision, which makes each DB unique and optimized for their 
    own specific use case and no general DB system can handle it all. 
    
    Often transactions are mostly used in writing data, but you can also have read-only transactions. For example,
    transactions for generating consistent reports by providing a time-based snapshot of data.

1. Data Manipulation Language

### Atomicity

Every transaction should be treated as indivisible unit of work (either all queries within it succeed or fail, no partials). 

This helps DB to remain in a consistent state by guaranteeing that all failed transactions are rolled back which helps
 to prevent data corruption and maintain consistency. It also simplifies error handling where developers don't need to 
handle rollbacks.

There are different ways to implement atomicity, few of which are:

- Logging: before writing changes to disk, they're written to undo/redo logs and only applied when commit is successful.
- Shadow Copy: changes are applied to a copy of original page. When the transaction is successful, 
  the pointers to data are updated to apply the changes.
- Two-Phase Commits: used in distributed systems, ensuring all peers commit or abort the transaction together.

### Isolation

Isolation property in transaction helps prevent concurrent operations from interfering with each other,
ensuring each transaction appears to run on its own.
It's managed through different isolation levels, which control how transactions interact with the concurrency
anomalies encountered. 

Concurrency anomalies (or Read Phenomenas) are undesirable side effects of running multiple transactions at same 
time. Few of these includes:

- **Dirty Reads**: when a transaction reads uncommitted changes made by another concurrent transaction, 
  which is rolled back. From DB point of view this change was never present in data as it was never committed 
  essentially making our read dirty.  
- **Non-Repeatable Reads**: when you read same entry more than once in a transaction, and it yields different values.
   For example, you read a value directly in first query and then collect sum in second query. 
  If the value is changed when collecting sum, this will result in inconsistent sum w.r.t data in first query.
  That’s why it's called non-repeatable as in you can't read repeated value in same transaction. 
- **Phantom Reads**: when re-reading a range of rows, a new row appears due to write by other transaction. 
  The reason it's different from Non-Repeatable Reads is due to the way **Repeatable Read** (1) isolation level is implemented.
    { .annotate }

    1.  Most DBs implement Repeatable Reads by keeping a version of rows being used in the transaction. This approach 
       doesn't help with Phantom Reads, as you can't version non-existent rows.

- **Lost Updates**: when two or more concurrent transactions read the same data, both make a modification based 
  on that data, and the second transaction's update overwrites the first one, effectively erasing its changes. 
  This leads to inconsistency as work updated by one transaction is lost due to overwrite from others.

To prevent these anomalies, Isolation property provide different levels of control. Below are few commonly
implemented Isolation levels listed from lowest to highest Isolation.

1. **Read committed**: transactions will only see committed changes. This solves dirty read as you’re sure the changes
   read are committed.
2. **Repeatable Read**: with this isolation level you can repeat reads consistently within your transaction, solving 
   non-repeatable and dirty read anomalies.
3. **Serializable**: concurrent transactions are executed as if they're being run one after another, essentially solving
   all concurrency anomalies.
4. **Snapshot**: allows transaction to read from a consistent snapshot of database without blocking writers.

??? note "Table Anomalies/Isolation"
    | *Isolation <br/>Level* | *Dirty Reads* | *Non-repeatable<br/> Reads* | *Phantom <br/>Reads* | *Lost <br/>Updates* |
    | --- |--------------------| --- | --- | --- |
    | *Read Committed* | :white_check_mark: | :x: | :x: | :x: |
    | *Repeatable Read* | :white_check_mark: | :white_check_mark: | :x: | :x: |
    | *Snapshot* | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
    | *Serializable* | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |


There’s mainly two different approaches for implementing isolation:

- **Pessimistic** approaches by using locks. These maybe on either row, table or page level. 
- **Optimistic** approaches keep track of transactions and fails one of them when they overstep each others isolation levels. 
  This approaches reduces the significant lock management overhead on DB, but requires additional handling for retries.  

!!! warning "NOTE"
    Postgres implements *Repeatable Read* as a *Snapshot isolation* level and as such you don’t get Phantom reads there
    but this might not be true for other DB system which implements Repeatable Read by maintaining version of rows.

The choice of isolation level balances data consistency with performance, as higher isolation provides more consistency
but can decrease performance.

### Consistency

Consistency ensures that a database remains in a valid state both before and after a transaction by guaranteeing 
adherence to all predefined rules, constraints, and triggers.

When defining *Consistency* across DB system, it can mean two different things:

1. *Consistency in Data* - consistent data w.r.t to the defined data model. 
   For example, having integrity across defined constraints (like primary key, foreign keys, data type), 
   cleaning orphaned references as per defined rules and constraints.
2. *Consistency in Read/Write* - consistently reading data across different instances of DB.

    ??? note "Read More here"
        Read consistency ensures a transaction sees the most recent committed changes immediately. 
        This consistency challenge is introduced due to *Replication*, specifically when data written to primary isn't 
        synced yet to replicas. This is usually done to optimize for performance.
        For example, **Eventual consistency** within a system provide higher performance but the application may 
        temporarily show stale data before eventually reflecting correct values. While **Synchronous replication** 
        offers stronger consistency at the cost of slower performance compared to asynchronous approach.

Consistency acts as a safeguard, ensuring that data integrity is maintained and preventing the database from entering 
an invalid or corrupted state due to incomplete or erroneous transactions.

### Durability

Durability ensures changes from a committed transaction are permanently stored on non-volatile storage (e.g., SSD, HDD)
— even if the system losses its power or crashes.

DB systems play around with this concept to optimize their performance since writing to disk is slower,
and instead you can write to memory first and then flush the changes to disk in bulk.
But this may compromise the durability under uncertain conditions,
so in addition DBs writes these changes in a compact format to a log file (`WAL` (1) ) on disk so that even if we 
lose the data in memory - the record can be replayed to recover the lost data. This is better because the changes 
written are compact and appended at the end of file.
{ .annotate }

1.  Write Ahead Log

!!! note "NOTE"
    The standard `write()` operation in OS caches writes in file system for better performance.
    If the system crashes during this time, the data in the cache is lost. Instead, DBs use the `fsync` operation
    to immediately write to disk, ensuring durability but at a performance cost.

For mission-critical systems, strong durability is non-negotiable; for less critical data,
eventual durability may be acceptable.