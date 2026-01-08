---
tags:
  - Trees
  - LC_Hard
  - Neetcode150
hide:
  - toc
---
# 124. Binary Tree Maximum Path Sum

[Problem Link](https://leetcode.com/problems/binary-tree-maximum-path-sum/){target=_blank}

A path can start and end at any node, but it must go **downward through parent–child connections**, and it cannot 
branch in both directions. To calculate the maximum path sum of the whole tree, we need to answer two questions for each 
node

1. What is the **maximum path sum passing through this node**?
2. What is the **maximum path sum we can extend upward** to this node’s parent?

Both these can be solved using a bottom up DFS, where we first compute the maximum path sums from the left and right
subtrees. If either side contributes a negative value, we discard it by taking `max(value, 0)`, since including a 
negative path would only reduce the total sum. The path that **passes through the current node** is `node.val + 
left_max + right_max`. This represents a complete path with the current node as the highest point. Since this path 
cannot be extended further, we update our global result by comparing this value. 

However, when returning to the parent, we can only choose **one direction** (left or right), so that returns value can
be consumed by path to make max sum. So the value returned by the helper function is `node.val + max(left_max, right_max)`.

??? note "Runtime Complexity"
    **Time**: $O(n)$, where $n$ is the number of nodes in the tree

    **Space**: $O(h)$, where $h$ is the height of the tree due to recursion stack

=== "Python"

    ```python
    --8<-- "docs/DSA/src/py/binary_tree_maximum_path_sum.py:7"
    ```

=== "Go"

    ```go
    --8<-- "docs/DSA/src/go/binary_tree_maximum_path_sum.go:4"
    ```