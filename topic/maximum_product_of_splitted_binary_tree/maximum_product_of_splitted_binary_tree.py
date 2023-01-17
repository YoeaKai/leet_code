from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def __init__(self) -> None:
        self.res = 0

    def maxProduct(self, root: Optional[TreeNode]) -> int:
        def dfs(root):
            if not root:
                return 0
            left, right = dfs(root.left), dfs(root.right)
            self.res = max(self.res, left * (valSum - left),
                           right * (valSum - right))
            return left + right + root.val

        self.res = 0
        valSum = 0
        valSum = dfs(root)
        dfs(root)
        return self.res % (10**9 + 7)
