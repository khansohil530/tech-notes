from typing import Optional, List
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
class Solution:
    def _helper(self, left, right):
        if left > right:
            return

        val = self.preorder[self.pre_idx]
        self.pre_idx += 1
        mid = self.inorder[val]
        root = TreeNode(val)
        root.left = self._helper(left, mid-1)
        root.right = self._helper(mid+1, right)
        return root

    def buildTree(self, preorder: List[int], inorder: List[int]) -> Optional[TreeNode]:
        self.inorder = {v: idx for idx, v in enumerate(inorder)}
        self.pre_idx = 0
        self.preorder = preorder
        return self._helper(0, len(inorder)-1)
