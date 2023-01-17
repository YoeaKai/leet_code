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
			int tmp = tires[i][0] = without_change[i][1] = tires[i][0];

			for (int j = 2; j < 20; j++)
			{
				if ((long long)without_change[i][j - 1] * tires[i][1] >= limit)
					break;
				without_change[i][j] = tmp * tires[i][1];
				tmp = without_change[i][j];
				// since we define it as the total time, rather than just the time for the j-th lap we have to make it prefix sum.
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
};