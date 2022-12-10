from topic.maximum_product_of_splitted_binary_tree.maximum_product_of_splitted_binary_tree import Solution
from topic.maximum_product_of_splitted_binary_tree.maximum_product_of_splitted_binary_tree import TreeNode


def run():
    solution = Solution()
    tree2 = TreeNode()
    tree2.val = 2
    tree = TreeNode()
    tree.val = 1
    tree.left = tree2

    print(solution.maxProduct(tree))


if __name__ == '__main__':
    print("----------Start----------")
    run()
    print("----------Finish----------")
