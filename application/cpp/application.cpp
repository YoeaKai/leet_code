#include <iostream>
#include <vector>
#include "../../topic/find_two_nonOverlapping_subArrays_each_with_target_sum/find_two_nonOverlapping_subArrays_each_with_target_sum.cpp"
using namespace std;

int main()
{
    Solution s;
    vector<int> v = {1, 3, 1};

    cout << "------start------" << endl;
    cout << s.minSumOfLengths(v, 4) << endl;
    cout << "------over------" << endl;
}
