from typing import List


class Solution:
    def maxArea(self, height: List[int]) -> int:
        result = 0
        l, r = 0, len(height)-1
        while l < r:
            curr = (r-l)*min(height[l], height[r])
            result = max(result, curr)
            if height[r] > height[l]:
                l+=1
            else:
                r-=1
        return result