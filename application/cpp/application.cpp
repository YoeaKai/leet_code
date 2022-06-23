#include <iostream>
#include <vector>
#include "../../topic/minimum_time_to_finish_the_race/minimum_time_to_finish_the_race.cpp"
using namespace std;

int main()
{
    Solution s;
    vector<vector<int> > a = vector<vector<int> >{{1, 10}, {2, 2}, {3, 4}};

    cout << "------start------" << endl;
    cout << s.minimumFinishTime(a, 6, 5) << endl;
    cout << "------over------" << endl;
}
