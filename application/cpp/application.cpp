#include <iostream>
#include <vector>
#include "../../topic/matchsticks_to_square/matchsticks_to_square.cpp"
using namespace std;

int main()
{
    Solution s;
    vector<int> a = {1, 1, 2, 2, 2};

    cout << "------start------" << endl;
    cout << s.makesquare(a) << endl;
    cout << "------over------" << endl;
}
