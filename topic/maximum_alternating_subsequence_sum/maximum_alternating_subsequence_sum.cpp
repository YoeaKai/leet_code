#include <vector>
using namespace std;

class Solution
{
public:
	long long maxAlternatingSum(vector<int> &nums)
	{
		long long odd = 0, even = 0;

		for (int &num : nums)
		{
			// long long tmp = even;
			even = max(even, odd + num);
			// odd = max(odd, tmp - num);
			// This odd change is equal to the content of the comment.
			odd = even - num;
		}

		return even;
	}
};
