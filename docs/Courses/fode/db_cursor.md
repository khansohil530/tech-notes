# Database Cursor

Suppose your SQL query is fetching lots of rows, and it takes some time before getting back the result. 
This is because the query needs to do plan query execution, gather all records, then move this data over to TCP
and transmit it over network and finally collect all the data on client side. All these steps become increasingly
slower as we fetch more and more rows. Also, sometimes few of the clients won't have the required memory to
store all the results from such huge queries.

This scenario can be avoided using server side **Cursor**, which allow you to encapsulate a query into a token and
fetch few entries at a time from the query result using that token. This allows you to have immediate result for
few queried entries allowing the UI to partially load data which wouldn't affect the user workflow. You can save 
memory usage on client side by processing few rows at a time. You can also cancel the query midway, if you're
satisfied with your results. And all these can be streamed through websocket allowing smooth workflow on client.
But this shifts the load of managing state of cursor to backend/server, which may causes resource starvation
if not managed properly. Also, long-running cursors involving transactions would continue to block writes
causing bad write performance.

The above-mentioned scenarios are applicable for **Server Side Cursors**, where the server is responsible for
managing the state of cursor. You can also use **Client-Side Cursors**, which provided limited sets of queried
entries by sorting and filtering data in batches (using keywords like `offset`, `limit`, `sort by` in SQL).

## Cursors in SQL Server

To understand more about how cursors are implemented in real production DBs, let's look at the implementation
of cursors in Microsoft SQL Server. Most of the information is derived from their docs referenced 
[here](https://learn.microsoft.com/en-us/sql/relational-databases/cursors?view=sql-server-ver17){target=_blank}.

They define cursors as an extension over results sets to provide processing like positioning the cursor to 
a specific row in result, retrieve/modify one or many rows from current position, allowing different
visibility to changes made by others and finally allowing access to stored procedures and triggers.
SQL Server supports four cursor types.

1. **Forward-Only**: is a forward-only and read-only cursor which doesn't support scrolling. They only allow you to 
   read data serially from start to end. Rows are retrieved from DB as they're fetched, which allows it to read 
   modification made by other users even after its declaration. This simplicity allows it to fetch result quickly
   while keeping low memory footprint.
2. **Static**: Builds the complete result set in a temporary DB (`tempdb`) when this cursor is opened allowing it
   to hide modification made by other users after cursor declaration.  Due to the static copy, you can scroll through
   result set quickly, but storing the entire result set uses more memory.
3. **Keyset**: controlled by a set of unique identifiers, or keys, known as the `keyset`, stores the key and respective
   row in temporary DB (`tempdb`). When the result is fetches, DB will fetch the related row specified by key from table
   as such it only detects modification like updates and deletes but not inserts since the key for new row wouldn't
   be present in `tempdb`. Since only keys are stored in `tempdb`, it uses lesser memory than `STATIC` cursor  but
   fetching respective rows requires additional lookup to underlying table making scrolling slower.
4. **Dynamic**: reflects all committed changes made to result set when scrolling over the cursor. This allows your
   cursor to view the most recent changes, beneficial displaying real-time updates. But since the ordering
   of data isn't constant you can use positioning reliably. 

