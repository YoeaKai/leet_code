from topic.all_paths_from_source_to_target.all_paths_from_source_to_target import Solution


def run():
    solution = Solution()
    print(solution.allPathsSourceTarget([[1, 2], [3], [3], []]))


if __name__ == '__main__':
    print("----------Start----------")
    run()
    print("----------Finish----------")
