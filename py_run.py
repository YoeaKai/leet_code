from topic.leaf_similar_trees.leaf_similar_trees import Solution
from topic.leaf_similar_trees.leaf_similar_trees import TreeNode


def run():
    solution = Solution()
    tree = TreeNode()
    tree.val = 1
    print(solution.leafSimilar(tree, tree))


if __name__ == '__main__':
    print("----------Start----------")
    run()
    print("----------Finish----------")
