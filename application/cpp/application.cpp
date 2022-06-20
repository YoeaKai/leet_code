#include <iostream>
#include <vector>
#include "../../topic/smallest_string_starting_from_leaf/smallest_string_starting_from_leaf.cpp"
using namespace std;

int main()
{
    Solution s;
    TreeNode c = TreeNode(2);
    TreeNode b = TreeNode(2);
    TreeNode a = TreeNode(1, &b, &b);

    cout << "------start------" << endl;
    cout << s.smallestFromLeaf(&a) << endl;
    cout << "------over------" << endl;
}
