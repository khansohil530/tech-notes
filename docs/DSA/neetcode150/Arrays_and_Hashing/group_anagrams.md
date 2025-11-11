---
tags:
  - Hash Table
  - Array
  - LC_Medium
  - Neetcode150
---
# :orange_circle: 49. Group Anagrams

[Problem Link](https://leetcode.com/problems/group-anagrams/description/){target=_blank}

You can check [this](valid_anagram.md) problem on how to find if two strings are anagrams.

If we map a string into an Array of size 26, where each element represents the occurrence
a character in a string. The associated character at any index in the Array is found by adding
the ASCII code of 'a' to the index -> $index+ascii(a)$. And we're using size 26, because the string
only have lowercase english characters which can be mapped to 26 positions in our array.

You can think of the generated Array as a kind of bitmap and strings which are anagram would've same bitmap. 
We can use this information to generate the group of Anagram string. 

??? note "Pseudocode"
    - Iterate over each string. During each iteration, 
      - Generate the bitmap of the string.
      - Store the string in a Hashmap where key is bitmap and value is list of string with same bitmap.
      - Finally returns the values of Hashmap, which would be grouped Anagram strings.

??? note "Runtime Complexity"
    `strs` -> length $n$, with string of maximum size $k$. 

    <b>Time</b>: $O(nk)$ <- $O(n)$ from iterating `strs`, $O(k)$ for generating bitmap of each string.
    
    <b>Space</b>: $O(n)$ <- from Hashmap for storing the groups.


=== "Python"
    ```python 
    --8<-- "docs/src/neetcode150/py/group_anagrams.py:2:"
    ```
=== "Go"
    ```go
    --8<-- "docs/src/neetcode150/go/group_anagrams.go:2:"
    ```

??? info "Implementation Note"
    **Python**: We're using tuple as key, because they're hashable object. If you use strings as key,
    consider the number of digits in counts when generating the string key.

    **Go**: keys to map must be `comparable` types. Slices (`[]int`) aren't comparablable but Arrays (`[N]T`) are
    comparable if their element Type (`T`) is comparable. 