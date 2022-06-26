#include <vector>
#include <map>
using namespace std;

class Solution
{
public:
  int findLen(vector<vector<pair<int, int> > > &dp, int r, int target)
  {
    int l = 0;

    while (l < r)
    {
      int mid = l + (r - l) / 2;
      if (dp[mid].back().first < target)
        l = mid + 1;
      else
        r = mid;
    }

    return l;
  }

  int binarySearch(vector<pair<int, int> > &dp, int r, int target)
  {
    int l = 0;

    while (l < r)
    {
      int mid = l + (r - l) / 2;
      if (dp[mid].first >= target)
        l = mid + 1;
      else
        r = mid;
    }

    return l;
  }

  int findNumberOfLIS(vector<int> &nums)
  {
    // The index of dp means the max length of LIS.
    // dp[i] store a vector that memorize pairs
    // of max value and its number of LIS from 0 to i.
    vector<vector<pair<int, int> > > dp(nums.size() + 1);
    int curMax = 0;

    for (int i = 0; i < nums.size(); ++i)
    {
      // bsearch insertion point
      int l = findLen(dp, curMax, nums[i]);

      // bsearch number of options
      int options = 1;
      int len = l - 1;
      if (len >= 0)
      {
        int l1 = binarySearch(dp[len], dp[len].size(), nums[i]);

        options = (l1 == 0) ? dp[len].back().second : dp[len].back().second - dp[len][l1 - 1].second;
      }

      dp[l].push_back({nums[i], (dp[l].empty() ? options : dp[l].back().second + options)});

      if (l == curMax)
        curMax++;
    }

    return dp[curMax - 1].back().second;
  }
};