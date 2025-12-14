---
tags:
  - Two Pointers
  - LC_Medium
  - Neetcode150
---
# 15. 3Sum

[Problem Link](https://leetcode.com/problems/3sum/description/){target=_blank}

This problem is similar to [Two Sum](two_sum.md), we can just use an outer iteration
to reduce it to a two sum. But core part would be, **how you'd avoid duplicate triplets in result?**

One approach could be by sorting the array, and since duplicate values would be adjacent we could directly
skip them during each iteration. Also, Sorting Array is $O(nlogn)$ operation, which wouldn't impact the
runtime of our $O(n^2)$ solution.  

??? note "Pseudocode"
    - Sort the input Array and start an outer loop as an indicative of first number in triplet.
    - Skip the number if it's same as previous number, as we've already solved for it and don't want duplicate.
    - Within inner loop since the array is sorted, you can use [similar](two_sum_2.md) approach as Two Sum 2.

??? note "Runtime Complexity"
    <b>Time</b>: $O(n^2)$, from two loops.
    
    <b>Space</b>: $O(1)$/$O(n)$, depending on sorting algorithm


=== "Python"

    ```python
    --8<-- "docs/DSA/neetcode150/src/py/3sum.py:2"
    ```

=== "Go"

    ```go
    --8<-- "docs/DSA/neetcode150/src/go/3sum.go:4"
    ```