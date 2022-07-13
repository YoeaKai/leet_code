#include <iostream>
#include <vector>
#include "../../topic/best_time_to_buy_and_sell_stock_with_cooldown/best_time_to_buy_and_sell_stock_with_cooldown.cpp"
using namespace std;

int main()
{
    Solution s;
    vector<int> a = {1, 2, 3, 0, 2};

    cout << "------start------" << endl;
    cout << s.maxProfit(a) << endl;
    cout << "------over------" << endl;
}
