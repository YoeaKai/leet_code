#include <iostream>
#include <vector>
#include "../../topic/maximum_length_of_pair_chain/maximum_length_of_pair_chain.cpp"
using namespace std;

int main()
{
    Solution s;
    vector<vector<int> > a = vector<vector<int> >{{2, 2}, {1, 10}, {3, 4}};

    cout << "------start------" << endl;
    cout << s.findLongestChain(a) << endl;
    cout << "------over------" << endl;
}
