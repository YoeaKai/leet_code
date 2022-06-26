#include <iostream>
#include <vector>
#include "../../topic/number_of_longest_increasing_subsequence/number_of_longest_increasing_subsequence.cpp"
using namespace std;

int main()
{
    Solution s;
    vector<int> a = vector<int>{2, 3, 2, 3, 4};

    cout << "------start------" << endl;
    cout << s.findNumberOfLIS(a) << endl;
    cout << "------over------" << endl;
}
