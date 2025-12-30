---
tags:
  - Binary Search
  - LC_Medium
  - Neetcode150
hide:
  - toc
---
# 875. Koko Eating Bananas

[Problem Link](https://leetcode.com/problems/koko-eating-bananas/description/){target=_blank}

We need to find the minimum eating rate $k$ such that Koko can eat all the bananas in $piles$ within $h$ hours.
The optimal value of $k$ lies between two bounds:

- Minimum possible eating rate, where Koko must eat at least $\lceil \sum(piles)/h \rceil$ bananas per hour otherwise, 
  even with perfect distribution, she cannot finish all bananas in $h$ hours
- Maximum possible eating rate, $\max piles$ since at this rate, Koko can finish each pile in at most one hour, and 
  therefore all piles within $h$ hours.

As we arrange the eating rates from the minimum to the maximum, we observe a monotonic pattern.
For smaller rates, Koko cannot finish all bananas within $h$ hours, and once a rate becomes valid, all larger rates 
remain valid. For example,
```text
piles = [3, 6, 7, 11], h = 8

minimum rate = ceil(sum(piles) / h) = 4
maximum rate = max(piles) = 11

Let k be the eating rate, and hrs be the total hours required at rate k

k:   1   2   3   | 4   5   6   7   8   9  10  11
hrs:27  14  10   | 8   8   6   5   5   5   5   4
ok?: ✘   ✘   ✘   | ✔   ✔   ✔   ✔   ✔   ✔   ✔   ✔
                 ↑
           transition point

```

This single transition point, from invalid to valid eating rates, allows us to efficiently search for the minimum valid
$k$ using binary search.

??? note "Runtime Complexity"
    <b>Time</b>: $O(nlogM)$, where $M = max(piles)$. The log $M$ factor comes from binary search over the eating rates, 
    and each check takes $O(n)$ time to compute the total hours.

    <b>Space</b>: $O(1)$, constant variables


=== "Python"

    ```python
    --8<-- "docs/DSA/src/py/koko_eating_bananas.py"
    ```

=== "Go"

    ```go
    --8<-- "docs/DSA/src/go/koko_eating_bananas.go:2"
    ```