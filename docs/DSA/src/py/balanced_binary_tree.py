from typing import Optional
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
class Solution:
    def _helper(self, node):
        if not node:
            return True, 0
        lb, lh = self._helper(node.left)
        rb, rh = self._helper(node.right)
        b = abs(lh-rh) <= 1 and lb and rb
        h = 1+max(lh, rh)
        return b, h

    def isBalanced(self, root: Optional[TreeNode]) -> bool:
        balanced, _ = self._helper(root)
        return balanced
