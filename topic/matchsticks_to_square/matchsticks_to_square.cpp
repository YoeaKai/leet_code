#include <vector>
#include <map>
#include <numeric>
using namespace std;

class Solution
{
public:
	// DP solution
	bool makesquare(vector<int> &nums)
	{
		int sum = accumulate(nums.begin(), nums.end(), 0), sideSum;
		if (sum % 4)
			return false;

		int n = nums.size(), all = (1 << n) - 1, sideLen = sum / 4;
		vector<int> sideCombinations, doubleSideCombinations(1 << n, false);

		for (int i = 0; i <= all; i++)
		{
			sideSum = 0;

			for (int j = 0; j <= 15; j++)
				if (i >> j & 1)
					sideSum += nums[j];

			if (sideSum == sideLen)
			{
				for (int k : sideCombinations)
				{
					if ((k & i) == 0) // There is no reused matchsticks.
					{
						int t = k | i;
						doubleSideCombinations[t] = true;
						if (doubleSideCombinations[all ^ t]) // XOR, means that the remaining matchsticks are used.
							return true;
					}
				}
				sideCombinations.push_back(i);
			}
		}

		return false;
	}

	// dfs solution
	bool dfs(vector<int> &sidesLength, const vector<int> &matches, int index, const int target)
	{
		if (index == matches.size())
			return sidesLength[0] == sidesLength[1] && sidesLength[1] == sidesLength[2] && sidesLength[2] == sidesLength[3];

		for (int i = 0; i < 4; ++i)
		{
			if (sidesLength[i] + matches[index] > target)
				continue;

			// Do not check the same case.
			int j = i;
			while (--j >= 0)
				if (sidesLength[i] == sidesLength[j])
					break;
			if (j != -1)
				continue;

			sidesLength[i] += matches[index];

			if (dfs(sidesLength, matches, index + 1, target))
				return true;

			sidesLength[i] -= matches[index];
		}

		return false;
	}

	bool makesquare2(vector<int> &matchsticks)
	{
		int sum = 0;
		for (int &val : matchsticks)
			sum += val;

		if (sum % 4 != 0)
			return false;

		sort(matchsticks.begin(), matchsticks.end(), [](const int &l, const int &r)
			 { return l > r; });

		vector<int> sidesLength(4, 0);

		return dfs(sidesLength, matchsticks, 0, sum / 4);
	}
};