---
tags:
  - Hash Table
  - Array
  - Heap/Priority Queue
  - LC_Medium
  - Neetcode150
  - Bucket Sort
---
# 347. Top K Frequent Elements

[Problem Link](https://leetcode.com/problems/top-k-frequent-elements/description/){target=_blank}

!!! info "Heap"
    This is a general problem designed to be solved efficiently using heap data structure.
    For now, we'll solve it using hashmaps.

The general idea to solve this is by generating a hashmap where the value is list of numbers and key is the frequency
of those number in `nums`. Now, you can use this hashmap to generate list of nums sorted in order of their frequency in
`nums`. But for this problem, we only want first $k$ elements from this list.

??? note "Pseudocode"
    - Generate hashmap where key is $num$ and value is frequency of that $num$ in `nums`.
    - Using above hashmap, generate another hashmap where key is a number and value is the list
      of numbers having $key$ frequency in `nums`.
    - Generate result having numbers sorted in order of highest frequency. 
      Since the highest frequency could be `len(nums)` while smallest -> $0$.
        - Iterate from the highest frequency to the smallest while adding numbers of
          frequencies found in our hashmap. 

??? note "Runtime Complexity"

    <b>Time</b>: $O(n)$, from iterating `nums`.
    
    <b>Space</b>: $O(n)$ <- from hashmaps.


=== "Python"
    ```python 
    --8<-- "docs/DSA/neetcode150/src/py/top_k_frequent_element.py:3:"
    ```
=== "Go"
    ```go
    --8<-- "docs/DSA/neetcode150/src/go/top_k_frequent_element.go:2:"
    ```
