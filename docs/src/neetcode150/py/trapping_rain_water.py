from typing import List


class Solution:
    def trap(self, height: List[int]) -> int:
        maxL, maxR = height[0], height[-1]
        l, r = 0, len(height)-1
        result = 0
        while l <= r:
            if height[l] < height[r]:
                maxL = max(maxL, height[l])
                result += maxL-height[l]
                l+=1
            else:
                maxR = max(maxR, height[r])
                result += maxR - height[r]
                r-=1
        return result
