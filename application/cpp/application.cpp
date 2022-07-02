#include <iostream>
#include <vector>
#include "../../topic/minimize_the_difference_between_target_and_chosen_elements/minimize_the_difference_between_target_and_chosen_elements.cpp"
using namespace std;

int main()
{
    Solution s;
    vector<vector<int> > a = {{1, 2, 3}, {4, 5, 6}, {7, 8, 9}};

    cout << "------start------" << endl;
    cout << s.minimizeTheDifference(a, 13) << endl;
    cout << "------over------" << endl;
}
