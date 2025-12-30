---
tags:
  - Array
  - Prefix Sum
  - LC_Medium
  - Neetcode150
hide:
  - toc
---
# 238. Product of Array Except Self

[Problem Link](https://leetcode.com/problems/product-of-array-except-self/description/){target=_blank}

The catch of problem is we can't use division operator. If that wasn't the case, we could simply have
aggregated multiplication of whole array and divide each element by it get their own product without self.


When translated to equations, this would look like: $\prod_{i=1}^{n} nums_i \over nums_i$. By removing the division
 -> $\prod_{i=1}^{i-1} nums_i * \prod_{i=i+1}^{n} nums_i$, which is basically product of prefix and postfix array
to index `i`.

One way to implement this is to store prefix and postfix multiplication in two separate arrays and combine them
to generate the result. 

!!! tip "Follow Up"
    Instead of saving the prefix and postfix multiplication within separate array, we can simply store them
    within output array (since multiplication is communicative) 

??? note "Pseudocode"
    - To generate prefix, iterate from start of `nums` and using a variable to aggregate the multiplication over
      the iteration.
    - To generate postfix, iterate from end of `nums` and repeat same.

??? note "Runtime Complexity"
    <b>Time</b>: $O(n)$
    
    <b>Space</b>: $O(1)$ <- output isn't considered


=== "Python"
    ```python
    --8<-- "docs/DSA/src/py/product_of_array_except_self.py:2"
    ```
=== "Go"
    ```go
    --8<-- "docs/DSA/src/go/product_of_array_except_self.go:2"
    ```