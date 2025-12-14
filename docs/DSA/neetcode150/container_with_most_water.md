---
tags:
  - Two Pointers
  - Greedy
  - LC_Medium
  - Neetcode150
---
# 11. Container With Most Water

[Problem Link](https://leetcode.com/problems/container-with-most-water/description/){target=_blank}

Area of container would be determined by two factors, 
- width 
- minimum height on either edges.

We can use two pointers starting from the largest width available. To iterate over to next candidate,
we can remove edges from either side. But we can **greedily** remove the shorter edge, as 
this would be always the limiting factor in area calculation -> $(r-l)*min(height[l],height[r])$.
Also, the solution from shorter edges has been already considered in current iteration, so we don't
have to worry about missing the current area in result.

??? note "Pseudocode"
    - Iterate over the array using two pointers, left and right.
    - During each iteration, calculate the area covered and update the global result if we've found larger value.
    - Later we can just shift the left or right pointers to move onto next candidate to find global maxima.

??? note "Runtime Complexity"
    <b>Time</b>: $O(n)$, since we're only iterating the nums array once.
    
    <b>Space</b>: $O(1)$, constant space from pointer variables.


=== "Python"

    ```python
    --8<-- "docs/DSA/neetcode150/src/py/container_with_most_water.py:2"
    ```

=== "Go"

    ```go
    --8<-- "docs/DSA/neetcode150/src/go/container_with_most_water.go:2"
    ```