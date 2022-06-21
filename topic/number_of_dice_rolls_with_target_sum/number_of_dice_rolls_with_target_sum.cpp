#include <vector>
using namespace std;
class Solution
{
public:
  //Further Optimizations
  int numRollsToTarget(int n, int k, int target)
  {
    vector<int> dp(target + 1);
    dp[0] = 1;
    int i, j, m;
    for (i = 1; i <= n; i++)
      for (j = target; j >= (i == n ? target : 0); j--)
        for (m = j - 1, dp[j] = 0; m >= max(0, j - k); m--)
          dp[j] = (dp[j] + dp[m]) % 1000000007;
    return dp[target];
  }

  // Original
  int numRollsToTarget2(int n, int k, int target)
  {
    vector<int> dp(target + 1);
    dp[0] = 1;

    for (int i = 1; i <= n; i++)
    {
      vector<int> tmp(target + 1);

      // For each number of face-up, find the possibility of each sum.
      for (int j = 1; j <= k; j++)
      {
        for (int m = j; m <= target; m++)
        {
          tmp[m] = (tmp[m] + dp[m - j]) % 1000000007;
        }
      }

      // // For each sum, find the possibility of each number of face-up.
      // for (int j = 1; j <= target; j++)
      // {
      //   for (int m = 1; m <= k && j >= m; m++)
      //   {
      //     tmp[j] = (tmp[j] + dp[j - m]) % 1000000007;
      //   }
      // }

      dp = tmp;
    }

    return dp[target];
  }
};