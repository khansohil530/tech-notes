---
tags:
  - Binary Search
  - LC_Easy
  - Neetcode150
hide:
  - toc
---
# 704. Binary Search

[Problem Link](https://leetcode.com/problems/binary-search/description/){target=_blank}

Using the [general](index.md) template,

- The search space is $[0,len(nums))$, so we initialize $left = 0$ and $right = len(nums)$.
- The condition is to find the minimal index $k$ such that $nums[k] >= target$.
- If the target exists in the array, the minimal value satisfying the condition will be `target` itself, and its 
  index will be stored in `left`.

??? note "Runtime Complexity"
    <b>Time</b>: $O(logn)$, search space is reduced by half in each iteration

    <b>Space</b>: $O(1)$, constant space from variables.


=== "Python"

    ```python
    --8<-- "docs/DSA/src/py/binary_search.py"
    ```

=== "Go"

    ```go
    --8<-- "docs/DSA/src/go/binary_search.go:2"
    ```