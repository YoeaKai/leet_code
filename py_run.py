from topic.concatenated_words.concatenated_words import Solution


def run():
    solution = Solution()
    print(solution.findAllConcatenatedWordsInADict(["cat", "dog", "catdog"]))


if __name__ == '__main__':
    print("----------Start----------")
    run()
    print("----------Finish----------")
