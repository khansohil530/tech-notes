---
tags:
  - Binary Search
  - LC_Medium
  - Neetcode150
hide:
  - toc
---
# 981. Time Based Key-Value Store

[Problem Link](https://leetcode.com/problems/time-based-key-value-store/description/){target=_blank}

We can store and retrieve standard key–value pairs in a hashmap in $O(1)$. However, this problem requires searching for 
a given $key$ and $timestamp$, the value stored at the largest timestamp $timestampPrev \le timestamp$.

The $key$ lookup can still be done in $O(1)$ using a hashmap. Since timestamps for each $key$ are monotonically 
increasing, we can store the corresponding values (along with timestamps) in an array. This allows us to use binary 
search to find the required value in $O(logn)$.

Using the [standard binary search template](index.md), we find the minimum $index$ such that $timestamp \lt arr[index]$. 
The answer is then at $index−1$, since we want the largest timestamp that is less than or equal to the given 
$timestamp$. 

??? note "Runtime Complexity"
    <b>Time</b>: $O(1)$ for Set and $O(logn)$ for Get

    <b>Space</b>:$O(mn)$, where $m$ is the number of keys and $n$ is the average number of values per key



=== "Python"

    ```python
    --8<-- "docs/DSA/src/py/time_based_key_value_store.py:2"
    ```

=== "Go"

    ```go
    --8<-- "docs/DSA/src/go/time_based_key_value_store.go:2"
    ```