---
tags:
  - Trees
  - LC_Hard
  - Neetcode150
hide:
  - toc
---
# 297. Serialize and Deserialize Binary Tree

[Problem Link](https://leetcode.com/problems/serialize-and-deserialize-binary-tree/){target=_blank}

The goal is to convert a binary tree into a string and then reconstruct the same tree from that string. In practice, 
we can choose any traversal algorithm, but we use **preorder traversal** since it allows us to encounter the `root` 
first, which simplifies the implementation. Another important point is to explicitly encode `null` children so that the
tree structure is preserved.

Below solution uses an **iterative preorder traversal**. We traverse the tree using a stack. When a node is visited, we
record its value and push its right and left children onto the stack. If a node is `None`, we record a special marker
`"N"`. By doing this consistently, the serialized string uniquely represents both the values and structure of the tree.

During deserialization, we reverse this process. The serialized string is split using a delimiter, and the first value
becomes the root. While rebuilding the tree, each node needs to be processed for both its left and right children. 
Since preorder traversal visits the left subtree before the right subtree, we may encounter the same parent node at
different stages.

To handle this, we keep an additional state variable that tracks whether a node's left subtree has already been 
processed. If the left subtree is complete, we then proceed to build the right subtree. This allows us to reconstruct
the tree correctly while following the preorder sequence.

??? note "Runtime Complexity"
    **Time**: $O(n)$, where $n$ is the number of nodes in the tree

    **Space**: $O(n)$, for the serialized string and stack used during traversal


=== "Python"

    ```python
    --8<-- "docs/DSA/src/py/serialize_and_deserialize_binary_tree.py:7"
    ```

=== "Go"

    ```go
    --8<-- "docs/DSA/src/go/serialize_and_deserialize_binary_tree.go:8"
    ```