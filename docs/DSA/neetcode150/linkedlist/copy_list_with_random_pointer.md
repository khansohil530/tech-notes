---
tags:
  - Linked List
  - LC_Medium
  - Neetcode150
hide:
  - toc
---
# 138. Copy List with Random Pointer

[Problem Link](https://leetcode.com/problems/copy-list-with-random-pointer/description/){target=_blank}

To create a deep copy of the linked list, we use a **hash map** to maintain a one-to-one mapping between each original
node and its cloned node.

In the first pass, we iterate through the original list and create a new node for each existing node, storing the 
mapping from the original node to its clone in the hash map. At this stage, only the node values are copied.

In the second pass, we iterate through the list again and assign the `Next` and `Random` pointers for each cloned node
by looking them up in the hash map. This ensures that both pointers correctly reference the corresponding cloned nodes.

Finally, we return the cloned head node obtained from the hash map, which represents the deep-copied list.

??? note "Runtime Complexity"
    <b>Time</b>: $O(n)$

    <b>Space</b>: $O(n)$


=== "Python"

    ```python
    --8<-- "docs/DSA/src/py/copy_list_with_random_pointer.py:8"
    ```

=== "Go"

    ```go
    --8<-- "docs/DSA/src/go/copy_list_with_random_pointer.go:8"
    ```