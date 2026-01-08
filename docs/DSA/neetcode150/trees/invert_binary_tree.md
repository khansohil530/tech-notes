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
# 226. Invert Binary Tree

[Problem Link](https://leetcode.com/problems/invert-binary-tree/){target=_blank}

Inverting binary tree applies uniformly at every node: the left and right children simply need to be swapped. 
Since this operation depends on the same transformation being applied to each subtree, the problem fits a recursive 
approach.

Starting from the root, we think in terms of subproblems: if we already know how to invert the left and right subtrees,
then inverting the current tree is just a matter of swapping those two results. This leads directly to a depth-first 
traversal where each node waits for its children to be processed before performing the swap.

The base case occurs when we reach a `None` node, which requires no action. As the recursion unwinds, each node swaps its
left and right pointers, effectively inverting the tree from the bottom up.

---

??? note "Pseudocode"
    ```text
        if node is null:
        return null
        left = invertTree(node.left)
        right = invertTree(node.right)
        node.left = right
        node.right = left
        return node
    ```


??? note "Runtime Complexity"
    <b>Time</b>: $O(n)$, where $n$ is the number of nodes in the tree

    <b>Space</b>: O(h) due to recursion stack, where $h$ is the height of the tree


=== "Python"

    ```python
    --8<-- "docs/DSA/src/py/invert_binary_tree.py:7"
    ```

=== "Go"

    ```go
    --8<-- "docs/DSA/src/go/invert_binary_tree.go:8"
    ```