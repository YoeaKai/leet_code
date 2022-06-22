#include <iostream>
#include <vector>
#include "../../topic/find_two_nonOverlapping_subArrays_each_with_target_sum/find_two_nonOverlapping_subArrays_each_with_target_sum.cpp"
using namespace std;

int main()
{
    Solution s;
    vector<int> v = {3, 2, 2, 4, 3};

    cout << "------start------" << endl;
    cout << s.minSumOfLengths2(v, 3) << endl;
    cout << "------over------" << endl;
}
