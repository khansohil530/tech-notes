---
tags:
  - Trees
  - LC_
  - Neetcode150
hide:
  - toc
---
# 1448. Count Good Nodes in Binary Tree

[Problem Link](https://leetcode.com/problems/count-good-nodes-in-binary-tree/){target=_blank}


A node is considered **good** if its value is greater than or equal to every value along the path from the root to that 
node. The main insight is that, at any point in the traversal, we only need to remember the **last good node** 
(i.e., the node with the maximum value seen so far on that path).

The solution uses a **recursive depth-first traversal**. The helper function takes two parameters:

- `node`: the current node being visited
- `lastGoodNode`: the node representing the maximum value encountered so far on the current path

At each step:

* If the current node is `nil`, we return `0` since it contributes no good nodes.
* If the current node’s value is greater than or equal to `lastGoodNode.Val`, it qualifies as a good node. We count 
  it and update `lastGoodNode` to the current node.
* We then recursively process the left and right subtrees, passing along the updated `lastGoodNode`.

By carrying the maximum-so-far information down the recursion, we ensure that each node is visited once and evaluated 
correctly without extra state or recomputation.

??? note "Pseudocode"
    ```
    function helper(node, lastGood):
    if node is null:
    return 0
        count = 0
        if node.val ≥ lastGood.val:
            count = 1
            lastGood = node
    
        return count 
               + helper(node.left, lastGood) 
               + helper(node.right, lastGood)
    ```

??? note "Runtime Complexity"
    **Time**: $O(n)$, where $n$ is the number of nodes in the tree

    **Space**: $O(h)$, where $h$ is the height of the tree due to recursion stack



=== "Python"

    ```python
    --8<-- "docs/DSA/src/py/count_good_nodes_in_binary_tree.py:6"
    ```

=== "Go"

    ```go
    --8<-- "docs/DSA/src/go/count_good_nodes_in_binary_tree.go:3"
    ```