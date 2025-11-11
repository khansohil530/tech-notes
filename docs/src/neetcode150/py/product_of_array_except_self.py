from typing import List
class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        result = []
        pre, post = 1, 1

        for i in range(len(nums)):
            result.append(pre)
            pre *= nums[i]

        for i in reversed(range(len(nums))):
            result[i] *= post
            post *= nums[i]
        print(result)
        return result
