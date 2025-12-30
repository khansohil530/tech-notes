---
tags:
  - Binary Search
hide:
  - toc
---

# Binary Search

Binary Search allows you to eliminate search space by half in every step using ordering in dataset, essentially making
search $O(logn)$. At core, it finds the boundary point where a condition switches from false to true (or vice versa) to
eliminate half the search space.

When implementing Binary Search, it's rather difficult to write a bug-free code in just a few minutes. 
Some of the general decision which impacts the algorithm:

- When to exit the loop? Should we use $left < right$ or $left <= right$ as the while loop condition?
- How to initialize the boundary variable $left$ and $right$?
- How to update the boundary? How to choose the appropriate combination from $left = mid$ , $left = mid + 1$ and 
  $right = mid$, $right = mid - 1$?

To help with this, people on problem-solving platforms have developed the following generalized template which can
be used to apply binary search, with a little tweaking to the requirements of problem.

!!! note ""
    Suppose we have a search space which is sorted in ascending order. For most tasks, we can transform the 
    requirement into the following generalized form:
    
    minimize $k$ such that, $condition(k) = True$ 

**Template ->**

```python
def binary_search(array) -> int:
    def condition(value) -> bool:
        pass

    left, right = min(search_space), max(search_space) # could be [0, n], [1, n] etc. Depends on problem
    while left < right:
        mid = left + (right - left) // 2
        if condition(mid):
            right = mid
        else:
            left = mid + 1
    return left
```

You only need to modify three parts when using this template to solve most of the binary search problems, without 
worrying about corner cases and bugs:

- Correctly initialize the boundary variables `left` and `right` to specify search space by setting left and right 
  boundaries which includes all possible elements.
- Decide return value. Is it return `left` or return `left - 1`? Remember that after exiting the while loop, `left` is 
  the minimal $k$ satisfying the `condition` function.
- Design the `condition` function. Usually the most difficult part, but becomes easy with practice.

!!! note ""
    To compute $mid$, we can also do $(left+right)//2$ but addition can cause overflow in some languages. Therefore, it's
    better to compute $mid = left + (right-left)//2$.

## Pattern

Problems solvable by binary search have the following characteristics:

- Clear monotonic behavior in your result set, like `false false false true true true`, `valid -> invalid`, 
  `too small -> too large`.
- Search space is ordered (explicitly or implicitly), like a sorted array or answer ranges b/w a boundary ($1...10^9$) 
  and you've to find maximum, minimum or threshold of time, capacity, speed.

This commonly appears in following problems patterns:

1. Searching in arrays: Exact match, First/last occurrence, Insert position
2. Boundary problems: First element $\ge$ target, Last element $\le$ target
3. Search on answer problems: Minimum speed, Maximum capacity, Smallest valid value
4. Rotated / modified arrays: Rotated sorted array, Nearly sorted array, Infinite array (conceptually)

