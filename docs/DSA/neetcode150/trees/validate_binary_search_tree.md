---
tags:
  - BS Trees
  - DFS
  - LC_Medium
  - Neetcode150
hide:
  - toc
---
# 98. Validate Binary Search Tree

[Problem Link](https://leetcode.com/problems/validate-binary-search-tree/){target=_blank}


For a valid BST, **each node must satisfy constraints imposed by all of its ancestors**, not 
just its immediate parent. Simply comparing a node with its left and right children is not sufficient.

This can be done by using a DFS traversal with two additional parameters `lower_bound` and `upper_bound`, 
which are used to check validity of each node:

* `lower_bound`: the minimum value the node is allowed to take
* `upper_bound`: the maximum value the node is allowed to take

At each node:

* If the node is `None`, it is trivially valid.
* Otherwise, the node’s value must lie within the current bounds.
* When moving to the left subtree, the upper bound becomes `node.val - 1`, since all values must be strictly smaller.
* When moving to the right subtree, the lower bound becomes `node.val + 1`, since all values must be strictly larger.

??? note "Pseudocode"
    ```text
    function validate(node, upper, lower):
    if node is null:
    return true
    if node.val < lower OR node.val > upper:
        return false
    
    return validate(node.left, node.val - 1, lower)
           AND validate(node.right, upper, node.val + 1)
    ```

??? note "Runtime Complexity"
    **Time**: $O(n)$, where $n$ is the number of nodes in the tree
    
    **Space**: $O(h)$, where $h$ is the height of the tree due to recursion stack



=== "Python"

    ```python
    --8<-- "docs/DSA/src/py/validate_binary_search_tree.py:7"
    ```

=== "Go"

    ```go
    --8<-- "docs/DSA/src/go/validate_binary_search_tree.go:5"
    ```