class Solution:
    def findDuplicate(self, nums: list[int]) -> int:
        slow, fast = 0, 0
        while fast < len(nums):
            slow = nums[slow]
            fast = nums[nums[fast]]
            if slow == fast:
                break

        first, second = 0, slow
        while first != second:
            first = nums[first]
            second = nums[second]
        return first