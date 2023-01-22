from topic.word_break.word_break import Solution


def run():
    solution = Solution()
    print(solution.wordBreak("catsandog", [
          "cats", "dog", "sand", "and", "cat"]))


if __name__ == '__main__':
    print("----------Start----------")
    run()
    print("----------Finish----------")
