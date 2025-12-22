---
tags:
  - LC_Medium
  - Sliding Window
  - Neetcode150
hide:
  - toc
---
# 424. Longest Repeating Character Replacement

[Problem Link](https://leetcode.com/problems/longest-repeating-character-replacement/description/){target=_blank}

We use a sliding window with two pointers $l$ and $r$ representing the current substring. The goal is to find the 
longest window that can be converted into a string of same characters using at most $k$ replacements.

For any window, the number of replacements needed is: $windowSize−max(character frequency in window)$. 
A window is valid if this value is at most $k$. Computing the maximum frequency from scratch on every iteration would 
cost $O(26)$ time (one per English letter), which is unnecessary.

Instead, we maintain a variable $maxf$ that stores the maximum frequency of any character seen so far while expanding 
the window. As $r$ moves right, we update the frequency map and $maxf$.

- If the window is valid ($r - l + 1 - maxf \le k$), we update the result with the current window size.
- If the window becomes invalid, we shrink it by moving $l$ forward and decreasing the frequency of $s[l]$.

When shrinking the window, the true maximum frequency inside the window may decrease, but we intentionally don’t
recompute $maxf$. Allowing $maxf$ to be an upper bound may temporarily keep an invalid window, but it never causes us
to miss the maximum valid window length. Using this formulation:

- window tracks the largest length achievable so far
- $maxf$ tracks the largest character frequency encountered so far

This avoids recomputing frequencies and ensures the algorithm runs in linear time $O(n)$.


??? info "For example"
    ```text
    s = "AABABBA", k = 1
    
    Index:  0 1 2 3 4 5 6
    Chars:  A A B A B B A
    
    ┌──────────────────────────┐
    │ Index:  0  1 2 3 4 5 6   │
    │ Chars: [A] A B A B B A   │
    ├──────────────────────────┤
    │l, r = 0, ws = 1 (r-l+1)  │
    │maxf = 1                  │
    │valid ✓ ws - maxf = 0<= 1 │
    └──────────────────────────┘
               │r++
    ┌──────────────────────────┐
    │ Index:  0 1  2 3 4 5 6   │
    │ Chars: [A A] B A B B A   │
    ├──────────────────────────┤
    │l=0, r = 1, ws = 2        │
    │maxf = 2                  │
    │valid ✓ ws - maxf = 0<= 1 │
    └──────────────────────────┘
               │r++
    ┌──────────────────────────┐
    │ Index:  0 1 2  3 4 5 6   │
    │ Chars: [A A B] A B B A   │
    ├──────────────────────────┤
    │l=0, r = 2, ws = 3        │
    │maxf = 2                  │
    │valid ✓ ws - maxf = 1<= 1 │
    └──────────────────────────┘
               │r++
    ┌──────────────────────────┐
    │ Index:  0 1 2 3  4 5 6   │
    │ Chars: [A A B A] B B A   │
    ├──────────────────────────┤
    │l=0, r = 3, ws = 4        │
    │maxf = 3                  │
    │valid ✓ ws - maxf = 1<= 1 │
    └──────────────────────────┘
               │r++
    ┌──────────────────────────┐
    │ Index:  0 1 2 3 4  5 6   │
    │ Chars: [A A B A B] B A   │
    ├──────────────────────────┤
    │l=0, r = 4, ws = 5        │
    │maxf = 3                  │<----- Important
    │valid ✗ ws - maxf = 2 > 1 │     Need to shrink
    └──────────────────────────┘
               │ l++
    ┌──────────────────────────┐
    │ Index: 0  1 2 3 4  5 6   │
    │ Chars: A [A B A B] B A   │
    ├──────────────────────────┤
    │l=1, r = 4, ws = 4        │
    │maxf = 3 <-------------------- this must be 2 as per our given window but we 
    │valid ✓ ws - maxf = 1<= 1 │    can avoid this without impacting window size
    └──────────────────────────┘
               │r++                 
               │                
    ┌──────────────────────────┐    
    │ Index: 0  1 2 3 4 5  6   │     
    │ Chars: A [A B A B B] A   │    
    ├──────────────────────────┤      
    │l=1, r = 5, ws = 5        │
    │maxf = 3                  │
    │valid ✗ ws - maxf = 2 > 1 │
    └──────────────────────────┘
               │ l++
    ┌──────────────────────────┐
    │ Index: 0 1  2 3 4 5  6   │
    │ Chars: A A [B A B B] A   │
    ├──────────────────────────┤
    │l=2, r = 5, ws = 4        │
    │maxf = 3                  │
    │valid ✓ ws - maxf = 1<= 1 │
    └──────────────────────────┘
               │r++
    ┌──────────────────────────┐
    │ Index: 0 1  2 3 4 5 6    │
    │ Chars: A A [B A B B A]   │
    ├──────────────────────────┤
    │l=2, r = 6, ws = 5        │
    │maxf = 3                  │
    │valid ✗ ws - maxf = 2> 1  │
    └──────────────────────────┘
               │l++
    ┌──────────────────────────┐
    │ Index: 0 1 2  3 4 5 6    │
    │ Chars: A A B [A B B A]   │
    ├──────────────────────────┤
    │l=2, r = 6, ws = 4        │
    │maxf = 3                  │
    │valid ✓ ws - maxf = 1<= 1 │
    └──────────────────────────┘
    ```


??? note "Runtime Complexity"
    <b>Time</b>: $O(n)$, from one pass
    
    <b>Space</b>: $O(1)$, constant space as only english letters used as keys in hashmap


=== "Python"

    ```python
    --8<-- "docs/DSA/neetcode150/src/py/longest_repeating_character_replacement.py:4"
    ```

=== "Go"

    ```go
    --8<-- "docs/DSA/neetcode150/src/go/longest_repeating_character_replacement.go:2"
    ```