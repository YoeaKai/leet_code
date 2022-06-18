#include <iostream>
#include <vector>
#include "../../topic/jump_game_VI/jump_game_VI.cpp"
using namespace std;

int main()
{
    Solution s;
    vector<int> strs = {10, -5, -2, 4, 0, 3};

    cout << "------start------" << endl;
    cout << s.maxResult(strs, 3) << endl;
    cout << "------over------" << endl;
}
