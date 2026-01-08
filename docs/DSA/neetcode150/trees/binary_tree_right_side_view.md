---
tags:
  - Trees
  - BFS
  - LC_Medium
  - Neetcode150
hide:
  - toc
---
# 199. Binary Tree Right Side View

[Problem Link](https://leetcode.com/problems/binary-tree-right-side-view/){target=_blank}

The right side view of a binary tree consists of the nodes that are **last visible at each level** when the tree is 
viewed from the right. This naturally leads to a [**level-order traversal**](binary_tree_level_order_traversal.md),
where we add the last node in each level.


??? note "Runtime Complexity"
    **Time**: $O(n)$, where $n$ is the number of nodes in the tree

    **Space**: $O(n)$, for the queue used during traversal


=== "Python"

    ```python
    --8<-- "docs/DSA/src/py/binary_tree_right_side_view.py:10"
    ```

=== "Go"

    ```go
    --8<-- "docs/DSA/src/go/binary_tree_right_side_view.go:3"
    ```