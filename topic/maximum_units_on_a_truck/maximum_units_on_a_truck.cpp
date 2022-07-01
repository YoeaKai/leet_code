#include <vector>
using namespace std;

class Solution
{
public:
	static bool compare(vector<int> &v1, vector<int> &v2)
	{
		return v1[1] > v2[1];
	}

	int maximumUnits(vector<vector<int> > &boxTypes, int truckSize)
	{
		sort(boxTypes.begin(), boxTypes.end(), compare);

		int count = 0, sum = 0;

		for (auto &box : boxTypes)
		{
			count += box[0];

			if (count > truckSize)
				return sum + (truckSize - (count - box[0])) * box[1];

			sum += box[0] * box[1];
		}

		return sum;
	}
};