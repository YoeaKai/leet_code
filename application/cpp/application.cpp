#include <iostream>
#include <vector>
#include "../../topic/longest_string_chain/longest_string_chain.cpp"
using namespace std;

int main()
{
    Solution s;
    vector<string> strs = {"xbc", "pcxbcf", "xb", "cxbc", "pcxbc"};

    cout << "------start------" << endl;
    cout << s.longestStrChain(strs) << endl;
    cout << "------over------" << endl;
}
