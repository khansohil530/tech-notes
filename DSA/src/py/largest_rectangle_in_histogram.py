class Solution:
    def largestRectangleArea(self, heights: list[int]) -> int:
        n = len(heights)
        lp = [-1]*n
        stk = []
        for i in range(n):
            while stk and heights[stk[-1]] >= heights[i]:
                stk.pop()
            if stk:
                lp[i] = stk[-1]
            stk.append(i)

        rp = [n]*n
        stk = []
        for i in range(n-1, -1, -1):
            while stk and heights[i] <= heights[stk[-1]]:
                stk.pop()
            if stk:
                rp[i] = stk[-1]
            stk.append(i)

        result = 0
        for i, height in enumerate(heights):
            area = height * (rp[i] - lp[i] + 1)
            result = max(result, area)

        return result
