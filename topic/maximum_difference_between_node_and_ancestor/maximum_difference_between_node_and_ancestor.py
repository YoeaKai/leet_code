from typing import Optional


# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def maxAncestorDiff(self, root: Optional[TreeNode]) -> int:
        def dfs(root, minValue, maxValue):
            if not root:
                return 0

            res = max(abs(root.val - minValue), abs(root.val - maxValue))

            minValue, maxValue = min(
                minValue, root.val), max(maxValue, root.val)

            return max(res, dfs(root.left, minValue, maxValue), dfs(root.right, minValue, maxValue))

        return dfs(root, root.val, root.val)
