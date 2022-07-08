#include <iostream>
#include <vector>
#include "../../topic/longest_increasing_subsequence/longest_increasing_subsequence.cpp"
using namespace std;

int main()
{
    Solution s;
    vector<int> a = {0, 1, -1, 3, 2, 3};

    cout << "------start------" << endl;
    cout << s.lengthOfLIS(a) << endl;
    cout << "------over------" << endl;
}
