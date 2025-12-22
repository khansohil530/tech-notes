---
tags:
  - Array
  - Hash Table
  - LC_Medium
  - Neetcode150
hide:
  - toc
---
# 36. Valid Sudoku

[Problem Link](https://leetcode.com/problems/valid-sudoku/description/){target=_blank}

To check if a sudo board is valid, you've to satisfy 3 types of criteria:
- each row shouldn't have repeating digit.
- each column shouldn't have repeating digit.
- each box (3x3 here) shouldn't have a repeating digit.

This can be checked pretty easily using hash maps (similar to [contains_dupliacte](contains_duplicate.md)).

??? question "How would you map cells within a box to same key?"
    $key=(floor(row/3))*3 + (floor(col/3))$

??? note "Pseudocode"
    - Iterate over the board using two loops. Within each iteration, if the cell isn't `.`
    - check if the value is present in your hashmap buckets.
        - if present, we can say the sudoku is invalid and return.
        - else add the values to respective hashmap key.

??? note "Runtime Complexity"
    number of elements in board -> n
        
    <b>Time</b>: $O(n)$, single iteration over each element with constant time checks.
    
    <b>Space</b>: $O(n)$ <- from hashmaps.


=== "Python"
    ```python
    --8<-- "docs/DSA/neetcode150/src/py/valid_sudoku.py:3"
    ```
=== "Go"
    ```go
    --8<-- "docs/DSA/neetcode150/src/go/valid_sudoku.go:2"
    ```