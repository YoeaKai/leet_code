#include <vector>
#include <unordered_map>
#include <unordered_set>
using namespace std;

class Solution
{
public:
	int longestConsecutive(vector<int> &nums)
	{
		unordered_set<int> set(nums.begin(), nums.end());
		int ret = 0;

		for (int n : nums)
		{
			if (set.find(n) == set.end())
				continue;

			set.erase(n);
			int pre = n - 1, next = n + 1;

			while (set.find(pre) != set.end())
				set.erase(pre--);

			while (set.find(next) != set.end())
				set.erase(next++);

			ret = max(ret, next - pre - 1);
		}

		return ret;
	}

	int longestConsecutive2(vector<int> &nums)
	{
		unordered_map<int, int> mp;
		int ret = 0;

		for (int num : nums)
		{
			if (!mp[num])
				ret = max(ret, mp[num] = mp[num + mp[num + 1]] = mp[num - mp[num - 1]] = mp[num + 1] + mp[num - 1] + 1);
		}

		return ret;

		// That is mean:
		// int val = mp[num + 1] + mp[num - 1] + 1;
		// mp[num] = val;
		// mp[num + mp[num + 1]] = val;
		// mp[num - mp[num - 1]] = val;
	}
};