class Solution:
    def minWindow(self, s: str, t: str) -> str:
        need = Counter(t)
        missing = len(t)
        l = R = L = 0
        for r, ch in enumerate(s, 1):
            if need[ch] > 0:
                missing -= 1
            need[ch]-=1
            if missing == 0:
                while l < r and need[s[l]] < 0:
                    need[s[l]]+=1
                    l+=1
                if R == 0 or r-l < R-L:
                    R, L = r, l
                need[s[l]] += 1
                missing += 1
                l += 1
        return s[L:R]