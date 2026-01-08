---
tags:
  - Trees
  - DFS
  - ZFunction
  - LC_Easy
  - Neetcode150
hide:
  - toc
---
# 572. Subtree of Another Tree

[Problem Link](https://leetcode.com/problems/subtree-of-another-tree/){target=_blank}

To check whether one tree is a subtree of another, the main challenge is comparing both **structure and values** 
efficiently. A direct recursive comparison at every node can work, but it's inefficient to check if both subtrees are
the same tree for every matching tree in the worst case. Instead, we can hash both trees in string and reframe the
problem as a **string-matching** task.

You can serialize both trees as a string using any standard traversal algorithms, just make sure the `null` child nodes
are noted in serialization Including null markers is crucial without them, different tree structures could produce the 
same string.

Once both trees are serialized, the problem reduces to checking whether the serialized `subRoot` string appears as a 
substring within the serialized `root` string.

To perform this check efficiently, we use the [**Z-algorithm**](../algos/zfunction.md). We concatenate the strings as
`subRoot_serialized + "|" + root_serialized` and compute the Z-array. The Z-value at a position tells us how many 
characters match the prefix starting from that index.

If at any position in the concatenated string the Z-value equals the length of the serialized `subRoot`, we have found 
an exact match, meaning the subtree exists.

This approach ensures that both tree structure and node values are matched exactly, while achieving linear-time string 
matching.

??? note "Pseudocode"
    ```text
    function isSubtree(root, subRoot):
        s = serialize(subRoot)
        t = serialize(root)
        z = Z_function(s + "|" + t)
        for i from len(s)+1 to end:
            if z[i] == len(s):
                return true
        return false
    ```

??? note "Runtime Complexity"
    **Time**: $O(n + m)$, where $n$ and $m$ are the number of nodes in `root` and `subRoot`

    **Space**: $O(n + m)$ for serialization and Z-array

=== "Python"

    ```python
    --8<-- "docs/DSA/src/py/subtree_of_another_tree.py:7"
    ```

=== "Go"

    ```go
    --8<-- "docs/DSA/src/go/subtree_of_another_tree.go:5"
    ```