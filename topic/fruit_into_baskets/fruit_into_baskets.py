from typing import List


class Solution:
    def totalFruit(self, fruits: List[int]) -> int:
        count, start, cur = {}, 0, 0

        for cur, fruit in enumerate(fruits):
            count[fruit] = count.get(fruit, 0) + 1
            if len(count) > 2:
                count[fruits[start]] -= 1
                if count[fruits[start]] == 0:
                    del count[fruits[start]]
                start += 1

        return cur - start + 1
