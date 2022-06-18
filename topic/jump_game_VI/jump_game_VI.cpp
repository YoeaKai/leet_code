#include <deque>
#include <vector>
using namespace std;

class Solution
{
public:
	int maxResult(vector<int> &nums, int k)
	{
		long curr = 0;
		deque<int> dq;

		for (int i = nums.size() - 1; i >= 0; i--)
		{
			curr = nums[i] + (dq.empty() ? 0 : nums[dq.front()]); // Get current best.

			// We remove all the smaller results.
			while (!dq.empty() && curr > nums[dq.back()])
				dq.pop_back();
			dq.push_back(i);

			// Erase all the indices in deque that are greater than or equal to i+k.
			if (dq.front() >= i + k)
				dq.pop_front();
			nums[i] = curr; // Use input array as auxillary array to store the best results.
		}
		return curr;
	}
};