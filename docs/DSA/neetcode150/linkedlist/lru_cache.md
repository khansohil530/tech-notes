---
tags:
  - Linked List
  - LC_Medium
  - Neetcode150
hide:
  - toc
---
# 146. LRU Cache

[Problem Link](https://leetcode.com/problems/lru-cache/description/){target=_blank}

To support both `get` and `put` operations in $O(1)$ time, we combine a **hash map** with a **doubly linked list**.

- The hash map stores key-to-node mappings, allowing direct access to cache entries in constant time. 
- The doubly linked list maintains the **usage order** of the cache, where the least recently used (LRU) node is near
  the head, and the most recently used (MRU) node is near the tail.

We use two dummy nodes, `lru` and `mru`, to simplify insertion and removal. When a key is accessed or updated, its node
is removed from its current position and reinserted just before the `mru` node, marking it as most recently used.

For the `get` operation, 

- if the key exists, we move the corresponding node to the MRU position and return its value. 
- If the key does not exist, we return `-1`.

For the `put` operation, 

- if the key already exists, we remove its current node before inserting the updated one. 
- After insertion, if the cache exceeds its capacity, we remove the node next to the `lru` dummy node, which represents
  the least recently used entry, and delete it from the hash map.

This design ensures constant-time operations while correctly maintaining the LRU eviction policy.

??? note "Runtime Complexity"
    <b>Time</b>: $O(1)$ for both `put` and `get`

    <b>Space</b>: $O(n)$


=== "Python"

    ```python
    --8<-- "docs/DSA/src/py/lru_cache.py"
    ```

=== "Go"

    ```go
    --8<-- "docs/DSA/src/go/lru_cache.go:2"
    ```