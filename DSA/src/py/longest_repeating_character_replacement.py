from collections import defaultdict


class Solution:
    def characterReplacement(self, s: str, k: int) -> int:
        maxf = 0
        l = 0
        counts = defaultdict(int)
        for r, ch in enumerate(s):
            counts[ch] += 1
            maxf = max(maxf, counts[ch])
            if r-l+1 - maxf > k: # no more subs
                counts[s[l]] -= 1
                l += 1
        return len(s) - l