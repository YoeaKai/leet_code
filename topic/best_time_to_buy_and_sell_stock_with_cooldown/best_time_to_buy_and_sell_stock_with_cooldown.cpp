#include <vector>
using namespace std;

class Solution
{
public:
	int maxProfit(vector<int> &prices)
	{
		// Three state
		int buyed = -prices[0], beforeCooldown = INT_MIN, afterCooldown = 0, preBuyed;

		for (int i = 1; i < prices.size(); i++)
		{
			preBuyed = buyed;

			buyed = max(preBuyed, afterCooldown - prices[i]);
			afterCooldown = max(beforeCooldown, afterCooldown);
			beforeCooldown = preBuyed + prices[i];
		}

		// From this:
		// for (int i = 1; i < prices.size(); i++)
		// {
		// 	preBuyed = buyed;
		// 	preBeforeCooldown = beforeCooldown;
		// 	preAfterCooldown = afterCooldown;

		// 	buyed = max(preBuyed, preAfterCooldown - prices[i]);
		// 	beforeCooldown = preBuyed + prices[i];
		// 	afterCooldown = max(preBeforeCooldown, preAfterCooldown);
		// }

		return max(beforeCooldown, afterCooldown);
	}
};