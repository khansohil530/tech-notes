class Solution:
    def dailyTemperatures(self, temperatures: list[int]) -> list[int]:
        result = [0]*len(temperatures)
        stk = []
        for r, temp in enumerate(temperatures):
            while len(stk) != 0 and temperatures[stk[-1]] < temp:
                l = stk.pop()
                result[l] = r-l

            stk.append(r)
        return result