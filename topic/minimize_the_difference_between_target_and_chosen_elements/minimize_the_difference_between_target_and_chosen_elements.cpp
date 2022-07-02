#include <iostream>
#include <vector>
using namespace std;

class Solution
{
public:
	int minimizeTheDifference(vector<vector<int> > &mat, int target)
	{
		bitset<4901> p(1);

		for (auto &r : mat)
		{
			bitset<4901> tmp;

			for (auto &i : r)
				tmp |= p << i;

			swap(p, tmp);
		}

		int result = INT_MAX;

		for (int i = 0; i < 4901; ++i)
			if (p[i])
			{
				result = min(result, abs(i - target));

				if (i > target)
					break;
			}

		return result;
	}
};
