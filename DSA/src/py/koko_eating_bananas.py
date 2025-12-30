import math

class Solution:
    def calcTime(self, piles: list[int], k: int) -> int:
        t = 0
        for pile in piles:
            t += math.ceil(pile/k)
        return t

    def minEatingSpeed(self, piles: list[int], h: int) -> int:
        right, left = max(piles), math.ceil(sum(piles)/h)
        while left < right:
            mid = left + (right-left)//2
            if self.calcTime(piles, mid) <= h:
                right = mid
            else:
                left = mid + 1

        return left