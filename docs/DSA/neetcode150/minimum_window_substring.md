---
tags:
  - LC_Hard
  - Sliding Window
  - Neetcode150
hide:
  - toc
---
# 76. Minimum Window Substring

[Problem Link](https://leetcode.com/problems/minimum-window-substring/description/){target=_blank}

We need to find the minimum length substring in $s$, such that all characters from $t$ is present in the substring. 
To solve it we can use a sliding window ($l$, $r$ bounds), in which a valid window would indicate a substring with
all characters in $t$. To do this, 

- we can maintain a variable $missing = length(t)$ which would indicate the number of missing character from $t$ in 
  current window. 
- To make sure that $missing$ is only updated for needed character, we'll use a hashmap ($need$) which indicates the
  frequency of consumed characters in our window and populate it with counts from $t$.

With this setup, when we shift our window to right: 

1. Check if the current character is needed using hashmap, if so we can decrement it from $missing$.
2. Then we decrement the count of current character regardless, since this step would simplify the code for shrinking the
   window to correct position.
3. When our $missing$ variable reaches $0$, it means we've found all the characters from $t$ in our window $(l, r)$.
       - Since we've been ignoring the start of window, we've to fix it to correct position. This is now simple because 
         we're maintaining $need$ hashmap which would've negative frequency for overconsumed characters and $0$ count for
         needed characters. This approach also handles few subtleties like for example, $window = BCBA$ and $t = ABC$,
         here we'll safely ignore the first $B$ as its already account for in smaller  $window=CBA$.
       - Now the start of window is appropriate, we can check it against the optimal result and update.
       - Since we've consumed this window, we'll shrink the window from left to move forward.


??? note "Runtime Complexity"
    <b>Time</b>: $O(n)$ as each character is only processed once.
    
    <b>Space</b>: $O(1)$ from hashmap storing only english character keys 


=== "Python"

    ```python
    --8<-- "docs/DSA/neetcode150/src/py/minimum_window_substring.py"
    ```

=== "Go"

    ```go
    --8<-- "docs/DSA/neetcode150/src/go/minimum_window_substring.go:2"
    ```