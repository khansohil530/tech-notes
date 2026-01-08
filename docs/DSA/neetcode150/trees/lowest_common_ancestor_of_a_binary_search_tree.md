---
tags:
  - BS Trees
  - LC_Medium
  - Neetcode150
hide:
  - toc
---
# 235. Lowest Common Ancestor of a Binary Search Tree

[Problem Link](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/){target=_blank}


Since it's a **Binary Search Tree (BST)**, we can leverage its property => for any node, all values in 
the left subtree are smaller, and all values in the right subtree are larger.

Starting from the root, we compare the values of nodes `p` and `q` with the current node:

- If both `p` and `q` have values smaller than the current node, then their lowest common ancestor must lie in the left
  subtree, so we move left.
- If both values are greater than the current node, the LCA must lie in the right subtree, so we move right.
- Otherwise, the current node is the split point where one node lies on one side and the other lies on the opposite
  side (or one of them is equal to the current node). This makes the current node the lowest common ancestor.

We continue this process iteratively until we find the split point, which is returned as the answer. This approach 
avoids unnecessary traversal and directly exploits the BST structure.

??? note "Pseudocode"
    ```text
    function LCA(root, p, q):
    while root is not null:
        if root.val > p.val AND root.val > q.val:
            root = root.left
        else if root.val < p.val AND root.val < q.val:
            root = root.right
        else:
            return root
    return null
    ```

??? note "Runtime Complexity"
    **Time**: $O(h)$, where $h$ is the height of the BST

    **Space**: $O(1)$, since the solution is iterative



=== "Python"

    ```python
    --8<-- "docs/DSA/src/py/lowest_common_ancestor_of_a_binary_search_tree.py:10"
    ```

=== "Go"

    ```go
    --8<-- "docs/DSA/src/go/lowest_common_ancestor_of_a_binary_search_tree.go:3"
    ```