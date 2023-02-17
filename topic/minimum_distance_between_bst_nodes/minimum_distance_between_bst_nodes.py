from typing import Optional
from sys import maxsize


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def minDiffInBST(self, root: Optional[TreeNode]) -> int:
        self.prev = -maxsize - 1

        def dfs(node) -> int:
            if not node:
                return maxsize
            min_diff = min(dfs(node.left), node.val - self.prev)
            self.prev = node.val
            return min(min_diff, dfs(node.right))

        return dfs(root)
