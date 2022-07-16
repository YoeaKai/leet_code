#include <vector>
#include <unordered_map>
using namespace std;

class Solution
{
public:
	int maxPoints(vector<vector<int> > &points)
	{
		int result = 0;

		for (int i = 0; i < points.size(); i++)
		{
			int samePoint = 1;
			unordered_map<double, int> map;

			for (int j = i + 1; j < points.size(); j++)
				if (points[i][0] == points[j][0] && points[i][1] == points[j][1])
					samePoint++;
				else if (points[i][0] == points[j][0])
					map[INT_MAX]++;
				else
					map[double(points[i][1] - points[j][1]) / double(points[i][0] - points[j][0])]++;

			int localMax = 0;

			for (auto &val : map)
				localMax = max(localMax, val.second);

			localMax += samePoint;
			result = max(result, localMax);
		}

		return result;
	}
};