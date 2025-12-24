---
tags:
  - Stack
  - LC_Easy
  - Neetcode150
hide:
  - toc
---
# 20. Valid Parentheses

[Problem Link](https://leetcode.com/problems/valid-parentheses/description/){target=_blank}

To check whether parentheses are valid, we need to ensure that every opening bracket is closed in the correct order. 
This can be done by pairing each closing bracket with the most recently opened bracket that has not yet been closed.
If the two brackets match, the pair is considered valid.

Because the last opened bracket must always be closed first, this process follows a Last-In-First-Out (LIFO) pattern. 
A stack is well suited for this task since it allows constant-time access to the most recent element.

??? note "Pseudocode"
    If the current character is an opening bracket:
    
    - Store it in a stack so it can be matched with a closing bracket later.
    - You can use a hashmap (for example, { "{": "}", "(": ")", "[": "]" }) to define valid pairs.

    If the current character is a closing bracket:

    - Check whether the stack is empty. If it is, there is no corresponding opening bracket, so the string is invalid.
    - Otherwise, compare the closing bracket with the expected closing bracket for the stack’s top element.
    - If they match, remove the top element from the stack (the pair is successfully closed).
    - If they do not match, the string is invalid.

    After processing all characters:

    - If the stack is empty, all brackets were properly closed.
    - If the stack still contains elements, some opening brackets were never closed.

??? note "Runtime Complexity"
    <b>Time</b>: $O(n)$, single pass + stack operation are constant time.

    <b>Space</b>: $O(n)$, stack can grow upto size of array.

=== "Python"

    ```python
    --8<-- "docs/DSA/neetcode150/src/py/valid_parentheses.py"
    ```

=== "Go"

    ```go
    --8<-- "docs/DSA/neetcode150/src/go/valid_parentheses.go:2"
    ```

??? info "Go Implementation"
    Time complexity of `stk = append(stk, val)` is amortized $O(1)$. Amortized because Go have to resize the slice when 
    its capacity is used, which is an $O(n)$ from copying and moving existing data. However, this happens infrequently
    making it $O(1)$ most of the time. 
    
    `stk = stk[:len(stk)-1]` is $O(1)$ since data is removed from end, Go just updates the header indicating length.