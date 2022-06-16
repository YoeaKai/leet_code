#include <iostream>
#include <vector>
#include "../../topic/convert_sorted_array_to_binary_search_tree/convert_sorted_array_to_binary_search_tree.cpp"
using namespace std;

int main()
{
    Solution s;
    vector<int> p = {-10, -3, 0, 5, 9};
    cout << "------start------" << endl;
    cout << s.sortedArrayToBST(p) << endl;
    cout << "------over------" << endl;
}
