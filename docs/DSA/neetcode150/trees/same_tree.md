---
tags:
  - Trees
  - DFS
  - LC_Easy
  - Neetcode150
hide:
  - toc
---
# 100. Same Tree

[Problem Link](https://leetcode.com/problems/same-tree/){target=_blank}

To determine whether two binary trees are the same, we compare them **node by node**, ensuring that both structure 
and values match exactly and do this **recursively** for all left and right subtrees. At each pair of nodes, we need to
 check the following possible cases:

- If both nodes are `None`, the trees match at this position.
- If both nodes exist, their values must be equal, and their left and right subtrees must also be identical.
- If one node exists and the other does not, the trees differ.

By applying this logic recursively, we ensure that every corresponding node in both trees is checked in the same 
relative position.

??? note "Pseudocode"
    ```text
        function isSameTree(p, q):
        if p is null and q is null:
        return true
        if p is null or q is null:
        return false
        if p.val ≠ q.val:
        return false
        return isSameTree(p.left, q.left) AND isSameTree(p.right, q.right)
    ```

??? note "Runtime Complexity"
    **Time**: O(n)
    **Space**: O(h)

=== "Python"

    ```python
    --8<-- "docs/DSA/src/py/same_tree.py:7"
    ```

=== "Go"

    ```go
    --8<-- "docs/DSA/src/go/same_tree.go:3"
    ```