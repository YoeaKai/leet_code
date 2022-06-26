#include <iostream>
#include <vector>
#include "../../topic/minimum_depth_of_binary_tree/minimum_depth_of_binary_tree.cpp"
using namespace std;

int main()
{
    Solution s;
    TreeNode e = TreeNode{7};
    TreeNode d = TreeNode{15};
    TreeNode c = TreeNode{20, &d, &e};
    TreeNode b = TreeNode{9};
    TreeNode a = TreeNode{3, &b, &c};

    cout << "------start------" << endl;
    cout << s.minDepth(&a) << endl;
    cout << "------over------" << endl;
}
