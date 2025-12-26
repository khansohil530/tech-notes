class Solution:
    def isValid(self, s: str) -> bool:
        stk = []
        op = {"(": ")", "{": "}", "[": "]"}
        for ch in s:
            if ch in op:
                stk.append(ch)
            else:
                if len(stk) == 0 or (op[stk[-1]] != ch):
                    return False
                stk.pop()
        return len(stk) == 0
