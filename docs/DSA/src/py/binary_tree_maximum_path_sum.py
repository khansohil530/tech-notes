from typing import Optional
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
class Solution:
    def _helper(self, node):
        if not node:
            return 0

        left_max = max(self._helper(node.left), 0)
        right_max = max(self._helper(node.right), 0)
        self.res = max(self.res, node.val + left_max + right_max)
        return node.val + max(left_max, right_max)



    def maxPathSum(self, root: Optional[TreeNode]) -> int:
        self.res = float('-inf')
        self._helper(root)
        return int(self.res)
