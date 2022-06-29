#include <vector>
using namespace std;

class Solution
{
public:
	vector<vector<int> > reconstructQueue(vector<vector<int> > &people)
	{
		sort(people.begin(), people.end(), [](vector<int> &p1, vector<int> &p2)
			 { return p1[0] > p2[0] || (p1[0] == p2[0] && p1[1] < p2[1]); });

		vector<vector<int> > result;

		for (auto person : people)
		{
			result.insert(result.begin() + person[1], person);
		}

		return result;
	}
};