from collections import deque
from typing import Optional, List


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
class Solution:
    def levelOrder(self, root: Optional[TreeNode]) -> List[List[int]]:
        levels = []
        q = deque([root])
        while q:
            level = []
            for _ in range(len(q)):
                node = q.pop()
                if not node:
                    continue
                level.append(node.val)
                q.appendleft(node.left)
                q.appendleft(node.right)
            if level:
                levels.append(level)
        return levels