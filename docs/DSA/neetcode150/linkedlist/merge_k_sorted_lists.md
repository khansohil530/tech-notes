---
tags:
  - Linked List
  - Heap
  - Merge Sort
  - LC_Hard
  - Neetcode150
hide:
  - toc
---
# 23. Merge k Sorted Lists

[Problem Link](https://leetcode.com/problems/merge-k-sorted-lists/description/){target=_blank}

We've previously solved [Merging two sorted linked lists](merge_two_sorted_lists.md) efficiently in linear time.
However, this problem requires merging `k` sorted lists, and a naive approach of merging them one by one would lead to
unnecessary repeated work and higher time complexity. Instead, we can apply a divide-and-conquer strategy, similar to 
the merge step in merge sort. Instead of merging all lists sequentially, we merge them in pairs, reducing the number of 
lists by half at each step.

Initially, all `k` lists are stored in an array. While more than one list remains:

- We iterate through the array in steps of two.
- Merge each adjacent pair of lists.
- Store the merged results in a temporary array.

If the number of lists is odd, the last list is carried forward as-is. After each pass, the total number of lists is 
reduced roughly by half. This process continues until only one merged list remains, which is returned as the final 
result.

Each node is processed once per merge level, and the number of merge levels is $O(logk)$. This ensures the total time
complexity stays optimal.

??? note "Runtime Complexity"
    <b>Time</b>: $O(Nlogk)$, where $N$ is the total number of nodes across all lists, $k$ is number of lists

    <b>Space</b>: $O(k)$


=== "Python"

    ```python
    --8<-- "docs/DSA/src/py/merge_k_sorted_lists.py"
    ```

=== "Go"

    ```go
    --8<-- "docs/DSA/src/go/merge_k_sorted_lists.go:2"
    ```