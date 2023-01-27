from typing import List


class Solution:
    def findAllConcatenatedWordsInADict(self, words: List[str]) -> List[str]:
        dic = set(words)
        exist = {}

        def dfs(word):
            if word in exist:
                return exist[word]

            exist[word] = False

            for idx in range(1, len(word)):
                prefix = word[:idx]
                suffix = word[idx:]

                if prefix in dic and suffix in dic:
                    exist[word] = True
                    break
                if prefix in dic and dfs(suffix):
                    exist[word] = True
                    break
                if dfs(prefix) and suffix in dic:
                    exist[word] = True
                    break

            return exist[word]

        return [word for word in words if dfs(word)]
