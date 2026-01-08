---
tags:
  - Trees
  - DFS
  - LC_Easy
  - Neetcode150
hide:
  - toc
---
# 543. Diameter of Binary Tree

[Problem Link](https://leetcode.com/problems/diameter-of-binary-tree/){target=_blank}

The diameter of a binary tree is the length of the longest path between any two nodes. This path **might not pass 
through the root**, so we need to consider all possibilities. 

$diameterNode = height(leftSubTree) + height(rightSubTree)$ => for each node, we need height and maximum diameter for 
left and right subtree. We can compute both values together using post-order variant of DFS. 

From the left and right children, we obtain their heights and diameters. The height of the current node is simply
`1 + max(left_height, right_height)`.

The maximum diameter of the current subtree node is maximum of either:

- Maximum diameter from left subtree
- Maximum diameter from right subtree
- Diameter passing through the current node, which has length `left_height + right_height`

This way we can calculate both height and maximum diameter for the current node in the same iteration. 
Finally, the diameter of the entire tree is obtained from the value returned at the root.

---

??? note "Pseudocode"
    ```text
    function helper(node):
    if node is null:
    return (0, 0)
    (lh, ld) = helper(node.left)
    (rh, rd) = helper(node.right)
    
    height = 1 + max(lh, rh)
    diameter = max(ld, rd, lh + rh)
    
    return (height, diameter)
    ```


??? note "Runtime Complexity"
    <b>Time</b>: $O(n)$

    <b>Space</b>: $O(h)$


=== "Python"

    ```python
    --8<-- "docs/DSA/src/py/diameter_of_binary_tree.py:7"
    ```

=== "Go"

    ```go
    --8<-- "docs/DSA/src/go/diameter_of_binary_tree.go:3"
    ```