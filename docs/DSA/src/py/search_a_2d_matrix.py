class Solution:
    def searchMatrix(self, matrix: list[list[int]], target: int) -> bool:
        m, n = len(matrix), len(matrix[0])
        l, r = 0, len(matrix)*len(matrix[0])
        mid = 0
        while l < r:
            mid = l + (r-l)//2
            if matrix[mid//n][mid%n] >= target:
                r = mid
            else:
                l = mid + 1

        if l < m*n and matrix[l//n][l%n] == target:
            return True
        return False