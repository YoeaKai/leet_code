#include <iostream>
#include <vector>
#include "../../topic/binary_tree_level_order_traversal_II/binary_tree_level_order_traversal_II.cpp"
using namespace std;

int main()
{
    Solution s;
    TreeNode c = TreeNode(20);
    TreeNode b = TreeNode(9);
    TreeNode a = TreeNode(3, &b, &c);

    cout << "------start------" << endl;
    vector<vector<int> > ans = s.levelOrderBottom(&a);
    cout << "------over------" << endl;
}
