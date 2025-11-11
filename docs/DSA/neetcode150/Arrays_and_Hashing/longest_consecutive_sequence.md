---
tags:
  - Array
  - Hash Table
  - Union Find
  - LC_Medium
  - Neetcode150
---
# :orange_circle: 128. Longest Consecutive Sequence



[Problem Link](https://leetcode.com/problems/longest-consecutive-sequence/description/){target=_blank}

**Brute Force**: Since we only need to find the longest consecutive sequence irrespective of order in array,
we can check the sequence starting each element and take the longest. This would result in $O(n^2)$ time 
complexity.

To optimize this, we can just check the sequence for $num_i$ whose previous number ($num_i-1$) isn't present in our 
array, as these number are guaranteed to generate unique sequence . This would result in checking the sequence for
elements only once. 

??? note "Pseudocode"
    - Iterate over the nums array. To keep checks constant time, we can also create a Hashset out of `nums` array.
    - For each `num`, if the previous (`num-1`) isn't present in our array means we've a unique sequence starting 
      from this `num`.
    - To generate this sequence, declare a length pointer and increment it until we exhaust the sequence numbers
      present in our hashset.
    - Finally use a global variable to maintain the maximum length.

??? note "Runtime Complexity"
    <b>Time</b>: $O(n)$, single iteration over each element with constant time checks.
    It might look like we're iterating within 2 loops, but the conditional will reduce the
    inner iteration such that we're not repeating checks.
    
    <b>Space</b>: $O(n)$ <- from hashmap.


=== "Python"
    ```python
    --8<-- "docs/src/neetcode150/py/longest_consecutive_sequence.py:3"
    ```
=== "Go"
    ```go
    --8<-- "docs/src/neetcode150/go/longest_consecutive_sequence.go:2"
    ```