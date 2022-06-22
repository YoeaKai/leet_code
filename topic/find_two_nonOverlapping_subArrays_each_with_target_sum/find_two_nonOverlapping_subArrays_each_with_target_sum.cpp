#include <vector>
#include <unordered_map>
using namespace std;

class Solution
{
public:
	// hash map
	int minSumOfLengths(vector<int> &arr, int target)
	{
		int ret = INT_MAX, dis = INT_MAX, sum = 0;
		unordered_map<int, int> mp;
		mp[0] = -1;

		for (int i = 0; i < arr.size(); i++)
		{
			sum += arr[i];
			mp[sum] = i;
		}

		sum = 0;

		for (int i = 0; i < arr.size(); i++)
		{
			sum += arr[i];

			if (mp.find(sum - target) != mp.end())
				if (i - mp[sum - target] < dis)
					dis = i - mp[sum - target];

			// mp.find(sum + target) searches for any sub-array starting with index i+1 with sum target.
			if (mp.find(sum + target) != mp.end() && dis != INT_MAX)
				ret = min(ret, mp[sum + target] - i + dis);
		}

		return ret == INT_MAX ? -1 : ret;
	}

	// sliding window, faster
	int minSumOfLengths2(vector<int> &arr, int target)
	{
		const int N = arr.size(), INF = 1e9;
		vector<int> dp(N, INF);
		int sum = 0, ret = INF;

		for (int hi = 0, lo = 0; hi < N; hi++)
		{
			sum += arr[hi];

			while (sum > target)
				sum -= arr[lo++];

			if (sum == target)
			{
				int len = hi - lo + 1;
				dp[hi] = len;
				if (lo > 0)
					ret = min(ret, len + dp[lo - 1]);
			}

			if (hi > 0)
				dp[hi] = min(dp[hi], dp[hi - 1]);
		}

		return ret == INF ? -1 : ret;
	}
};