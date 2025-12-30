---
tags:
  - Stack
  - Monotonic Stack
  - LC_Medium
  - Neetcode150
hide:
  - toc
---
# 739. Daily Temperatures

[Problem Link](https://leetcode.com/problems/daily-temperatures/description/){target=_blank}

For each day, the goal is to find the next future day with a higher temperature. Instead of using a brute-force 
approach with two nested loops, we can resolve multiple previously colder days in a single pass by using a stack.

As we iterate through the days, we store the indices of unresolved days in a stack. Before pushing the current day, we 
check whether it is warmer than the days represented at the top of the stack. While the stack is not empty and the 
current temperature is higher, we pop those indices and compute how many days they had to wait.

This process maintains the stack in a monotonically decreasing order of temperatures, ensuring that all colder days are
processed as soon as a warmer day appears and that no unresolved day is missed.

??? note "Pseudocode"

    - Process the temperature list from left to right while maintaining a stack that stores indices of days whose warmer
      temperature has not yet been found.
    - For the current day $r$ with temperature $temp$, while the stack is not empty and the last recorded temperature is
      less than $temp$, we've found the next warmer day for that last temperature.
        - Pop the index $l$ from the stack and set $result[l] = r - l$, which is the number of days waited.
    - After resolving all smaller temperatures, push the current temperature index $r$ onto the stack.
    - Any indices left in the stack do not have a warmer future day, so their result remains 0.

??? note "Runtime Complexity"
    <b>Time</b>: $O(n)$, since we only iterate temperatures once, and the stack push/pop operation are constant time

    <b>Space</b>: $O(n)$, stack can grow upto size of temperatures

=== "Python"

    ```python
    --8<-- "docs/DSA/src/py/daily_temperatures.py"
    ```

=== "Go"

    ```go
    --8<-- "docs/DSA/src/go/daily_temperatures.go:2"
    ```