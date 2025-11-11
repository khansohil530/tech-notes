from typing import List
class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        s = dict()
        for i, num in enumerate(nums):
            if num in s:
                return [i, s[num]]
            s[target - num] = i
        return [-1, -1]