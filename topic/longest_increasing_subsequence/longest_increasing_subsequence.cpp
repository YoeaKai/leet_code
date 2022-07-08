#include <vector>
#include <map>
using namespace std;

class Solution
{
public:
	// Greedy with Binary Search
	int lengthOfLIS(vector<int> &nums)
	{
		vector<int> sub;

		for (int num : nums)
		{
			if (sub.empty() || sub[sub.size() - 1] < num)
			{
				sub.push_back(num);
			}
			else
			{
				auto it = lower_bound(sub.begin(), sub.end(), num);
				*it = num;
			}
		}

		return sub.size();
	}

	// Dynamic Programming
	int lengthOfLIS2(vector<int> &nums)
	{
		int result;
		map<int, int> times;

		for (int &num : nums)
		{
			int maxLen = 1;

			for (auto &val : times)
				if (val.first >= num)
					break;
				else if (maxLen < val.second + 1)
					maxLen = val.second + 1;

			times[num] = max(times[num], maxLen);
			result = max(result, maxLen);
		}

		return result;
	}
};