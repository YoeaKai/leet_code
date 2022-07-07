#include <vector>
using namespace std;

class Solution
{
public:
	int maxScore(vector<int> &cardPoints, int k)
	{
		int sum = 0;

		for (int i = 0; i < k; i++)
			sum += cardPoints[i];

		int result = sum, n = cardPoints.size() - 1;

		for (int i = k - 1; i >= 0; i--)
		{
			sum += cardPoints[n - (k - 1 - i)] - cardPoints[i];

			if (sum > result)
				result = sum;
		}

		return result;
	}
};