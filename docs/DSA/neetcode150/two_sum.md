---
tags:
  - Hash Table
  - Array
  - LC_Easy
  - Neetcode150
hide:
  - toc
---
# 1. Two Sum

[Problem Link](https://leetcode.com/problems/two-sum/description/){target=_blank}

You can think of the problem as, 
**given a number $num$ find the index of $target-num$**.

This can be done by storing the index of each $num$ in a data structure like **Hashmap** which can be queried to
fetch index of a value in a constant time.

??? note "Pseudocode"
    - Iterate over the $nums$ array. For each $num$,
      - check if the $target-num$ is present in our map.
          - If yes, we can directly return the current index and saved index from here
          - Else, store the index of current num in our map and continue. 

??? note "Runtime Complexity"
    <b>Time</b>: $O(n)$, since we're only iterating the nums array once and querying hashmap is an $O(1)$ operation.
    
    <b>Space</b>: $O(n)$, due to the map used for storing num -> index mapping

=== "Python"
    ```python 
    --8<-- "docs/DSA/neetcode150/src/py/two_sum.py:2:"
    ```
=== "Go"
    ```go
    --8<-- "docs/DSA/neetcode150/src/go/two_sum.go:2:"
    ```