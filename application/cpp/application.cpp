#include <iostream>
#include <vector>
#include "../../topic/maximum_width_of_binary_tree/maximum_width_of_binary_tree.cpp"
using namespace std;

int main()
{
    Solution s;
    TreeNode e = TreeNode(5);
    TreeNode d = TreeNode(4);
    TreeNode c = TreeNode(3, nullptr, &e);
    TreeNode b = TreeNode(2, &d, nullptr);
    TreeNode a = TreeNode(1, &b, &c);

    cout << "------start------" << endl;
    cout << s.widthOfBinaryTree(&a) << endl;
    cout << "------over------" << endl;
}
