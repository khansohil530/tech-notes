class Solution:
    def carFleet(self, target: int, position: list[int], speed: list[int]) -> int:
        result = 0
        stk = []
        for dist, speed in sorted(zip(position, speed), reverse=True):
            stk.append((target-dist)/speed)
            if len(stk) >= 2 and stk[-1] <= stk[-2]:
                stk.pop()
        return len(stk)