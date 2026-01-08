from typing import Optional
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
class Solution:
    def _helper(self, node, upper_bound, lower_bound):
        if not node:
            return True

        return lower_bound <= node.val <= upper_bound and \
            self._helper(node.left, node.val-1, lower_bound) and \
            self._helper(node.right, upper_bound, node.val+1)

    def isValidBST(self, root: Optional[TreeNode]) -> bool:
        return self._helper(root, float('inf'), float('-inf'))
