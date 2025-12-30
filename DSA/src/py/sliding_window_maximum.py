from typing import List
from collections import deque

class Solution:
    def maxSlidingWindow(self, nums: List[int], k: int) -> List[int]:
        q = deque()
        result = []
        l = 0
        for r, num in enumerate(nums):
            while q and nums[q[-1]] < num:
                _ = q.pop()

            q.append(r)
            while q and q[0] < l:
                _ = q.popleft()

            if r-l+1 == k:
                result.append(nums[q[0]])
                l+=1
        return result