from typing import List
from collections import Counter, defaultdict
class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        freq = Counter(nums)
        buckets = defaultdict(list)
        for num, count in freq.items():
            buckets[count].append(num)

        result = []
        for count in range(len(nums), 0, -1):
            if count in buckets:
                result.extend(buckets[count])

        return result[:k]