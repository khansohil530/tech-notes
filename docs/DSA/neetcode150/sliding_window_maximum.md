---
tags:
  - LC_Hard
  - Sliding Window
  - Neetcode150
hide:
  - toc
---
# 239. Sliding Window Maximum

[Problem Link](https://leetcode.com/problems/sliding-window-maximum/description/){target=_blank}

You've to minimize the following operations:

- keep track of maximum values in window, 
- while also thinking of a way to invalidate the maximum value and update it as window shifts to right.

The easiest solution is to use a **Maximum Heap**, which would give you $O(1)$ access to max value while update/remove
would be $O(logn)$. Additionally, you can attach the index metadata in heap to invalidate any entry which 
are out of bounds. This would result in an $O(nlogn)$ time, and $O(n)$ space.

The optimal solution uses Monotonic Queue to achieve $O(n)$ space and time. You can maintain a monotonic queue whose
head will always point the maximum value of window. To achieve this, before adding any new value to queue you've to 
remove all values smaller than current so that it's the queue is sorted. For example, 
```text
nums=[1,3,-1,2]
q = [3] - add -1 -> q = [3, -1] - add 2 -> q = [3, 2]
```
After adding current value to queue, we've to check if the front value is within window. So we'll remove all front 
value outside window. Finally, the front value would indicate the maximum value within window. For example,
```text
nums=[1,3,-1,2], k = 2
window = [1,3], q = [3]
    | add -1
window = [3, -1], q = [3, -1]
    | add 2
window = [-1, 2], q = [3, 2] -> top invalid, pop front 
                             -> q = [2]
```


??? note "Runtime Complexity"
    <b>Time</b>: $O(n)$ constant time as each value is only processed once
                        and deque operations are $O(1)$

    <b>Space</b>: $O(n)$, queue can grow upto the size of $nums$


=== "Python"

    ```python
    --8<-- "docs/DSA/neetcode150/src/py/sliding_window_maximum.py:4"
    ```

=== "Go"

    ```go
    --8<-- "docs/DSA/neetcode150/src/go/sliding_window_maximum.go:2"
    ```