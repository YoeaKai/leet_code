#include <vector>
#include <unordered_map>
using namespace std;

class Solution
{
public:
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
};