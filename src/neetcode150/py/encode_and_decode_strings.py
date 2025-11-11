from typing import List
class Solution:
    DELIMITER = "#"
    def encode(self, strs: List[str]) -> str:
        tokens = []
        for s in strs:
            token = f"{len(s)}{self.DELIMITER}{s}"
            tokens.append(token)

        return "".join(tokens)

    def decode(self, s: str) -> List[str]:
        start = curr = 0
        strs = []
        while curr < len(s):
            while s[curr] != self.DELIMITER:
                curr+=1

            size = int(s[start:curr])
            strs.append(s[curr+1: curr+size+1])
            start = curr+size+1
            curr = curr+size+1

        return strs