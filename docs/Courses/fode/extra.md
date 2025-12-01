# Extra Topics

## Index Selectivity

You want to index column which provides as few rows when filtered through. 
For example, we index a column which stores Gender which can have 3 values, 
and filtering any gender would result in massive result set. So instead database would go for heap scan.  

## Postgres TupleId
All indexes in Postgres point to a tupleId which is used as the key to cluster around the table.
Whenever you make any update, a new tupleId will be generated for the same row.
Due to this Postgres needs to update all the indexes pointing to this row with old tupleId to new tupleId, 
which isn’t optimal when we’ve a lots of indexes. Still it tries to optimize this by using 
*Heap Only Tuple* (HOT) optimization where it’ll immediately update the tupleId in index of the updated column.
But this could still cause issue for index pointing to old tupleId. This is resolved by storing some metadata
on old tupleId page, which points to the latest tupleId of this row **only if the new tupleId is on same page**. 
This can be used to advantage by using the *fill factor* configuration which tells the limit upto which a
page can be filled to leave some space for updates and inserts. 

## WAL, Redo and Undo Logs

Logs help DBMS to ensure durability and crash recoverability. Whenever you commit writes, it must be persisted 
by database. It basically means the data must be present even after shutting of the DBMS. 
One way to do this is to directly flush the changes to disk after each commit, but this way your commit
operation will become slower as writing to disk is heavy (because writes to disk are in whole pages, 
not individual bytes). Instead, databases, keep changes in memory and marks the pages as dirty to indicate the 
page has been updated. This approach is fast but can compromise with the persistence of data. To be 100% sure
with persistence, database instead maintains logs contains each of these changes as a tiny delta which are 
appended to the end of the log file and can be replayed to update the state of database. This log is called
*Write Ahead Log* (WAL). WAL can’t grow infinitely large, so after it grew to some size we’ll flush the changes
in WAL to disk and clear our WAL since it no longer needs to maintain the logs. And we’ll restart again.
This flushing is of changes to disk is called *checkpointing*. Checkpointing operation are very heavy 
operations as it includes a lots of IOs and compute operations which can cause spike in systems resource usage
and impact its performance. To make this tolerable we can make checkpoints smaller so that we can flush
to disk frequently, but not smaller enough to impact the writes operation.  

## Endurance of SSDs
SSDs store data within page present in fixed blocks. There’s no mechanical apparatus like hard disk which
makes them faster compared to them. However, SSDs can update certain page until a limit after which the bytes
are no longer usable essentially reducing the size of SSD. Due to this workload involving updates aren’t
considered good for SSDs. For example, B-Tree indexes which restructure themselves are considered bad for SSDs
but index like LSM-Tree which only appends entries are considered optimal for SSDs.

## Postgres Architecture

Postgres is a SQL row-based database which follows MVCC storage model where each row can have multiple 
physical version on disk with the last version is the latest. Basically every insert/update/delete operation
create a newer tupleId of row indicating the current version (lookup pros and cons of this decision).
Then it uses processes instead of threads for work. Let’s discuss all the process below: 

- *Post master*: first process to spawn on startup, acts a parent process for all other processes,
  works at listener to connect external application over network. Every other process is forked from this process.
- *Backend Processes*: each client connection receives its own backend processes to receive and process the 
  request. This is a bad choice as processes requires more memory and CPU context switching which impacts
  the performance. But postgres avoid this by offloading most of the work outside these processes.
- *Shared Memory*: Or shared buffer pool where most of the data which is shared among processes are present like
  WAL records, pages, etc.
- *Background workers*: backend processes uses these workers to outsource most of its work like querying
  based on the generated plan. If a parallel plan is needed, these background workers will be picked up and
  assigned respective work.
- *Auxiliary Processes*: 
  - `bw` (background writer) wakes up periodically,and write pages from shared memory to disk to free up the memory.
  - `cp` (checkpointer) directly flushes the WAL records and pages to disk and creates a checkpoint which 
    indicates that data uptil this point is consistent. 
  - `lg` (logger) is used to writes logs. 
  - `avl` (Auto vacuum launcher) launches autovacuum workers. 
  - `wa`(WAL archiver) responsible for backing up WAL records
  - `wr` (WAL receiver) runs on replica to replicate data from WAL records
  - `ww` (WAL writer) writes record to WAL and flush them to disk.
  - `st` (startup process) which is actually the first process to start whose role is to check if the pages
    are consistent with WAL records, if not mark them as dirty pages. As this needs to be done before any
    client connects to database, it should be the first process to start.
- *AutoVacuum Workers*: Periodically Vacuums the database which essentially means it frees up old tupleId
  which are no longer required by any transactions. Vacuum includes much other stuff you can explore online.
- *WAL Senders*: responsible for sending WAL records from client to replicas.


## Table Joins using Hash Tables 
In simple terms one relation is mapped to another using key and value as respective column and foreign key.
Usually the column with less value is picked up as key, as the hashtable will be smaller than.
To map the value to foreign key — you fetch the foreign key row, and the row from source table with same
foreign key, map it in hash table and finally u can use it to join the tables. 


## Storage of NULL value in database systems
NULL in database system indicate that the given place isn’t allowed to store data.
Postgres uses a null bitmap in front of every row to indicate if the given column is null.
The size of this bitmap starts with 8bits for 8 column each and with more columns it increments with 8 bytes 
in size (so 9-64 column would use another 8 bytes). This helps you save space on disk, as you no longer
have to persist the column for null value and allows you to fit more rows within same page.
However, be careful when working with NULLs as they’re widely inconsistent:

- `SELECT COUNT(FIELD)` would ignore counts of row having null value in given field but 
  `SELECT COUNT(*)` would give you correct count.
- You can’t compare null value, you can just check if It's null or not null. You can’t use `T in [NULL]`
- Not all database support NULLs in indexes, check if beforehand.


## Write Amplifications
Basically when the actual work done for write is much more than the logical work required. 
For example, postgres creates new row even for every update/delete (for versioning).
Now this new row with new tupleId must be updated in all existing index on this table. 
This is somewhat optimized by updating the index of columns whose value had been changes and rest of
untouched column index can be updated in background. Those old rows will now also point to this latest version
known as *heap only tuple* (HOT), because these version must be present in same page which is managed by
using *fill factor* configuration. Then you’ve WAL writes to achieve durability. All these amplification
happens at database level where the database system is responsible. 

SSDs Disks/Storage also causes write amplifications. SSDs uses charge entrapment where electrons are trapped
in different levels on a atom. The way these electrons are trapped can be used as a way to store information
in them and this configuration of electrons isn’t lost even in absence of electricity. 
These cells are then arranged into rows which are then arranged into pages and pages into blocks.
For Database Systems, we just need to be aware of page and block level. 

- Writes to disk are on page level, and for SSDs you can simply write data to new pages easily.
- When you want to update the data on a already existing page, you’ll write the data to a new page and mark the existing page as stale.
- Finally, you can’t clean single pages in SSDs, you’ve to clean the whole block to free up space.

To clean up blocks with both stale and active pages, SSDs have a garbage collection program which moves
active data to new block and then clean the block to free up storage. All these processes due to update
requires additional work in SSDs causing write amplification.