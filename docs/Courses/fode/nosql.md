# NoSQL Overview

DB Systems can be fundamentally divided into two parts based on their architecture:

- **Frontend**: communication layer which implements the APIs exposes internal functionality to clients and decides
  the format of data used for this communication. The most popular format of communication with DB was using table,
  which SQL excelled in but as internet evolved various other data structures like JSON became commonly used across
  different places. This change became a catalyst to development of DBs using different data formats for communication,
  like documents data which used JSON like structure for encoding its data. All these data format were designed because
  of the need of performance in their specific use case. 
- **Storage Engine**: primarily focuses on storing data on disk efficiently, which involves working with bytes so
  format of data is irrelevant here. Other responsibility of storage engine involves working with indexes to 
  store/fetch data, managing data files, using compression to save storage space, providing crash recovery,
  and other features like ACID.

The major difference in SQL and NoSQL is between Frontend where the data format is changed to document
from rows, and the API format which changed from SQL to simple getter and setter commands.

## MongoDB Architecture

MongoDB is a document based NoSQL database popular for its schemaless tables.

**MongoDB version <4.2**: This initial version used **Memory Map Index** on `_id` which is a B-Tree index where
the leaf node contained a 64-bit pointer to the document. The 64-bit was composed of filename (32-bit) and
offset (32-bit) to locate the document in that file, using which the OS can directly jump to the document
location and retrieve it. The downside of this design was that any update in page size or data files could 
mess up the whole offset based index. Also, it only supported collection level lock for concurrent transactions.

**MongoDB version 4.2-5.2**: The storage engine was replaced with **WiredTiger** which solve the problem of
collection level locks by allowing document level locks. It introduced compression which allowed mongo to fit more
documents within a page.  The storage model now included a hidden clustered index (B+ Tree) on a field
*recordId*. Any indexed field would reference this recordId which would then in turn point you to respective
page on the clustered index. The problem now what that primary index on `_id` became slower since we’ve
to do two lookups (find recordId → find page).

**MongoDB version > 5.3**: Introduced clustered collection, which basically built a clustered index
around `_id` field (avoiding recordId). The problem with this was that `_id` is 12 bytes long,
due to which secondary indexes grew much larger.

## MemCached Architecture

Memcached is a high-performance, distributed, in-memory key-value cache used to speed up dynamic web
applications by reducing database load. Its architecture is intentionally simple and optimized for speed.
It uses a client-server architecture where server are responsible for storing data in RAM as key-value pair,
and client are libraries in apps which decides how to fetch the key within the cache.
!!! note "Distributed"
    The cluster is logically distributed, but coordination is handled entirely by the clients.
    This makes the system highly scalable and avoids complex distributed consensus.


*Memory Management*: It organizes memory allocation in slabs where each slab holds item of same size. 
As new items are added, they’re written to a pre-allocated page serially. The page is divided into equal fixed 
size chunks whose determined by the assigned slab class. Each item uses whole chunk/s to stored their
information, as such you might have unused memory within each chunk. This is minimized by using the most
appropriate slab class for each item. The Slab class vary from class 1 (chunk size of 72 bytes) to
class 43 (chunk size of 1MB). This is done to avoid memory fragmentation while keeping allocation fast and 
predicatable to ensure high throughput for caching. 

!!! note "Fragmentation"
    When storing data sequentially without any strategy, freeing unused memory leaves small gaps of
    free memory scattered across the physical memory, this problem is known as *fragmentation*.
    Fragmentation makes it difficult to get a continuous block of memory large enough to store your 
    new data item even though there’s more than enough memory present. OS overcomes this problem using 
    *virtual memory*, which basically gives us a continuous block but behind the scene is mapped to 
    multiple small area on physical memory. This still isn’t optimal because to fetch a single block
    OS will have to fetch multiple pages and reassemble the memory fragments, which is why it's always
    better to avoid memory fragmentation.  

- *Threading*: Memcached used TCP transport as default to connect to remote clients. 
   The listener thread creates a TCP socket on port 11211. After accepting the connection, 
   it's distributed among a pool of worker threads which is responsible for the requested read/write
- *LRU*: Memcached use LRU eviction policy when it can’t find any space for new keys. 
  Even if you put a TTL on a key that it can’t expire before given time, the eviction policy can still
  remove this key. LRU is implemented as a linked list where each node is a key-value pair in linked list,
  and each slab has its own linked list. When an item is accessed, its move to the head of Linked List.
  As result, unused items are pushed down to the tail and can be removed when needed. 
  One of the disadvantage of using Linked List LRU approach, you need to lock the entire linked list before
  any update which serializes write operations and the multithreaded model wasn’t effective. 
  This model was then updated to have LRU Linked List per slab, which reduced the locking to per slab.
  Later on, LRU updates were made once every 60 second to reduce locking further. In 2018, the model was 
  completely redesigned by breaking LRU into subclass based on temperature but the problem of lock still
  persists in keys belonging to same temperature.
- *Reads and Writes*: It uses hash to index the key to a memory location where its value is present.
  For reads, it’ll look up the key on the linked list at designated memory location and update the key’s
  position to head. For writes, if the memory location is free, it’ll create a new pointer and a slab class
  is assigned otherwise it’ll handle *Collision* using chaining.  If the chain becomes too large to impact 
  read, Memcached resizes the hash and shifts everything to flattens the structure.

## Redis Architecture

Redis is an in-memory data structure store that supports caching, message queuing, real-time analytics, and more.
Its architecture is more feature-rich than Memcached while still being extremely fast.

It uses a single threaded event loop model for all its operation. To allow processing of multiple clients in 
parallel, it uses I/O multiplexer (epoll/kqueue/select) over the single thread. 

One of the biggest difference b/w Memcached is its built in support for advanced data structures like
Streams, Bitmaps, HyperLogLogs and more. All its primary data lives in RAM which allows faster read and writes,
but it also supports durability optionally. You can persist data to disk in two ways:

- Journaling using an **append only log** (AOL) file, where logs are added for every insert/update. 
  These logs can be later replayed to restore the data of Redis upto the latest state of logs.
  This requires another thread to append this logs.
- Snapshots, where data is flushed to disk periodically. This could risk data loss but the process is 
  much faster and the backup file is much compact and smaller.

For communication, it uses its own wire protocol (known as RESP) build on top of TCP request/response model.

Other popular features supported by Redis includes:

- built-in publish–subscribe messaging system, where clients can publish messages to channels and subscribers
  to these channels would receive them in real time.
- supports built-in distributed clustering, across different models like 
  - Sharding where Redis Cluster is partitioned across multiple nodes using hash slots.
  - Leader-Follower async replication
- supports module to extend custom features like RedisBloom for supporting BloomFilters.

