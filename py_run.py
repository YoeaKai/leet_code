from topic.keys_and_rooms.keys_and_rooms import Solution


def run():
    solution = Solution()
    print(solution.canVisitAllRooms([[1], [2], [3], []]))
    print(solution.canVisitAllRooms([[1, 3], [3, 0, 1], [2], [0]]))


if __name__ == '__main__':
    print("----------Start----------")
    run()
    print("----------Finish----------")
