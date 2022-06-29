#include <iostream>
#include <vector>
#include "../../topic/queue_reconstruction_by_height/queue_reconstruction_by_height.cpp"
using namespace std;

int main()
{
    Solution s;
    vector<vector<int> > a = {{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}};

    cout << "------start------" << endl;
    a = s.reconstructQueue(a);
    // cout << s.reconstructQueue(&a) << endl;
    cout << "------over------" << endl;
}
