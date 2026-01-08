from typing import Optional
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
class Solution:
    def _helper(self, node):
        if not node:
            return 0, 0
        left_height, left_diameter = self._helper(node.left)
        right_height, right_diameter = self._helper(node.right)
        height = 1+max(left_height, right_height)
        diameter = max(left_diameter, right_diameter, 1+left_height+right_height)
        return height, diameter

    def diameterOfBinaryTree(self, root: Optional[TreeNode]) -> int:
         _, diameter = self._helper(root)
         return diameter
