---
tags:
  - Linked List
  - Two Pointers
  - LC_Medium
  - Neetcode150
hide:
  - toc
---
# 19. Remove Nth Node From End of List

[Problem Link](https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/){target=_blank}


To remove the $n^{th}$ node from the end of the list in a single pass, we use a **dummy node** and two pointers, 
`first` and `second`. The dummy node simplifies edge cases, such as when the head node needs to be removed.

Both pointers start at the dummy node. We first move the `second` pointer forward by `n` nodes, creating a fixed gap 
of `n` between `first` and `second`.

Next, we move both pointers together until `second` reaches the last node. At this point, `first` points to the node 
just before the one that needs to be removed. We then skip the target node by updating `first.next`.

Finally, we return `dummy.next` as the new head of the modified list.


??? note "Runtime Complexity"
    <b>Time</b>: $O(n)$

    <b>Space</b>: $O(1)$


=== "Python"

    ```python
    --8<-- "docs/DSA/src/py/remove_nth_node_from_end_of_list.py:7"
    ```

=== "Go"

    ```go
    --8<-- "docs/DSA/src/go/remove_nth_node_from_end_of_list.go:2"
    ```