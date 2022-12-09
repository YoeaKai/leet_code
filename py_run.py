from topic.maximum_difference_between_node_and_ancestor.maximum_difference_between_node_and_ancestor import Solution
from topic.maximum_difference_between_node_and_ancestor.maximum_difference_between_node_and_ancestor import TreeNode


def run():
    solution = Solution()
    tree2 = TreeNode()
    tree2.val = 2
    tree = TreeNode()
    tree.val = 1
    tree.left = tree2

    print(solution.maxAncestorDiff(tree, tree))


if __name__ == '__main__':
    print("----------Start----------")
    run()
    print("----------Finish----------")
