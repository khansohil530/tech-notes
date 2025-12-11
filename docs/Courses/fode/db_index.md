# B-Tree Index

Indexes in DBs are used to speed up your read queries.

Without Indexes, when you query for some records on your table, the DB would to sequentially search all the table
pages and gather records matching your query. This internally involves many steps like loading the page from disk
into memory and filtering the tuples, which is very inefficient and the time took grows linearly with the dataset.
To optimize this, you can maintain a sorted order in your table which allows you to use Binary Search to cut the time
taken to $O(log n)$. But this sort of algorithm operates very poorly at disk level, since it requires you to load
random pages causing lots of random I/O which under utilizes page buffer pool in memory. To minimize random disk I/O,
computer scientists invented a balanced tree data structure where each node (DB page) stored multiple keys along 
with pointers to multiple child nodes similar to below diagram.

``` mermaid
--8<-- "docs/Courses/fode/diagram/btree_index.mmd"
```

The tree consists of **root node** which is the starting point for search, **branch nodes** consisting of 
keys and pointers to next level of nodes, and **leaf nodes** which consists of actual keys and respective data.
With this you can store the root and branch nodes in memory most of the time as they're required for navigating
through the tree and only evict the leaf nodes as required, essentially reducing disk I/Os required to find a key.
The data structure is named as **B+ Tree** or commonly known as **B Tree** index.
Additionally, the leaf nodes also have reference to their respective neighbouring leaf node which saves you
the cost of traversing the whole tree in ordered to fetch a range of key.

Let's look into how DBs use this B-Tree structure to perform different CRUD operations.

1. **Search**: The search begins from the root node to respective leaf node by repeatedly comparing the search key
   with keys in current node and moving to appropriate child node. For example, if search key is between two keys in
   a node, move to the child pointer between them. Since the keys in each node are ordered, we can use binary search
   to determine the key location in $O(logk)$ time(k -> number of keys in each node or the degree of B-Tree).
2. **Insert**: Find the leaf node responsible for holding the new key and place the key into the leaf node. If the node
   is overfilled a **page split** occurs.
     - Divide the leaf node into two leaf nodes.
     - Add the middle key from leaf node into parent node.
     - Update the pointers to child nodes in parent node.
   
     This process is cascading as page split in low level nodes can move upto higher level nodes which result in 
     lots of movement of data on disk than required, causing **write amplifications** and spikes in disk usage.
     You can minimize this effect by using sequential index keys where the index is filled from left to right.

3. **Delete**: Find the node containing the key to be deleted and removing the key. If the node has too few keys after
   deletion, it must be restructured by either **Borrowing** a key from sibling node or **Merging** with a sibling node 
   which can cascade to higher level nodes potentially reducing the height of tree. 
4. **Update**: It usually involves a search for the key and then performing a delete operation, followed by an insert 
   for the new value.

To test it out and visualize these operation, you can use this 
[website](https://www.cs.usfca.edu/~galles/visualization/BTree.html){target=_blank}.  

## Secondary Index

But sometimes the table would also be queried using other fields which aren't primary key. In such cases you
can develop a **Secondary Index** on such field, which basically creates same B-Tree Data structure with the 
leaf nodes pointing to primary key or tupleId (in case of Postgres). 
Few things to keep in mind when using Secondary Index

- Some DBs (like InnoDB), Secondary Indexes use primary key internally to locate the row. In such case, the size
  of `pk` needs to be kept in check and choosing large `pk` can lead to larger secondary index essentially
  degrading its performance.
- Each additional secondary index causes more write amplification, as the DB now have to update all these structures
  which related incoming change in table. Postgres handles this a little differently by lazily updating the pointers
  with a regular cleanup process, and in the meantime it'll just mark the pointed tuple data with respective change so that
  DB can make the right decision.
- While `pk` is guaranteed to be unique, secondary index keys can be duplicate. Indexes perform best when the indexed 
  keys are selective, i.e. they map to small set of tuples. If a non-selective key is queried using index, the DB would
  have to query all the pointers which are scattered across pages which could potentially lead to more I/Os. In such
  cases, DB query optimizers decides to choose a less efficient execution plan like full table scans. Writes are similarly
  influenced by non-selective indexes, as DB have to maintain the large list of pointers for each such operation.


## Composite Indexes

Queries which include multiple AND filters for search can hugely benefit from  **Composite Indexes**.
**Composite Indexes** allows you to include more one keys in same index which are internally sorted and stored by
concatenating them together from left to right. For example, index on `(A, B, C)` would store the key as `A_B_C`.
The order in which columns are provided when creating composite index matters since you can still use the index
to filter tuples by selectively using left most columns in provided order but same can't be done for rest of
columns.

## UUID in B-Tree Indexes

UUID4 are completely random identifiers, two UUID4 Ids generated consecutively can never be ordered one after another.
Using such completely random (like UUID4) values in B-Tree index are disastrous and should be avoided as they negatively
impact both reads and writes.

- Writing random keys in B-Tree would result in frequent *page split* and fragmentation of data. This could be avoided
  if we had a sequential key for index which would fill up the index entries sequentially from left to right.
- Reading requires loading whole pages into a shared buffer pool memory. If the buffer is full, DB will eliminate the 
  oldest page to load up the current page. And in case of UUIDs, we’ll end up loading and removing page randomly 
  because we’ve no ordering. Instead, if we had sequential key for index, we’d have related pages which loaded up
  into memory, so if let’s say there's a surge in read for a category of products, they’ll be present in same nearby
  pages because their id are sequential. 

## Long-running transactions in Postgres 

In Postgres, any DML transaction touching a row creates a new version of that row. 
If the row is referenced in indexes, those need to be updated with the new tuple id as well. 
There are exceptions with optimization such as **heap only tuples (HOT)** where all the index doesn’t need to be 
updated immediately but that only happens if the page where the row lives have enough space (fill factor < 100%)

If a long transaction that’s updated millions of rows rolls back, then the new row versions created by this
transaction (millions in my case) are now invalid and shouldn’t be read by any new transaction. 
You have many ways to address this, 

- do you clean all dead rows *eagerly* on transaction rollback?
- Or do you do it *lazily* as a post-process?
- Or do you lock the table and clean those up until the database fully restarts?

Postgres does the *lazy approach*, using `VACCUM` command which is called periodically to remove dead rows and
free up space on the page.

What's the harm of leaving those dead rows in?  It's not really correctness issues at all,
in fact, transactions know not to read those dead rows by checking the state of the transaction that created them.
This is however an expensive check, the check to see if the transaction that created this row is committed
or rolled back. Also, the fact that those dead rows live in disk pages with alive rows makes an IO
inefficient as the database has to filter out dead rows. For example, a page may have contained 1000 rows,
but only 1 live row and 999 dead rows, the database will make that IO but only will get a single row of it.
Repeat that and you end up making more IOs. More IOs = slower performance.

Other databases do the eager approach and won’t let you even start the database before rolling back completely,
using undo logs. Both approaches have their pros and cons and at the end it really upto your workload which
approach suits you best.
