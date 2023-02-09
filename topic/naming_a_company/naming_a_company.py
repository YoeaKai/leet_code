from typing import List
from collections import defaultdict


class Solution:
    def distinctNames(self, ideas: List[str]) -> int:
        count = defaultdict(set)
        for idea in ideas:
            count[idea[0]].add(hash(idea[1:]))

        res = 0

        for idea_1, set_1 in count.items():
            for idea_2, set_2 in count.items():
                if idea_1 >= idea_2:
                    continue
                same = len(set_1 & set_2)
                res += (len(set_1) - same) * (len(set_2) - same)

        return res * 2
