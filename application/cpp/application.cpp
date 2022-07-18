#include <iostream>
#include <vector>
#include "../../topic/find_the_town_judge/find_the_town_judge.cpp"
using namespace std;

int main()
{
    Solution s;
    vector<vector<int> > a = {{1, 2}};

    cout << "------start------" << endl;
    cout << s.findJudge(2, a) << endl;
    cout << "------over------" << endl;
}
