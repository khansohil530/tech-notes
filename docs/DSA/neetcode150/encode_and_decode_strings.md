---
tags:
  - Hash Table
  - LC_Medium
  - Neetcode150
---
# 271. Encode and Decode Strings

[Problem Link](https://neetcode.io/problems/string-encode-and-decode?list=neetcode150){target=_blank}

One of the most common pattern for encoding any data structure into stream of continuous data type is
to use the size of data structure to denote the amount of data to read from stream for parsing single
unit of data. But, since our data might also include numbers there would be no way to differentiate the
size data from data structure value. To solve this, you can simply use a placeholder value between two
of these values. This would lead us to create streams as follows -> `<size><placeholder><data>...`

??? note "Pseudocode"
    **Encode**: Since our data structure is simply list of strings, we can

    - Calculate the size of each string
    - Generate the encoded token for this string -> `<size><placeholder><string>`
    - Join all the token strings to get encoded data.

    **Decode**: You can use two pointers to indicate the start and current position in stream.

    - We need to parse the stream until we reach end of it
    - Start by parsing the size of next data token, by reading data until you encounter the placeholder value.
    - Next decode the size into a number and read next `size` amount of data (skipping the placeholder).
    - Update the start and current pointer to end of current token and continue.

??? note "Runtime Complexity"
    <b>Time</b>: $O(n)$ <- encode, $O(n)$ <- decode
    
    <b>Space</b>: $O(n)$ <- encode (from storing tokens/output), $O(n)$ <- decode 


=== "Python"
    ```python
    --8<-- "docs/DSA/neetcode150/src/py/encode_and_decode_strings.py:2"
    ```
=== "Go"
    ```go
    --8<-- "docs/DSA/neetcode150/src/go/encode_and_decode_strings.go:10"
    ```