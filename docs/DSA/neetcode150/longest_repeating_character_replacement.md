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

We can use a sliding window with two pointers $l$ and $r$ representing the current substring. The goal is to find the
longest window that can be converted into a string of repeating characters using at most $k$ replacements.

For any valid window, the substring can be made of repeating character by switching the lower frequency characters, i.e.
$windowSize - max(counts) <= k$. But computing maximum count would give us $O(26)$ time (from 26 english letters). 
Instead, we can keep track of maximum repeating character in a variable ($maxf$). As we expand the window by moving $r$,
we update frequency map of characters and $maxf$

- if the window is valid, we can consider the window size in our result set $max(result, r-l+1)$.
- But if its invalid, i.e. $r-l+1-maxf > k$, we've to shrink window by moving $l$ forward and reducing respective count
  in hashmap.

!!! info "When shrinking, why do we update the frequency map but don’t decrease $maxf$?"
    This is safe because $maxf$ is allowed to be an upper bound rather than the exact maximum frequency in the current 
    window. An overestimated $maxf$ may temporarily allow an invalid window, but it not harm the optimal window size. 
    Since we’re only interested in the maximum window length, this approach remains correct while keeping the algorithm
    linear.
    ??? info "For example,"
        ```text
        s = "AABABBA", k = 1
        ---
        Index:  0 1 2 3 4 5 6
        Chars:  A A B A B B A
        ---
        r=0
        [A]            l=0
        maxf=1, size=1, subs=0 ✓
        
        r=1
        [A A]          l=0
        maxf=2, size=2, subs=0 ✓
        
        r=2
        [A A B]        l=0
        maxf=2, size=3, subs=1 ✓
        
        r=3
        [A A B A]      l=0
        maxf=3, size=4, subs=1 ✓   → result=4
        
        r=4
        [A A B A B]    l=0
        maxf=3, size=5, subs=2 ✗   → shrink
        
        remove s[l]=A, l=1
        [  A B A B]    l=1
        counts(A=2,B=2), maxf=3 (stale)
        size=4, subs=1 ✓   (kept)
        
        r=5
        [  A B A B B]  l=1
        counts(A=2,B=3), maxf=3
        size=5, subs=2 ✗   → shrink
        
        remove s[l]=A, l=2
        [    B A B B]  l=2
        size=4, subs=1 ✓ 
        ```


??? note "Runtime Complexity"
    <b>Time</b>: $O(n)$, from one pass
    
    <b>Space</b>: $O(n)$, from hashmap


=== "Python"

    ```python
    --8<-- "docs/DSA/neetcode150/src/py/longest_repeating_character_replacement.py:4"
    ```

=== "Go"

    ```go
    --8<-- "docs/DSA/neetcode150/src/go/longest_repeating_character_replacement.go:2"
    ```