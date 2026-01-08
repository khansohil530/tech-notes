from typing import Optional
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
class Solution:
    def z_function(self, s: str) -> list[int]:
        n = len(s)
        z = [0]*n
        l=r=0
        for i in range(1, n):
            if i > r:
                l=r=i
                while r < n and s[r] == s[r-l]:
                    r+=1
                z[i] = r-l
                r-=1
            else:
                k = i-l
                if z[k] < r-i+1:
                    z[i] = z[k]
                else:
                    l = i
                    while r < n and s[r] == s[r-l]:
                        r+=1
                    z[i] = r-l
                    r-=1
        return z

    def _seralize(self, root: Optional[TreeNode]) -> str:
        if not root:
            return "$N"
        return f"${root.val}{self._seralize(root.left)}{self._seralize(root.right)}"

    def isSubtree(self, root: Optional[TreeNode], subRoot: Optional[TreeNode]) -> bool:
        substr = self._seralize(subRoot)
        text = self._seralize(root)
        z = self.z_function(f"{substr}|{text}")
        k = len(substr)
        for i in range(len(text)):
            if z[i+k+1] == k:
                return True
        return False