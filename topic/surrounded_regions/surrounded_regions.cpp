#include <vector>
using namespace std;
class Solution
{
public:
	// i = hight, j = width, m = max hight, n = max width
	void findNonFlipped(vector<vector<char> > &board, int i, int j, int m, int n)
	{
		if (i < 0 or j < 0 or i >= m or j >= n or board[i][j] != 'O')
			return;
		board[i][j] = '#';
		findNonFlipped(board, i - 1, j, m, n);
		findNonFlipped(board, i + 1, j, m, n);
		findNonFlipped(board, i, j - 1, m, n);
		findNonFlipped(board, i, j + 1, m, n);
	}

	void solve(vector<vector<char> > &board)
	{
		int m = board.size();
		int n = board[0].size();

		//Moving over firts and last column
		for (int i = 0; i < m; i++)
		{
			if (board[i][0] == 'O')
				findNonFlipped(board, i, 0, m, n);
			if (board[i][n - 1] == 'O')
				findNonFlipped(board, i, n - 1, m, n);
		}

		//Moving over first and last row
		for (int j = 0; j < n; j++)
		{
			if (board[0][j] == 'O')
				findNonFlipped(board, 0, j, m, n);
			if (board[m - 1][j] == 'O')
				findNonFlipped(board, m - 1, j, m, n);
		}

		for (int i = 0; i < m; i++)
			for (int j = 0; j < n; j++)
			{
				if (board[i][j] == 'O')
					board[i][j] = 'X';
				if (board[i][j] == '#')
					board[i][j] = 'O';
			}
	}
};