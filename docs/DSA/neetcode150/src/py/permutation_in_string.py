from collections import Counter


class Solution:
    def checkInclusion(self, s1: str, s2: str) -> bool:
        counts = Counter(s1)
        l = 0
        for r, ch in enumerate(s2):
            counts[ch] -= 1
            while counts[ch] < 0:
                counts[s2[l]] += 1
                l += 1

            if r - l + 1 == len(s1):
                return True
        return False
