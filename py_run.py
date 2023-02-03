from topic.verifying_an_alien_dictionary.verifying_an_alien_dictionary import Solution


def run():
    solution = Solution()
    print(solution.isAlienSorted_2(
        ["apple", "app"], "abcdefghijklmnopqrstuvwxyz"))
    print(solution.isAlienSorted_2(
        ["word", "world", "row"], "worldabcefghijkmnpqstuvxyz"))
    print(solution.isAlienSorted_2(
        ["hello", "leetcode"], "hlabcdefgijkmnopqrstuvwxyz"))


if __name__ == '__main__':
    print("----------Start----------")
    run()
    print("----------Finish----------")
