#include <string>
#include <map>
#include <unordered_set>
using namespace std;

class Solution
{
public:
	int minDeletions(string s)
	{
		int cnt[26] = {}, ret = 0;
		for (auto &c : s)
			cnt[c - 'a']++;

		sort(begin(cnt), end(cnt));

		for (int i = 24; i >= 0 && cnt[i] > 0; --i)
		{
			ret += max(0, cnt[i] - max(0, cnt[i + 1] - 1));
			cnt[i] = min(cnt[i], max(0, cnt[i + 1] - 1));
		}

		return ret;
	}
};