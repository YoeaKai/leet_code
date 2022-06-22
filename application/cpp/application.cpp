#include <iostream>
#include <vector>
#include "../../topic/sort_list/sort_list.cpp"
using namespace std;

int main()
{
    Solution s;
    ListNode d = ListNode(1);
    ListNode c = ListNode(2, &d);
    ListNode b = ListNode(3, &c);
    ListNode a = ListNode(4, &b);

    cout << "------start------" << endl;
    s.sortList(&a);
    cout << "------over------" << endl;
}
