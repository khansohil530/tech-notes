# Database Replication

Replication in DB system involves sharing data between redundant database instances in order to
improve accessibility, reliability and fault tolerance. There are different kind of architecture for this

1. *Master/Standby Replication (Leader/Follower)*: Single database instance (node) will take all write operations
   and distribute to Standby nodes. Standby nodes are only used for read operations. The consistency of reads for
   standby nodes depends upon how long it take to propagate writes from master to standby, essentially ranging 
   the consistency from synchronous (strong) consistency to asynchronous (eventual) consistency.
2. *Multi-Master Replication*: Similarly, to improve write you can have multiple master nodes. But this
   strategy is more complicated and needs to handle conflicts when writing. 

The mode of replication can be *synchronous* or *asynchronous* depending on how data propagation is handled.
Synchronous replication makes client wait till the transaction is completed on master as well as standby nodes.
Cassandra uses this mode where the client accepts writes as successful until *quorum* (n+1/2) nodes have
replicated the data. Asynchronous replication shows writes as successful write after it's committed on master.
The replication is handed over to some backend process which moves data to standby nodes periodically.

Advantage of using replication is you get horizontal scaling, and you can split standby database into 
regions so that queries closer to certain group of user can use their region specific database instance.

Disadvantage are your system might not like eventual consistency model and when going for strong 
consistency the writes are slower. And if you want to go multi-master architecture, it's much more complex.