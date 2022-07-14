#include <iostream>
#include <vector>
#include "../../topic/construct_binary_tree_from_preorder_and_inorder_traversal/cpp/construct_binary_tree_from_preorder_and_inorder_traversal.cpp"
using namespace std;

int main()
{
    Solution s;
    vector<int> pre = {3, 9, 20, 15, 7};
    vector<int> in = {9, 3, 15, 20, 7};

    cout << "------start------" << endl;
    cout << s.buildTree(pre, in) << endl;
    cout << "------over------" << endl;
}
