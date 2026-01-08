class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
class Solution:
    def _helper(self, node: TreeNode, lastGoodNode: TreeNode) -> int:
        if not node:
            return 0

        tmp = 0
        if node.val >= lastGoodNode.val:
            tmp = 1
            lastGoodNode = node

        return tmp + self._helper(node.left, lastGoodNode) + self._helper(node.right, lastGoodNode)

    def goodNodes(self, root: TreeNode) -> int:
        return self._helper(root, root)