---
tags:
  - Two Pointers
  - LC_Hard
  - Neetcode150
---
# :red_circle: 42. Trapping Rain Water

[Problem Link](https://leetcode.com/problems/trapping-rain-water/description/){target=_blank}

Think of the water amount as sum of water collected at each index.
How would you calculate the trappable water at any index? This water level is bounded by the minimum of
maximum heights on either side of given index.

![Image Example](../../../static/trapping_rain_water.png){loading=lazy}

With this, you can create such prefix arrays for maximum left and right heights of respective index and finally
calculate the water trapped at each level. This would give us $O(n)$ time and space complexity.

We can optimize this to use $O(1)$ space by using a two pointers technique.
For this, we'll need to initialize two pointers at left and right ends, and two variables storing the maximum
left and right heights of among iterated values.
For example, with `l` and `r` index pointers,  lets say maxL < maxR.
For this, we can say for sure that the water collected for l cell is bounded by maxL.
Because any consecutive height on right of `l` would be equal or greater than maxR which wouldn't impact
our computation of water -> $min(maxL, maxR)-height[i]$.

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
    --8<-- "docs/src/neetcode150/py/trapping_rain_water.py:2"
    ```

=== "Go"

    ```go
    --8<-- "docs/src/neetcode150/go/trapping_rain_water.go:2"
    ```