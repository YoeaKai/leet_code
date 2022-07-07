#include <iostream>
#include <vector>
#include "../../topic/maximum_points_you_can_obtain_from_cards/maximum_points_you_can_obtain_from_cards.cpp"
using namespace std;

int main()
{
    Solution s;
    vector<int> a = {1, 2, 3, 4, 5, 6, 1};

    cout << "------start------" << endl;
    cout << s.maxScore(a, 3) << endl;
    cout << "------over------" << endl;
}
