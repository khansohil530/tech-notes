---
tags:
  - Binary Search
  - LC_Medium
  - Neetcode150
hide:
  - toc
---
# 153. Find Minimum in Rotated Sorted Array

[Problem Link](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/description/){target=_blank}

The array is sorted in ascending order and then rotated, which means it consists of two sorted parts and the minimum 
element is the point where the rotation happens. A linear scan would work, but the sorted structure allows a more 
efficient binary search approach.

With binary search, comparing $nums[mid]$ with $nums[right]$ tells which side contains the minimum. 

- If $nums[mid] \lt nums[right]$, the right half from $mid$ to $right$ is sorted, so the minimum must be at $mid$ or
  to its left.
- Otherwise, the minimum lies strictly to the right of $mid$

This process shrinks the search space while always keeping the minimum within the current bounds. When the loop ends, 
$left$ points to the minimum element.

??? note "Runtime Complexity"
    <b>Time</b>: $O(logn)$

    <b>Space</b>: $O(1)$


=== "Python"

    ```python
    --8<-- "docs/DSA/src/py/find_minimum_in_rotated_sorted_array.py"
    ```

=== "Go"

    ```go
    --8<-- "docs/DSA/src/go/find_minimum_in_rotated_sorted_array.go:2"
    ```