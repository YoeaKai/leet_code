#include <vector>
using namespace std;

class Solution
{
public:
	int findLongestChain(vector<vector<int> > &pairs)
	{
		sort(pairs.begin(), pairs.end());

		int cur = pairs[0][1], ret = 1;

		for (auto &p : pairs)
		{
			if (p[1] < cur)
			{
				cur = p[1];
			}
			else if (p[0] > cur)
			{
				cur = p[1];
				ret++;
			}
		}

		return ret;
	}
};