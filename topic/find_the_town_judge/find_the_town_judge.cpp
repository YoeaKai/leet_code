#include <vector>
using namespace std;

class Solution
{
public:
	int findJudge(int n, vector<vector<int> > &trust)
	{
		int len = trust.size();
		vector<int> count(n);

		for (auto &val : trust)
		{
			count[val[0] - 1] = INT_MIN;
			count[val[1] - 1]++;
		}

		for (int i = 0; i < n; i++)
			if (count[i] == n - 1)
				return i + 1;

		return -1;
	}
};