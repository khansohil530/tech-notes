---
tags:
  - Stack
  - LC_Medium
  - Neetcode150
hide:
  - toc
---
# 150. Evaluate Reverse Polish Notation

[Problem Link](https://leetcode.com/problems/evaluate-reverse-polish-notation/description/){target=_blank}

Each operator requires the two most recent operands, and this requirement naturally chains as we combine unprocessed 
numbers with intermediate results. For example:
```text
tokens = ["1", "2", "3", "+", "-"]
           │    └─────────┘    │
           │        │ 2+3 = 5  │
           └───────────────────┘
                    │ 1-5 = -4
                     
```

The above sequence shows a last-in, first-out (LIFO) access pattern. Because of this, the problem can be solved using a 
stack while processing tokens from left to right. The stack stores both integer operands and intermediate results.

- Operand: If the current token is a number, push it onto the stack.
- Operator: If the token is an operator (+, -, *, /), pop the top two values from the stack, apply the operation in the 
  correct order (second operand operator first operand), and push the result back onto the stack.

Final Result: After all tokens are processed, the top of the stack contains the evaluated result.

??? note "Runtime Complexity"
    <b>Time</b>: $O(n)$, single pass over $tokens$ array

    <b>Space</b>: $O(n)$, $stk$ can grow upto size of $tokens$ array

=== "Python"

    ```python
    --8<-- "docs/DSA/src/py/evaluate_reverse_polish_notation.py"
    ```

=== "Go"

    ```go
    --8<-- "docs/DSA/src/go/evaluate_reverse_polish_notation.go:4"
    ```