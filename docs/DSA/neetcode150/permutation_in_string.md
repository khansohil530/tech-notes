---
tags:
  - LC_Medium
  - Sliding Window
  - Neetcode150
hide:
  - toc
---
# 567. Permutation in String

[Problem Link](https://leetcode.com/problems/permutation-in-string/description/){target=_blank}

We use a sliding window with two pointers $l$ and $r$ the current substring (in $s2$). The window will represent substring 
only consisting characters from $s1$. This way when our window reaches length of $s1$, we can say we've found the 
permutation. To grow/shrink our window, we need a hashmap storing frequency of $s1$. We'll use this map to consume 
characters as our window grow, and when the frequency reaches $\lt 0$, we can shrink the window. At any point, 
if $r-l+1 == len(s)$, we've found the permutation of $s1$.


??? note "Runtime Complexity"
    <b>Time</b>: O(n); each character is only processed once
    
    <b>Space</b>: O(1); constant space as only english letters used as keys in hashmap


=== "Python"

    ```python
    --8<-- "docs/DSA/src/py/permutation_in_string.py:4"
    ```

=== "Go"

    ```go
    --8<-- "docs/DSA/src/go/permutation_in_string.go:2"
    ```