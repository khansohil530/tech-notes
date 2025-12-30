---
tags:
  - Binary Search
  - LC_Medium
  - Neetcode150
hide:
  - toc
---
# 33. Search in Rotated Sorted Array

[Problem Link](https://leetcode.com/problems/search-in-rotated-sorted-array/description/){target=_blank}

This problem can be reduced into simple Binary search by solving it in two steps. 

Since the array is a rotated version of a sorted array, the first task is to find the rotation point, or $pivot$,
which is the index of the smallest element. This has been solved in [previous](find_minimum_in_rotated_sorted_array.md) 
problem.

Once the $pivot$ is found, the array can be viewed as two individually sorted subarrays, one from the $pivot$ to the end,
and one from the start to just before the $pivot$. At this point, the problem becomes a standard binary search.
You've to just determine which of these two sorted ranges could contain the $target$:

- if $nums[pivot] \le target \le nums[right]$, the $target$ is in right segment from pivot
- otherwise, it's in left segment from pivot.

Finally, run a normal binary search on the chosen segment to find the target.

??? note "Runtime Complexity"
    <b>Time</b>: $O(logn)$

    <b>Space</b>: $O(1)$


=== "Python"

    ```python
    --8<-- "docs/DSA/src/py/search_in_rotated_sorted_array.py"
    ```

=== "Go"

    ```go
    --8<-- "docs/DSA/src/go/search_in_rotated_sorted_array.go:2"
    ```