from topic.single_threaded_CPU.single_threaded_CPU import Solution


def run():
    solution = Solution()
    print(solution.getOrder([[1, 2], [2, 4], [3, 2], [4, 1]]))


if __name__ == '__main__':
    print("----------Start----------")
    run()
    print("----------Finish----------")
