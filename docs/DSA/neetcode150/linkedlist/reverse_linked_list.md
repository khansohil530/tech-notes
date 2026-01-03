---
tags:
  - Linked List
  - LC_Easy
  - Neetcode150
hide:
  - toc
---
# 206. Reverse Linked List

[Problem Link](https://leetcode.com/problems/reverse-linked-list/description/){target=_blank}

To solve this in a single pass, we can use two pointers, `prev` and `curr`.
The `prev` pointer represents the previous node and will become the `next` pointer for the current node `curr`.
![reverse_linked_list.png](static/reverse_linked_list.png){loading=lazy width=600vw align=right}
While reversing the link for `curr`, we must first store its original `next` node in a temporary variable. 
This prevents losing the remaining part of the list and allows us to move forward to the next node.


??? note "Runtime Complexity"
    <b>Time</b>: $O(n)$

    <b>Space</b>: $O(1)$


=== "Python"

    ```python
    --8<-- "docs/DSA/src/py/reverse_linked_list.py:8"
    ```

=== "Go"

    ```go
    --8<-- "docs/DSA/src/go/reverse_linked_list.go:8"
    ```