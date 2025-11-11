from typing import List
class Solution:
    def _get_bitmap(self, s: str) -> tuple:
        l = [0] * 26
        for ch in s:
            l[ord(ch) - ord('a')] += 1
        return tuple(l)

    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        freq = {}  # bitmap -> str

        for s in strs:
            bitmap = self._get_bitmap(s)
            if bitmap in freq:
                freq[bitmap].append(s)
            else:
                freq[bitmap] = [s]
        return list(freq.values())