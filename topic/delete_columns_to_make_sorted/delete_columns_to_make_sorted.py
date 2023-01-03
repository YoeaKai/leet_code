from typing import List


class Solution:
    def minDeletionSize(self, strs: List[str]) -> int:
        return sum(column != sorted(column) for column in map(list, zip(*strs)))
