---
tags:
  - Two Pointers
  - LC_Medium
  - Neetcode150
hide:
  - toc
---
# 167. Two Sum II - Input Array Is Sorted

[Problem Link](https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/){target=_blank}

Similar to problem [TwoSum](two_sum.md), except the `nums` array is sorted now. 
We can greedily use this information by comparing sum of left and right end. If our sum exceed target,
we should reduce it by decreasing the right pointer. Else if our sum is smaller than target, we can increase
our current sum by increasing left pointer.


??? note "Runtime Complexity"
    <b>Time</b>: $O(n)$, since we're only iterating the nums array once.
    
    <b>Space</b>: $O(1)$, constant space from two pointers.


=== "Python"

    ```python
    --8<-- "docs/DSA/src/py/two_sum_2.py:2"
    ```

=== "Go"

    ```go
    --8<-- "docs/DSA/src/go/two_sum_2.go:2"
    ```