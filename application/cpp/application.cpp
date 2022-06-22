#include <iostream>
#include <vector>
#include "../../topic/longest_consecutive_sequence/longest_consecutive_sequence.cpp"
using namespace std;

int main()
{
    Solution s;
    vector<int> a = vector<int>{9};

    cout << "------start------" << endl;
    cout << s.longestConsecutive(a) << endl;
    cout << "------over------" << endl;
}
