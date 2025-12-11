---
tags:
  - Hash Table
  - Array
  - LC_Easy
  - Neetcode150
---
# :green_circle: 217. Contains Duplicate

[Problem Link](https://leetcode.com/problems/contains-duplicate/description/){target=_blank}

**Brute Force**: You can iterate twice over the Array to check if element from outer iteration is present using inner 
iteration. This would yield following runtime complexity: Time -> $O(n^2)$, Space -> $O(1)$

We can optimize the operation of checking if an element is present from $O(n)$ to $O(1)$ using a data structure like 
**HashSet** which gives us constant time querying operation. The tradeoff being increase in space complexity to
$O(n)$, which majority of time is acceptable as memory isn't as scare as old days.

??? note "Pseudocode"
    - Iterate over the $nums$ array. For each $num$,
      - check if the $num$ is present in our hashset. 
          - If yes, we can return immediately
          - Else, store current $num$ in our hashset and continue. 

??? note "Runtime Complexity"
    <b>Time</b>: $O(n)$, since we're only iterating the nums array once and querying hashset, an $O(1)$ operation.
    
    <b>Space</b>: $O(n)$, due to the hashset used for storing num.


=== "Python"

    ```python
    --8<-- "docs/DSA/neetcode150/src/py/contains_duplicate.py:2"
    ```

=== "Go"

    ```go
    --8<-- "docs/DSA/neetcode150/src/go/contains_duplicate.go:2"
    ```