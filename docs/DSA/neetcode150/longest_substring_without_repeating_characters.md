---
tags:
  - LC_Medium
  - Sliding Window
  - Neetcode150
hide:
  - toc
---
# 3. Longest Substring Without Repeating Characters

[Problem Link](https://leetcode.com/problems/longest-substring-without-repeating-characters/description/){target=_blank}

Similar to [last problem](best_time_to_buy_and_sell_stock.md), try to think of way to minimize windows from double loops.

We'll use a sliding window with two pointers $l$ and $r$ such that the window always contains unique characters. 
We expand the window by moving r from left to right and fix it whenever it becomes invalid.
When processing $s[r]$:

- If $s[r]$ hasn’t appeared in the current window, we include it and update the maximum window length.
- If $s[r]$ has appeared before, we move $l$ to one position after its last occurrence to restore uniqueness.

To do this efficiently, we store each character’s most recent index in a hashmap. There’s no need to explicitly remove 
characters between the old $l$ and the new $l$ as advancing $l$ implicitly invalidates them, and the hashmap entries are 
overwritten as we continue scanning. 

??? note "Runtime Complexity"
    <b>Time</b>: $O(n)$, each character is only processed once
    
    <b>Space</b>: $O(n)$, from hashmap


=== "Python"

    ```python
    --8<-- "docs/DSA/neetcode150/src/py/longest_substring_without_repeating_characters.py"
    ```

=== "Go"

    ```go
    --8<-- "docs/DSA/neetcode150/src/go/longest_substring_without_repeating_characters.go:2"
    ```