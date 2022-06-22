#include <vector>
#include <queue>
using namespace std;

class Solution
{
public:
	// STL
	int findKthLargest(vector<int> &nums, int k)
	{
		// sort(nums.begin(), nums.end());
		// return nums[nums.size() - k];

		// nth_element(nums.begin(), nums.begin() + k - 1, nums.end(), greater<int>());

		partial_sort(nums.begin(), nums.begin() + k, nums.end(), greater<int>());
		return nums[k - 1];
	}

	// min-heap using priority_queue
	int findKthLargest2(vector<int> &nums, int k)
	{
		priority_queue<int, vector<int>, greater<int> > pq;

		for (int num : nums)
		{
			pq.push(num);
			if (pq.size() > k)
				pq.pop();
		}

		return pq.top();
	}
};