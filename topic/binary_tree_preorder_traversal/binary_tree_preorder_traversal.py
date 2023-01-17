from typing import Optional, List


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def preorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
        if not root:
            return []

        return [root.val] + self.preorderTraversal(root.left) + self.preorderTraversal(root.right)

    def preorderTraversal_2(self, root: Optional[TreeNode]) -> List[int]:
        preorder = []

        def dfs(node):
            if not node:
                return
            preorder.append(node.val)
            dfs(node.left)
            dfs(node.right)

        dfs(root)

        return preorder
