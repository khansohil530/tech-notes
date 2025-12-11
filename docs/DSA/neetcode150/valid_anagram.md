---
tags:
  - Hash Table
  - LC_Easy
  - Neetcode150
---
# :green_circle: 242. Valid Anagram

[Problem Link](https://leetcode.com/problems/valid-anagram/description/){target=_blank}

Given strings `s` and `t`, they're Anagram if we can rearrange the characters of one to form another.
This can be reworded as, **both strings should have same characters**. We can check this by storing the
frequency of one string and match them with the other.  

??? note "Pseudocode"
    - If strings have different number of characters, we can return early as they'll never have same characters.
      - Otherwise, we'll iterate over one string and store the frequency of each character in a hashmap.
      - Iterate over other string to reduce the frequency of encountered character.
    - During which, we can say the strings aren’t anagrams if we encounter any character which either isn't present 
                in our hashmap, or the frequency of character is exhausted.

??? note "Runtime Complexity"
    <b>Time</b>: $O(n)$, since we're only iterating the strings twice, with insertion and querying hashmap being $O(1)$
                operation.
    
    <b>Space</b>: $O(n)$, due to the hashmap used for storing characters frequency.


=== "Python"
    ```python
    --8<-- "docs/DSA/neetcode150/src/py/valid_anagram.py"
    ```
=== "Go"
    ```go
    --8<-- "docs/DSA/neetcode150/src/go/valid_anagram.go:2"
    ```