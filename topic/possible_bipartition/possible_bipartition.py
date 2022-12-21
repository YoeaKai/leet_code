from typing import List


class Solution:
    def possibleBipartition(self, n: int, dislikes: List[List[int]]) -> bool:
        def dfs(person, person_group):
            group[person] = person_group
            for dislike_person in dislike_map[person]:
                if group[dislike_person] == group[person]:
                    return False
                if group[dislike_person] == -1:
                    if not dfs(dislike_person, 1 - person_group):
                        return False
            return True

        dislike_map = [[] for _ in range(n + 1)]
        for dislike in dislikes:
            dislike_map[dislike[0]].append(dislike[1])
            dislike_map[dislike[1]].append(dislike[0])

        # 0 stands for group 0 and 1 stands for group 1.
        group = [-1] * (n + 1)

        for i in range(0, n):
            if group[i] == -1:
                if not dfs(i, 0):
                    return False

        return True
