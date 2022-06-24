#include <vector>
using namespace std;
class Solution
{
	int limit = 2e8;

public:
	vector<vector<int> > removeUnused(vector<vector<int> > &tires)
	{
		sort(tires.begin(), tires.end());
		vector<vector<int> > ret;

		for (auto &t : tires)
			if (ret.empty() || ret.back()[1] > t[1])
				ret.push_back(t);

		return ret;
	}

	int minimumFinishTime(vector<vector<int> > &tires, int changeTime, int numLaps)
	{
		tires = removeUnused(tires);
		int n = tires.size();

		// to handle the cases where numLaps is small
		// without_change[i][j]: the total time to run j laps consecutively with tire i
		vector<vector<int> > without_change(n, vector<int>(20, limit));
		for (int i = 0; i < n; i++)
		{
			without_change[i][1] = tires[i][0];
			for (int j = 2; j < 20; j++)
			{
				if ((long long)without_change[i][j - 1] * tires[i][1] >= limit)
					break;
				without_change[i][j] = without_change[i][j - 1] * tires[i][1];
			}
			// since we define it as the total time, rather than just the time for the j-th lap
			// we have to make it prefix sum
			for (int j = 2; j < 20; j++)
			{
				if ((long long)without_change[i][j - 1] + without_change[i][j] >= limit)
					break;
				without_change[i][j] += without_change[i][j - 1];
			}
		}

		// dp[x]: the minimum time to finish x laps
		vector<int> dp(numLaps + 1, limit);
		for (int i = 0; i < n; i++)
		{
			dp[1] = min(dp[1], tires[i][0]);
		}
		for (int x = 1; x <= numLaps; x++)
		{
			if (x < 20)
			{
				// x is small enough, so an optimal solution might never changes tires!
				for (int i = 0; i < n; i++)
				{
					dp[x] = min(dp[x], without_change[i][x]);
				}
			}
			for (int j = x - 1; j > 0 && j >= x - 18; j--)
			{
				dp[x] = min(dp[x], dp[j] + changeTime + dp[x - j]);
			}
		}

		return dp[numLaps];
	}

	// Optimize failed
	// int minimumFinishTime(vector<vector<int> > &tires, int changeTime, int numLaps)
	// {
	// 	tires = removeUnused(tires);
	// 	int n = tires.size();

	// 	// Find the minimum time for each lap.
	// 	vector<int> withoutChangeMin(20, limit);
	// 	for (int i = 0; i < n; i++)
	// 	{
	// 		int pre = tires[i][0], cur = tires[i][0], tmp;
	// 		withoutChangeMin[1] = min(withoutChangeMin[1], cur);

	// 		for (int j = 2; j < 20; j++)
	// 		{
	// 			if ((long long)pre * tires[i][1] >= limit)
	// 				break;
	// 			tmp = cur;
	// 			cur = pre * tires[i][1];
	// 			pre = cur;
	// 			// since we define it as the total time, rather than just the time for the j-th lap we have to make it prefix sum.
	// 			cur += tmp;
	// 			withoutChangeMin[j] = min(withoutChangeMin[j], cur);
	// 		}
	// 	}

	// 	// withoutChangeMin change to minTime.
	// 	for (int x = 1; x <= numLaps; x++)
	// 	{
	// 		for (int j = x - 1; j > 0 && j >= x - 18; j--)
	// 		{
	// 			withoutChangeMin[x] = min(withoutChangeMin[x], withoutChangeMin[j] + changeTime + withoutChangeMin[x - j]);
	// 		}
	// 	}

	// 	return withoutChangeMin[numLaps];
	// }
};