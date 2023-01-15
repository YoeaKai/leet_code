class Solution:
    def smallestEquivalentString(self, s1: str, s2: str, baseStr: str) -> str:
        smaller_dict = {}

        def find(char):
            smaller_dict.setdefault(char,char)
            if char != smaller_dict[char]:
                smaller_dict[char] = find(smaller_dict[char])
            return smaller_dict[char]

        def union(c1,c2):
            smallest1 = find(c1)
            smallest2 = find(c2)

            if smallest1>smallest2:
                smaller_dict[smallest1] = smallest2
            else:
                smaller_dict[smallest2] = smallest1

        for idx in range(len(s1)):
            union(s1[idx],s2[idx])
        
        res = []

        for char in baseStr:
            res.append(find(char))
            
        return ''.join(res)
