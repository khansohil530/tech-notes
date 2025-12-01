# Database Engines

Database Engines or Storage Engines are binaries which provides low level disk storage operations like storing and
retrieving data from disk, providing compression over stored data, handling crash recovery, indexing over data, etc.
When you run an DML query, the parse and optimizer translates the query into these low level operation which
are then executed by engine.

Other benefits of designing these engine is you don't have to start from scratch to build an DB system, which 
supports your specific use case. You can build new features on top of existing engines.
Some DB systems (like MySQL and MariaDB) even allows you to switch engines. Below are the most popular
DB engines developed over time as requirement for applications evolved.

## MyISAM

MyISAM (which stands for **I**ndexed **S**equential **A**ccess **M**ethod) was default storage engine for MySQL in 
its earlier days (before version 5.5). It became popular during early internet days due to its 
simplicity (lightweight) and faster reads. But it didn't support essential features which are required by
modern application like
- No transactions, ACID making data storage unsafe and inconsistent.
- Write operations locked entire table resulting in poor write performance in concurrent environmnet.
- Didn't provide support for foreign key

## InnoDB

Developed to address the problems encountered by MyISAM, InnoDB became the default storage engine for MySQL
since version 5.5 by providing essential features like ACID transactions, row level locking, crash recovery,
foreign key constraints and more. It stores data in a B+ Tree Clustered Index around primary key.  

## SQLite

A lightweight, serverless, self-contained SQL engine which became popular for being extremely lightweight,
allowing it to be embedded into applications directly. There's no separate process running the DB in background,
it's just your application process which reads and writes data from the file directly. Additionally, you don't
need to configure or setup any additional step, just include the SQLite file. Since whole database is stored
in single file, copying, versioning, deploying is very fast and simple. It's ACID compliant, works across 
different platforms (like Windows, Linux, macOS). It's use across various popular apps like browsers, mobile apps,
and IoT devices. However, due to this simplicity its not ideal for large datasets or high throughput writes.

## BerkeleyDB

An embedded key–value storage engine that provides ACID transactions, multiple indexing methods,
and configurable concurrency control. It's extremely fast, reliable, and flexible, making it perfect for
embedded systems, but it doesn’t provide SQL and requires the application to handle schema and data logic.

## LevelDB

LevelDB is a fast, embedded, single-threaded, key–value storage engine using an LSM-tree design which allows
fast sequential writes, high compression, lower write amplification and efficient storage layout. 
Range or prefix queries over sorted key are much faster. Writes are optimized using LSM design while
Reads are optimized using a combination of Memtable, SSTable and Bloom Filters.
It's popularly used across application as Blockchain, Caching layers, and Web Browsers (much lighter than SQLite).

## RocksDB

Built as a fork from LevelDB, RocksDB introduced enhancements to provide low latency operation, higher write
throughput optimized for SSDs and large scale production workloads. It's popularly used across systems like
Kafka streams, Blockchains, Cockroach DB. You can check out the summary of major enhancements over its
ancestor in following table.


| Feature                 | **LevelDB**                    | **RocksDB**                                         |
| ----------------------- | ------------------------------ | --------------------------------------------------- |
| **Concurrency**         | Single-writer, limited readers | Multi-threaded, many writers, parallel compaction   |
| **Performance**         | Good for small apps            | Extremely high throughput for large workloads       |
| **Compaction**          | Simple, single-threaded        | Multi-threaded, advanced, tunable compactions       |
| **Transactions**        | No transactions                | Full ACID transactions via WriteBatch + WAL         |
| **Column families**     | Not supported                  | Supported (namespaces like MySQL tables)            |
| **Tuning options**      | Very few                       | Hundreds of tuning knobs for memory, IO, compaction |
| **Memory optimization** | Basic                          | Block cache, compressed cache, rate-limiting        |
| **Storage types**       | Disk only                      | Optimized for SSD / flash, persistent memory        |
| **Backup & restore**    | Manual                         | Built-in APIs for backup, checkpoints, replication  |
| **Use cases**           | Small embedded apps            | Large-scale server apps, distributed systems        |


