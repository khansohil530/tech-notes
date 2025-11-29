# Working with Indexes

This page will focus on how indexes works in DBs by demonstration using Postgres. 
We'll be using following Employees table with given schema

```postgresql
\d employees
                            Table "public.employees"
 Column |  Type   | Collation | Nullable |                Default                
--------+---------+-----------+----------+---------------------------------------
 id     | integer |           | not null | nextval('employees_id_seq'::regclass)
 name   | text    |           |          | 
Indexes:
    "employees_pkey" PRIMARY KEY, btree (id)

```

By default, Postgres builds B-Tree index around `pk`. To look around how a query performs, 
you also get `EXPLAIN ANALYZE` command which provides the query plan along with costs and other important information.
??? note "EXPLAIN ANALYZE"
    Postgres provides you two commands, `EXPLAIN` and `EXPLAIN ANALYZE` to understand how the query planner executes 
    (or intends to execute) a SQL query. It shows info like predicted plan steps (e.g., Seq Scan, Index Scan, Hash Join)
    estimated costs (`cost=...`), row counts, and row widths. 

    - `EXPLAIN` shows the query execution plan without actually running the query.
    - `EXPLAIN ANALYZE` executes the query and shows the actual query plan, execution time, and other actual figures.


Let’s look around how different queries behave when working with index in Postgres. The actual figures might vary
when executing same query due to optimizations in between like caching.

1. `SELECT` → indexed field, `WHERE` → indexed field (ignore the example query, it's just for demo purpose) 
    
    ```postgresql
    explain analyze select id from employees where id = 2000;
                                                              QUERY PLAN                                                           
    -------------------------------------------------------------------------------------------------------------------------------
     Index Only Scan using employees_pkey on employees  (cost=0.42..4.44 rows=1 width=4) (actual time=0.031..0.033 rows=1 loops=1)
       Index Cond: (id = 2000)
       Heap Fetches: 0
     Planning Time: 0.332 ms
     Execution Time: 0.072 ms
    (5 rows)
    ```
    DB decides to use *Index Only Scan* because we’ve an index on `id` which is used for filtering.
   *Heap Fetch* are 0 because the field we’re fetching is present in index itself, so we didn’t have to go to heap.

2. `SELECT` → non-indexed field, `WHERE` → indexed field
    
    ```postgresql
    explain analyze select name from employees where id = 50000;
                                                            QUERY PLAN                                                        
    --------------------------------------------------------------------------------------------------------------------------
     Index Scan using employees_pkey on employees  (cost=0.42..8.44 rows=1 width=6) (actual time=0.042..0.044 rows=1 loops=1)
       Index Cond: (id = 50000)
     Planning Time: 0.092 ms
     Execution Time: 0.152 ms
    (4 rows)
    ```
     `Index Scan` → for identifying the rows, then we’ve to go to heap for fetching the `name` field in `SELECT`.

3. `SELECT` → indexed field, `WHERE` → non-indexed field
    
    ```postgresql
    explain analyze select id from employees where name = 'P7o';
                                                           QUERY PLAN                                                       
    ------------------------------------------------------------------------------------------------------------------------
     Gather  (cost=1000.00..11310.94 rows=6 width=4) (actual time=0.729..52.861 rows=1 loops=1)
       Workers Planned: 2
       Workers Launched: 2
       ->  Parallel Seq Scan on employees  (cost=0.00..10310.34 rows=2 width=4) (actual time=11.900..27.233 rows=0 loops=3)
             Filter: (name = 'P7o'::text)
             Rows Removed by Filter: 333333
     Planning Time: 0.105 ms
     Execution Time: 52.886 ms
    (8 rows)
    ```
    
    Postgres will check if we’ve an index on the `WHERE` clause, if not we’ve to perform a Parallel Seq Scan
   (*Full Table Scan*). Still it tries to optimize this by using multiple worker.
    
4. Let’s index our name field and filter using a pattern match
    
    ```postgresql
    create index employees_name on employees(name);
    explain analyze select id from employees where name like '%P7o%';
                                                           QUERY PLAN                                                       
    ------------------------------------------------------------------------------------------------------------------------
     Gather  (cost=1000.00..11319.34 rows=90 width=4) (actual time=0.350..73.237 rows=11 loops=1)
       Workers Planned: 2
       Workers Launched: 2
       ->  Parallel Seq Scan on employees  (cost=0.00..10310.34 rows=38 width=4) (actual time=6.337..42.018 rows=4 loops=3)
             Filter: (name ~~ '%P7o%'::text)
             Rows Removed by Filter: 333330
     Planning Time: 0.344 ms
     Execution Time: 73.324 ms
    (8 rows)
    ```
    Even though we’d an index on `name` the DB couldn’t use it because we’re filtering for a pattern and
    not exact value. Since the pattern could fit multiple value, DB decides its more efficient to perform
    sequential scan.

## Different Execution Plans in Postgres

- **Sequential Scan (or Full Table Scan)**: When Postgres goes directly to the heap to fetches the query results. 
  This can be normally identified when query is using no filtering or filtering using field without index.
  But an unexpected scenario where this could happen is when it expects the query to bring a lots of rows even
  if we’ve filtering using indexed field. For example, `id!=10` will result only in single row so why go to
  index to fetch `id!=10` when its more efficient to just fetch the rows from heap and discard `id=10` 

    ```postgresql
    explain select name from grades where id!=10; 
                            QUERY PLAN                        
    ----------------------------------------------------------
     Seq Scan on grades  (cost=0.00..10.26 rows=500 width=15)
       Filter: (id <> 10)
    (2 rows) 
    ```

- **Bitmap Index Scan**: Postgres first creates a bitmap (bits where the position of bit indicates the page 
  number) for pages in heap and then scans the index to set bits which satisfies the condition. 
  After completing this, all the pages with set a bit in bitmap are fetched in one go. This is usually performed
  when we don’t have lots of rows which proves sequential scan efficient but enough rows that requires fetching
  in bulk.
    ```postgresql
    explain select name from grades where g = 25;
                                  QUERY PLAN                               
    -----------------------------------------------------------------------
     Bitmap Heap Scan on grades  (cost=4.18..8.43 rows=4 width=15)
       Recheck Cond: (g = 25)
       ->  Bitmap Index Scan on grades_g  (cost=0.00..4.18 rows=4 width=0)
             Index Cond: (g = 25)
    (4 rows)   
    ```
  Fetching rows from page is performed in **Bitmap Heap Scan** which discard rows which doesn’t satisfy the 
  query condition. Similar to this, there’s **BitmapAnd** and **BitmapOr** Scan which are used when using
  more than 2 indexed fields for filtering. In which case, Postgres would develop Bitmap for both the
  filters separately and then merge them into one using AND/OR operation depending on filtering condition.

- **Index Scan**: Postgres mostly uses index whenever we’ve a filtering criteria which uses indexed field.
  But it also depends on the amount of data we’re fetching, like previously discussed in Full Table Scans.
  If we’ve lots of rows its usually inefficient to fetch results from Index and then fetch the rows from heap.

- **Index Only Scan**: Postgres performs this when it can fetch the information asked by query from index itself
  (without going to heap). This is more efficient than index scans and can be made useful by adding
  a non-key column to the index. For example,
    ```postgresql
    create index grades_g on grades(g) include (name);
    ```
    We’ve included column `name` as non-key to index `grades_g` where `g` is the key column.
  So all the filtering will be performed on key column, but we can also fetch non-key column from index directly. 
  However, beware of including a non-key column to index, as it’ll increase the size of index and 
  will certainly impact on the cost for querying the index (as we’ve to load more pages for same index).

## When does DB use Index?

Suppose we’ve a table `T` with index on column `t1` (`idx_t1`) and index on column `t2` (`idx_t2`).
How will database plan for following query: `select * from T where t1=1 and t2=4` ?

- If we’ve lots and lots of rows, the optimizer will go ahead with full table scans and filter out data from there.
- If we’ve very few rows, the optimizer will use a single index to fetch the intermediate rows and
  filter out from them based on the other columns condition. Which index is used to fetch intermediate rows?
  The one which yields lesser rows in case of AND operation.
- If we’ve good enough rows (not too few or too many), we’ll develop bitmaps from both the indexes and
  BitmapAnd then to get final bitmap to scan the heap.

!!! note ""
    This decision about how many rows will turn up in a query is estimated based on statistics 
    which are precalculated by DB for each table. So always remember to update stats on tables before
    performing any critical operation.

You can also force database to use certain index by hinting it in query. 
For example, `select * from T where t1=1 and t2=4 /*+ index[t1 idx_t1] */`

Above case study also hints on what kind of column you should create index on to have maximum efficiency.
For example, if we’ve a column `state` and most of the rows use a same value for `state` creating an index on it
won’t help with searching your query because you’ll have so many rows with same state that its much more
efficient to just perform a sequential scan. 

!!! note "Create Index Concurrently"
    Most databases blocks writes when creating an index and this could impact live production system. 
    To solve this issue, postgres provide this feature to create index concurrently without stopping
    write in b/w the process.
    Command: `create index concurrently <index-name> on <table>(<column>)`
    It’d essentially create index sequentially and before exiting it’d wait for all ongoing transactions to
    complete so that they’ve been accounted for within the index.   

## Bloom Filters
Take a case where we need to query username to check if it’s been already taken by another user or not.

Directly querying the DB for presence of username is very slow if we’ve a lots of users signing up.
Instead, we can use intermediate cache to store username already taken up and query from these,
but this approach doubles our memory footprint.

To resolve this issue, we can use bloom filter which is essentially a fixed size bitmap on which set
bit indicates the possibility presence of that bit number and unset indicates its absence 
(the index of bit we need to check for a username can be found by using a hashing function % size of bitmap).
With this we can easily redirect most of absent usernames but to confirm the presence of one, 
we’ll have to query our database since the bit might be set by other username which collided on
same index of our bitmap. 
!!! note ""
    If all bits as set on our bloom filter, it'll become useless and if we’re always increasing the size
    of our bloom filter we’re moving toward more memory footprint. The actual implementation of bloom filter
    accounts for this pretty well, making it essentially works like above.
