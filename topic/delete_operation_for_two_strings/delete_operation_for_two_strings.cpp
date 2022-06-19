#include <vector>
#include <string>
using namespace std;
class Solution
{
public:
    int minDistance(string word1, string word2)
    {
        int n1 = word1.size(), n2 = word2.size();
        vector<int> dp(n2 + 1, 0);

        for (int j = 0; j <= n2; j++)
        {
            dp[j] = j;
        }

        int pre, tmp;
        for (int i = 0; i < n1; i++)
        {
            pre = dp[0];
            dp[0] += 1;

            for (int j = 0; j < n2; j++)
            {
                tmp = dp[j + 1];
                dp[j + 1] = word1[i] == word2[j] ? pre : min(dp[j], dp[j + 1]) + 1;
                pre = tmp;
            }
        }

        return dp[n2];
    }
};