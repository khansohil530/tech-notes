---
tags:
  - Stack
  - LC_Medium
  - Neetcode150
hide:
  - toc
---
# 155. Min Stack

[Problem Link](https://leetcode.com/problems/min-stack/description/){target=_blank}

Use two stacks: one stack ($storeStk$) to hold all the pushed values, and another stack ($minStk$) to track only the 
candidate minimum values, with the current minimum always at the top.

- Push: Push the value into $storeStk$. Also push it into $minStk$ if $minStk$ is empty or the value is less than or
  equal to the current top of $minStk$.
- Pop: Pop the top element from $storeStk$. If the popped value is equal to the top of $minStk$, pop from $minStk$ as
  well.
- GetMin: Return the top element of $minStk$.


??? note "Runtime Complexity"
    <b>Time</b>: $O(1)$

    <b>Space</b>: $O(n)$

=== "Python"

    ```python
    --8<-- "docs/DSA/neetcode150/src/py/min_stack.py"
    ```

=== "Go"

    ```go
    --8<-- "docs/DSA/neetcode150/src/go/min_stack.go:2"
    ```