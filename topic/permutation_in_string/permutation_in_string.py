from collections import Counter


class Solution:
    def checkInclusion(self, s1: str, s2: str) -> bool:
        counter, n, match = Counter(s1), len(s1), 0

        for i in range(len(s2)):
            if s2[i] in counter:
                if not counter[s2[i]]:
                    match -= 1
                counter[s2[i]] -= 1
                if not counter[s2[i]]:
                    match += 1

            if i >= n and s2[i-n] in counter:
                if not counter[s2[i-n]]:
                    match -= 1
                counter[s2[i-n]] += 1
                if not counter[s2[i-n]]:
                    match += 1

            if match == len(counter):
                return True

        return False
