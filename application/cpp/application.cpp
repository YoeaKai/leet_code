#include <iostream>
#include <vector>
#include "../../topic/ones_and_zeroes/ones_and_zeroes.cpp"
using namespace std;

int main()
{
    Solution s;
    vector<string> strs = {"10", "0001", "111001", "1", "0"};
    vector<string> strs2 = {"10", "0", "1"};

    cout << "------start------" << endl;
    cout << s.findMaxForm(strs, 5, 3) << endl;
    cout << s.findMaxForm(strs2, 1, 1) << endl;
    cout << "------over------" << endl;
}
