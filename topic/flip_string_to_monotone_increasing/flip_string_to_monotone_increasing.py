class Solution:
    # Dynamic programming solution
    def minFlipsMonoIncr(self, s: str) -> int:
        flip = 1
        no_flip = 0

        for idx in range(1, len(s)):
            pre_flip = flip
            if s[idx] == '0':
                flip = min(no_flip, flip)+1
                # if s[idx-1] == '0', no_flip no change.
                if s[idx-1] == '1':
                    no_flip = pre_flip
            else:
                if s[idx-1] == '0':
                    flip = no_flip + 1
                else:
                    flip += 1
                no_flip = min(no_flip, pre_flip)

        return min(flip, no_flip)

    # The solution of counting the changed number
    def minFlipsMonoIncr_2(self, s: str) -> int:
        count_0 = s.count('0')
        count_1 = 0
        res = len(s) - count_0

        for idx in range(len(s)):
            if s[idx] == '0':
                count_0 -= 1
            elif s[idx] == '1':
                res = min(res, count_0 + count_1)
                count_1 += 1

        return res
