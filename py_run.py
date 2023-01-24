from topic.snakes_and_ladders.snakes_and_ladders import Solution


def run():
    solution = Solution()
    print(solution.snakesAndLadders([[-1, -1, -1, -1, -1, -1], [-1, -1, -1, -1, -1, -1], [-1, -1, -1, -
          1, -1, -1], [-1, 35, -1, -1, 13, -1], [-1, -1, -1, -1, -1, -1], [-1, 15, -1, -1, -1, -1]]))


if __name__ == '__main__':
    print("----------Start----------")
    run()
    print("----------Finish----------")
