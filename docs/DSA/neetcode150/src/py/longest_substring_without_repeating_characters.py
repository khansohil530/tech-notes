class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        idxMap = dict()
        l = 0
        for r, ch in enumerate(s):
            if ch in idxMap and idxMap[ch] >= l:
                l = idxMap[ch]+1
            idxMap[ch] = r
        return len(s)-l