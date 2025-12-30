---
tags:
  - Two Pointers
  - Greedy
  - LC_Medium
  - Neetcode150
hide:
  - toc
---
# 11. Container With Most Water

[Problem Link](https://leetcode.com/problems/container-with-most-water/description/){target=_blank}

Area of container would be determined by two factors: -> **width** and -> **minimum height from its edges**.
To find the maximum area, we can iterate over different container ends using two loops and select the container with
maximum area value, but this approach uses $O(n^2)$ runtime.

Instead, we can **greedily** start the iteration from the large available width i.e. either ends of array using two
pointers. To move onto next optimal candidate, we can remove the shorter edge, as this would be always the limiting 
factor in area calculation -> $(r-l)*min(height[l],height[r])$. Also, the solution from shorter edges has been already 
considered in current iteration, so we don't have to worry about missing the current area in result.

??? note "Pseudocode"
    - Iterate over the array using two pointers, left and right.
    - During each iteration, calculate the area covered and update the global result if we've found larger value.
    - Later we can just shift the left or right pointers to move onto next candidate to find global maxima.

??? note "Runtime Complexity"
    <b>Time</b>: $O(n)$, since we're only iterating the nums array once.
    
    <b>Space</b>: $O(1)$, constant space from pointer variables.


=== "Python"

    ```python
    --8<-- "docs/DSA/src/py/container_with_most_water.py:2"
    ```

=== "Go"

    ```go
    --8<-- "docs/DSA/src/go/container_with_most_water.go:2"
    ```