---
tags:
  - Linked List
  - LC_Easy
  - Neetcode150
hide:
  - toc
---
# 21. Merge Two Sorted Lists

[Problem Link](https://leetcode.com/problems/merge-two-sorted-lists/description/){target=_blank}

To solve this in a single pass, we can use a dummy node along with a pointer `curr`.
At each step, we compare the current nodes of both lists and link the smaller one to `curr`.

After linking, we advance the pointer in the list from which the node was taken, and move `curr` forward.
Once one list is exhausted, we directly link the remaining nodes of the other list, since they are already sorted.
 


??? note "Runtime Complexity"
    <b>Time</b>: $O(n)$

    <b>Space</b>: $O(1)$


=== "Python"

    ```python
    --8<-- "docs/DSA/src/py/merge_two_sorted_lists.py:8"
    ```

=== "Go"

    ```go
    --8<-- "docs/DSA/src/go/merge_two_sorted_lists.go:2"
    ```