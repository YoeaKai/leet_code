from typing import List
from heapq import nlargest


class Solution:
    def longestPath(self, parent: List[int], s: str) -> int:
        children = [[] for _ in range(len(s))]

        for idx, p in enumerate(parent):
            if p >= 0:
                children[p].append(idx)
        
        res = [0]

        def dfs(node):
            path_len = [0]
            for child in children[node]:
                cur = dfs(child)
                if s[node] != s[child]:
                    path_len.append(cur)

            # Connect the two longest child.
            path_len = nlargest(2, path_len)
            res[0] = max(res[0], sum(path_len) + 1)

            # Choose the longest child for parents.
            return max(path_len) + 1
        
        dfs(0)

        return res[0]
