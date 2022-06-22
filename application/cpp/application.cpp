#include <iostream>
#include <vector>
#include "../../topic/kth_largest_element_in_an_array/kth_largest_element_in_an_array.cpp"
using namespace std;

int main()
{
    Solution s;
    vector<int> v = {3, 2, 1, 5, 6, 4};

    cout << "------start------" << endl;
    cout << s.findKthLargest2(v, 2) << endl;
    cout << "------over------" << endl;
}
