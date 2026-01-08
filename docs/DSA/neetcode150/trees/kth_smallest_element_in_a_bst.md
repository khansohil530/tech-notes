---
tags:
  - BS Trees
  - LC_Medium
  - Neetcode150
hide:
  - toc
---
# 230. Kth Smallest Element in a BST

[Problem Link](https://leetcode.com/problems/kth-smallest-element-in-a-bst/){target=_blank}

**Inorder traversal** on a BST visits nodes in ascending order**. This property allows us to directly find the k-th 
smallest element, but think about how to do it without storing all values.

We perform an inorder traversal (left → node → right) while maintaining a counter `k`. Each time we visit a node, we
decrement `k`. When `k` reaches zero, the current node’s value is exactly the k-th smallest element.

When implementing this recursively, you can use global variables for `k` and `res` which would capture result relatively
easy. In below code, the helper function implements Inorder traversal and when we reach $k=0$, we'll simply capture that
node instead of worrying about returning the values back up to recursion. By leveraging the BST property and stopping 
as soon as the answer is found, we avoid unnecessary work and extra space.


??? note "Runtime Complexity"
    **Time**: $O(h + k)$ in the best case, $O(n)$ in the worst case

    **Space**: $O(h)$, where $h$ is the height of the tree due to recursion stack


=== "Python"

    ```python
    --8<-- "docs/DSA/src/py/kth_smallest_element_in_a_bst.py:7"
    ```

=== "Go"

    ```go
    --8<-- "docs/DSA/src/go/kth_smallest_element_in_a_bst.go:3"
    ```