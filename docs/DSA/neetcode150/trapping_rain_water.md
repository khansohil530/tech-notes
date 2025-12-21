---
tags:
  - Two Pointers
  - LC_Hard
  - Neetcode150
hide:
  - toc
---
# 42. Trapping Rain Water

[Problem Link](https://leetcode.com/problems/trapping-rain-water/description/){target=_blank}

![Image Example](static/trapping_rain_water.png){loading=lazy width=400vw align=right}

Think of the total trapped water as sum of water collected at each index.
How would you calculate the trappable water at any index? The water level is bounded by the lesser of
maximum heights on left and right of given index. 

One way is to create prefix arrays storing maximum left and right heights for each index and finally calculate the 
water trapped at each level, but that’d add $O(n)$ space complexity. We can optimize the solution to O(1) extra space
using the two-pointer technique.

Start the iteration by declaring two pointers at left ($l$) and right ($r$) end of array, along with two variables to 
store maximum height seen so far from given $l$ ($maxL$) and $r$ ($maxR$). With this information, the amount of water 
that can be trapped at any possible is determined by: $water = min(maxL, maxR) - height[i]$. Now, we can move only the
pointer on the side with the smaller maximum height, because that side becomes the limiting factor. For example,

- $maxL < maxR$, then the left side limits the water level as the right side is guaranteed to have a boundary at least
  as tall $maxL$ (from $< maxR$), hence the water trapped at $l$ depends only on $maxL$.
  $\therefore water[l] = maxL - height[l]$ and we can move forward the $l$ as it's considered in out result
- $maxR <= maxL$, similarly here the right side limits the water level $\therefore water[r] = maxR - height[r]$.

??? note "Pseudocode"
    - Declare two pointers for indicating current left and right position in iteration.
    - Declare two variables to store the maximum left and right heights.
    - Iterate over the array until we've computed water level for all cells i.e $l<=r$.
    - For each iteration,
        - if $maxL < maxR$, we'll compute water level for `l` and increment the counter to next index
        - else, we'll computer water level for `r` and decrement the counter to next index.

??? note "Runtime Complexity"
    <b>Time</b>: $O(n)$, since we're only iterating the nums array once.
    
    <b>Space</b>: $O(1)$, constant time from two pointers.


=== "Python"

    ```python
    --8<-- "docs/DSA/neetcode150/src/py/trapping_rain_water.py:2"
    ```

=== "Go"

    ```go
    --8<-- "docs/DSA/neetcode150/src/go/trapping_rain_water.go:2"
    ```