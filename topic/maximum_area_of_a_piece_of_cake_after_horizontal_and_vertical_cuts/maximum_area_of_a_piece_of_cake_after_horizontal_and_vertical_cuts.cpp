#include <vector>
using namespace std;

class Solution
{
public:
	int maxArea(int h, int w, vector<int> &horizontalCuts, vector<int> &verticalCuts)
	{
		sort(horizontalCuts.begin(), horizontalCuts.end());
		sort(verticalCuts.begin(), verticalCuts.end());

		int heightMax = horizontalCuts[0], weightMax = verticalCuts[0], i = 1;

		for (; i < horizontalCuts.size(); i++)
		{
			horizontalCuts[i - 1] = horizontalCuts[i] - horizontalCuts[i - 1];

			if (horizontalCuts[i - 1] > heightMax)
				heightMax = horizontalCuts[i - 1];
		}

		if (h - horizontalCuts[i - 1] > heightMax)
			heightMax = h - horizontalCuts[i - 1];

		for (i = 1; i < verticalCuts.size(); i++)
		{
			verticalCuts[i - 1] = verticalCuts[i] - verticalCuts[i - 1];

			if (verticalCuts[i - 1] > weightMax)
				weightMax = verticalCuts[i - 1];
		}

		if (w - verticalCuts[i - 1] > weightMax)
			weightMax = w - verticalCuts[i - 1];

		return ((long)heightMax * weightMax) % 1000000007;
	}
};