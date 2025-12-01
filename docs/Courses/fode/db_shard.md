# Database Sharding

Another way to divide you dataset into smaller chunks is using Shards. It divides the table into smaller tables 
similar to partitioning, but the key difference is each table lives in a separate DB instance with sharding.
Using separate DB instance would provide u additional benefits like more resources (CPU, memory) for each
sharded table, or network advantage by geographically locating the shard closer to client, or provide specific
security standards for specific group of users. 

You can create shards based on keys like, zipcode which represent an area geographically,
or using range of values on some number field. But what if your have to shard on some random text?
Mostly shard keys which can’t be distinguished into groups are grouped using **Consistent Hashing**.

??? note "Consistent Hashing"
    Consistent hashing is an algorithm which distributes different shards as points on a ring like data structure
    and the range of values which falls on a pie of the ring belong to a single shard.
    To map values to ring, you’ve different hash function which evenly distributes the shard keys so that no 
    single shard is overused.

Following are few key advantages and disadvantages of sharding summarized briefly.

| Advantage                                                                                                     | Disadvantage                                                                                    |
|---------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------|
| You get scalability in data, memory, CPU, etc                                                                 | Makes the client complex as it needs to be aware of each shard.                                 |
| You get smaller tables and as such smaller indexes                                                            | Transaction across shard wouldn’t be atomic anymore.                                            |
| You get security as data can live in separate database instances for users which require most secure storage. | Rollbacks would be expensive                                                                    |
|                                                                                                               | Schema changes are hard                                                                         |
|                                                                                                               | The query must know which shard to hit, otherwise the client would’ve to search every database. |

**So when should you use sharding to optimize your DB performance?**
Sharding should be your last option for optimizing your database performances due to the complications involved 
which most of the time are unnecessary. There are many other optimization trick you can do before ultimately using
sharding

- You can look into *horizontal partitioning* before it which allows you to divide your table into smaller
  chunks and provide smaller index for all these partitions.
- Then, if your reads are slow, you can look into *replication* — where you can employ multiple read replicas 
  to distribute the load over a single server. 
- If you’re facing slow writes, maybe try to separate servers based on regions.

Sharding does provide you with scaled writes and read, but you’ll have to leave behind features like ACID
transactions. Also, the client coupling to database is strong since client needs to be aware of each shard.
And resharding or changing business logic also becomes complicated. So avoid using it until its absolutely 
necessary.

!!! note "Vitess"
    To avoid coupling you can transfer sharding client logic to a backend app which will handle which shard 
    should the query be directed to. One such backend app is [Vitess](https://vitess.io/docs/21.0/overview/whatisvitess/).
