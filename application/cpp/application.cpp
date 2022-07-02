#include <iostream>
#include <vector>
#include "../../topic/maximum_area_of_a_piece_of_cake_after_horizontal_and_vertical_cuts/maximum_area_of_a_piece_of_cake_after_horizontal_and_vertical_cuts.cpp"
using namespace std;

int main()
{
    Solution s;
    vector<int> a = {2};
    vector<int> b = {2};

    cout << "------start------" << endl;
    cout << s.maxArea(1000000000, 1000000000, a, b) << endl;
    cout << "------over------" << endl;
}
