from topic.number_of_good_paths.number_of_good_paths import Solution


def run():
    solution = Solution()
    print(solution.numberOfGoodPaths([1,3,2,1,3], [[0,1],[0,2],[2,3],[2,4]]))


if __name__ == '__main__':
    print("----------Start----------")
    run()
    print("----------Finish----------")
