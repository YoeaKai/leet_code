#include <string>
using namespace std;

class Solution
{
public:
	bool isSubsequence(string s, string t)
	{
		int idx = 0, n = s.size(), m = t.size();

		for (int i = 0; i < m && idx < n; i++)
			if (t[i] == s[idx])
				idx++;

		return idx == n;
	}
};