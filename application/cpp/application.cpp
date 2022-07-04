#include <iostream>
#include <vector>
#include "../../topic/wiggle_subsequence/wiggle_subsequence.cpp"
using namespace std;

int main()
{
    Solution s;
    vector<int> a = {1, 7, 4, 9, 2, 5};

    cout << "------start------" << endl;
    cout << s.wiggleMaxLength(a) << endl;
    cout << "------over------" << endl;
}
