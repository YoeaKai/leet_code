#include <string>
#include <vector>
#include <unordered_map>
using namespace std;

class Solution
{
public:
	static bool compare(const string &s1, const string &s2)
	{
		return s1.size() < s2.size();
	}

	int longestStrChain(vector<string> &words)
	{
		sort(words.begin(), words.end(), compare);
		unordered_map<string, int> dp;
		int maxLength = 0;

		for (auto &word : words)
		{
			for (int i = 0; i < word.size(); i++)
			{
				string pre = word.substr(0, i) + word.substr(i + 1);
				dp[word] = max(dp[word], dp[pre] + 1);
			}
			if (dp[word] > maxLength)
			{
				maxLength = dp[word];
			}
		}

		return maxLength;
	}
};