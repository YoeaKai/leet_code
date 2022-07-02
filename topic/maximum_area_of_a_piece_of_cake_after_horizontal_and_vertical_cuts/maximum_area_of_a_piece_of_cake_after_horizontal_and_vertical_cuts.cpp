#include <vector>
using namespace std;

class Solution
{
	int getMax(int &edge, vector<int> &cuts)
	{
		sort(cuts.begin(), cuts.end());

		int len = cuts.size();
		int maxVal = max(edge - cuts[len - 1], cuts[0]);

		for (int i = 1; i < len; i++)
			maxVal = max(maxVal, cuts[i] - cuts[i - 1]);

		return maxVal;
	}

public:
	int maxArea(int h, int w, vector<int> &horizontalCuts, vector<int> &verticalCuts)
	{
		return (long)getMax(h, horizontalCuts) * getMax(w, verticalCuts) % 1000000007;
	}
};
