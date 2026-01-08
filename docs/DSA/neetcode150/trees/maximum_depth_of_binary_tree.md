---
tags:
  - Trees
  - DFS
  - BFS
  - LC_Easy
  - Neetcode150
hide:
  - toc
---
# 104. Maximum Depth of Binary Tree

[Problem Link](https://leetcode.com/problems/maximum-depth-of-binary-tree/){target=_blank}

The **depth** of a binary tree is defined as the length of the longest path from the root to any leaf. This naturally 
suggests a **recursive, bottom-up** way of thinking.

For any given node, if we already know the maximum depth of its left and right subtrees, then the depth of this node is
simply **one (for the current node) plus the larger of those two depths**. This breaks the problem into identical 
subproblems on smaller trees.

The base case occurs when we reach a `None` node, which contributes a depth of `0`. As the recursion unwinds, each node
computes its depth using the results from its children, and the maximum value propagates back up to the root.

---

??? note "Pseudocode"
    ```text
        if node is null:
        return 0
        leftDepth = maxDepth(node.left)
        rightDepth = maxDepth(node.right)
        return 1 + max(leftDepth, rightDepth)
    ```

??? note "Runtime Complexity"
    <b>Time</b>: $O(n)$

    <b>Space</b>: $O(h)$


=== "Python"

    ```python
    --8<-- "docs/DSA/src/py/maximum_depth_of_binary_tree.py:7"
    ```

=== "Go"

    ```go
    --8<-- "docs/DSA/src/go/maximum_depth_of_binary_tree.go:3"
    ```