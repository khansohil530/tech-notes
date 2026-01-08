---
tags:
  - Trees
  - DFS
  - LC_Easy
  - Neetcode150
hide:
  - toc
---
# 110. Balanced Binary Tree

[Problem Link](https://leetcode.com/problems/balanced-binary-tree/){target=_blank}

A binary tree is height-balanced if, for every node, the heights of its left and right subtrees differ by at most one. 
Checking balance and computing height are **interdependent**, so doing them separately would lead to repeated work.
To avoid this, we use a **bottom-up traversal** where each recursive call returns two values:

- Whether the subtree is balanced
- The height of the subtree

For a given node, we first obtain the balance status and heights of its left and right subtrees. The current node is
balanced if both subtrees are balanced and $|leftHeight-rightHeight| \le 1$. The height of the current node is 
computed as $1 + max(leftHeight, rightHeight)$.

By propagating both balance information and height upward together, we ensure that each node is processed only once. 
The final result is the balance status returned from the root.

??? note "Pseudocode"
    ```
    function helper(node):
    if node is null:
    return (true, 0)
    
    (leftBalanced, leftHeight) = helper(node.left)
    (rightBalanced, rightHeight) = helper(node.right)
    
    balanced = leftBalanced AND rightBalanced AND abs(leftHeight - rightHeight) ≤ 1
    height = 1 + max(leftHeight, rightHeight)
    
    return (balanced, height)
    ```

??? note "Runtime Complexity"
    **Time**: $O(n)$

    **Space**: $O(h)$

=== "Python"

    ```python
    --8<-- "docs/DSA/src/py/balanced_binary_tree.py:7"
    ```

=== "Go"

    ```go
    --8<-- "docs/DSA/src/go/balanced_binary_tree.go:5"
    ```