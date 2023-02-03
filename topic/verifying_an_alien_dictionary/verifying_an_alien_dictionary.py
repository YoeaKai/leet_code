from typing import List


class Solution:
    def isAlienSorted(self, words, order):
        dic = {c: i for i, c in enumerate(order)}
        words = [[dic[c] for c in word] for word in words]
        return all(word_1 <= word_2 for word_1, word_2 in zip(words, words[1:]))

    def isAlienSorted_2(self, words: List[str], order: str) -> bool:
        def index(c):
            return ord(c) - 97  # The ASCII of 'a' is 97.

        dic = [0 for i in range(26)]
        for idx in range(1, len(order)):
            dic[index(order[idx])] = idx

        def isBigger(s1, s2):
            for idx in range(min(len(s1), len(s2))):
                if dic[index(s1[idx])] != dic[index(s2[idx])]:
                    return dic[index(s1[idx])] > dic[index(s2[idx])]
            return len(s1) > len(s2)

        for idx in range(1, len(words)):
            if isBigger(words[idx-1], words[idx]):
                return False

        return True
