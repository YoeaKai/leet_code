from topic.possible_bipartition.possible_bipartition import Solution


def run():
    solution = Solution()
    print(solution.possibleBipartition(4, [[1, 2], [1, 3], [2, 4]]))


if __name__ == '__main__':
    print("----------Start----------")
    run()
    print("----------Finish----------")
