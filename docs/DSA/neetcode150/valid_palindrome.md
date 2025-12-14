---
tags:
  - Two Pointers
  - LC_Easy
  - Neetcode150
---
# 125. Valid Palindrome

[Problem Link](https://leetcode.com/problems/valid-palindrome/description/){target=_blank}

Check Palindrome by comparing starting and ending characters in string. You can easily do this by 
taking two pointers, pointing the respective position from start and end which should be same to
form palindrome. 

??? note "Pseudocode"
    - Initialize start and end pointers
    - Iterate over the string using the pointers until they meet or cross each other over bound.

??? note "Runtime Complexity"
    <b>Time</b>: $O(n)$, single iteration over string.
    
    <b>Space</b>: $O(1)$, constant space by two pointers.


=== "Python"

    ```python
    --8<-- "docs/DSA/neetcode150/src/py/valid_palindrome.py:1"
    ```

=== "Go"

    ```go
    --8<-- "docs/DSA/neetcode150/src/go/valid_palindrome.go:5"
    ```

??? info "Go Implementation note"
    Strings in Go are sequence of bytes , so when you access any particular index - it'll return the respective byte.
    While characters are represented using `rune` which represents a Unicode code point.
    Byte occupies 1 byte size while Rune can use upto 4 bytes size, which is why converting a byte using `rune(.)`
    isn’t always safe. It's safe for ASCII characters which fall within 1 byte range.