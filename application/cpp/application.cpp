#include <iostream>
#include <vector>
#include "../../topic/find_mode_in_binary_search_tree/find_mode_in_binary_search_tree.cpp"
using namespace std;

int main()
{
    Solution s;
    TreeNode c = TreeNode(2);
    TreeNode b = TreeNode(2, &c, nullptr);
    TreeNode a = TreeNode(1, nullptr, &b);

    cout << "------start------" << endl;
    vector<int> ans = s.findMode(&a);
    cout << "------over------" << endl;
}
