#include <vector>
using namespace std;

class Solution
{
public:
	int
	minMoves2(vector<int> &nums)
	{
		sort(nums.begin(), nums.end());

		int sum = 0, n = nums.size();
		int mid = nums[n / 2];

		for (int i = 0; i < n; i++)
		{
			sum += abs(nums[i] - mid);
		}

		return sum;
	}
};