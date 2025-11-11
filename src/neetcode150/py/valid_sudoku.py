from collections import defaultdict
from typing import List

class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        rows = defaultdict(set)
        cols = defaultdict(set)
        box = defaultdict(set)
        for r in range(len(board)):
            for c in range(len(board[0])):
                cell = board[r][c]
                b = (r // 3) * 3 + c // 3
                if cell == ".":
                    continue
                elif cell in rows[r] or cell in cols[c] or cell in box[b]:
                    return False

                rows[r].add(cell)
                cols[c].add(cell)
                box[b].add(cell)
        return True

