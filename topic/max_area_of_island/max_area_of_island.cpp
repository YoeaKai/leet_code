#include <vector>
using namespace std;

class Solution
{
public:
	int area(vector<vector<int> > &grid, int i, int j)
	{
		if (i < 0 || i >= grid.size() || j < 0 || j >= grid[i].size() || grid[i][j] == 0)
			return 0;

		grid[i][j] = 0;

		return area(grid, i - 1, j) + area(grid, i + 1, j) + area(grid, i, j - 1) + area(grid, i, j + 1) + 1;
	}

	int maxAreaOfIsland(vector<vector<int> > &grid)
	{
		int maxArea = 0;

		for (int i = 0; i < grid.size(); i++)
			for (int j = 0; j < grid[0].size(); j++)
				if (grid[i][j] == 1)
					maxArea = max(maxArea, area(grid, i, j));

		return maxArea;
	}
};