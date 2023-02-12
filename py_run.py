from topic.minimum_fuel_cost_to_report_to_the_capital.minimum_fuel_cost_to_report_to_the_capital import Solution


def run():
    solution = Solution()
    print(solution.minimumFuelCost(
        [[3, 1], [3, 2], [1, 0], [0, 4], [0, 5], [4, 6]], 2))


if __name__ == '__main__':
    print("----------Start----------")
    run()
    print("----------Finish----------")
