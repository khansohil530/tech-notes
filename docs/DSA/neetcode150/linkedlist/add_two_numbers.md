---
tags:
  - Linked List
  - LC_Medium
  - Neetcode150
hide:
  - toc
---
# 2. Add Two Numbers

[Problem Link](https://leetcode.com/problems/add-two-numbers/description/){target=_blank}


To add the two numbers represented by linked lists, we traverse both lists simultaneously while maintaining a `carry`
for digit overflow.

We use a **dummy node** and a pointer `curr` to build the result list incrementally. At each step, we compute the sum 
of the current digits from `l1` and `l2`, along with the carried value from the previous step.

If a list is exhausted, we simply skip adding its value. We then create a new node with the digit `total % 10` and 
update the carry as `total // 10`.

After processing both lists, if a carry remains, we append a final node to the result list. The result is returned 
starting from `dummy.next`, which represents the sum in reverse order.


??? note "Runtime Complexity"
    <b>Time</b>: $O(n)$ 

    <b>Space</b>: $O(1)$ if output isn't considered, $O(n)$ if output is considered


=== "Python"

    ```python
    --8<-- "docs/DSA/src/py/add_two_numbers.py:6"
    ```

=== "Go"

    ```go
    --8<-- "docs/DSA/src/go/add_two_numbers.go:2"
    ```