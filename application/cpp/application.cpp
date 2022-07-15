#include <iostream>
#include <vector>
#include "../../topic/max_area_of_island/max_area_of_island.cpp"
using namespace std;

int main()
{
    Solution s;
    vector<vector<int> > a = {{0, 0}};

    cout << "------start------" << endl;
    cout << s.maxAreaOfIsland(a) << endl;
    cout << "------over------" << endl;
}
