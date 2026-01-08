---
tags:
  - Trees
  - BFS
  - LC_Medium
  - Neetcode150
hide:
  - toc
---
# 102. Binary Tree Level Order Traversal

[Problem Link](https://leetcode.com/problems/binary-tree-level-order-traversal/){target=_blank}

Level order traversal requires visiting nodes **level by level**, from top to bottom and left to right. This naturally 
suggests using a **queue**, since queues process elements in the same order they are added.

We start by placing the root node into the queue. Each iteration of the outer loop represents processing one full level
of the tree. At the beginning of each level, we record the current size of the queue, which tells us how many nodes 
belong to that level.

We then process exactly that many nodes:

- Remove a node from the queue.
- Add its value to the current level list.
- Push its left and right children into the queue for processing in the next level.

Null nodes are skipped to avoid unnecessary work. After processing all nodes of the current level, if the level 
contains any values, we append it to the final result. This continues until the queue is empty, meaning all levels 
have been traversed.

??? note "Runtime Complexity"
    **Time**: $O(n)$, where $n$ is the number of nodes in the tree

    **Space**: $O(n)$, for the queue used during traversal



=== "Python"

    ```python
    --8<-- "docs/DSA/src/py/binary_tree_level_order_traversal.py:10"
    ```

=== "Go"

    ```go
    --8<-- "docs/DSA/src/go/binary_tree_level_order_traversal.go:3"
    ```