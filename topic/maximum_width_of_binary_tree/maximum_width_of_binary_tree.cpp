#include <queue>
using namespace std;

struct TreeNode
{
	int val;
	TreeNode *left;
	TreeNode *right;
	TreeNode() : val(0), left(nullptr), right(nullptr) {}
	TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
	TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
};

class Solution
{
public:
	int widthOfBinaryTree(TreeNode *root)
	{
		// pair<node, order>
		queue<pair<TreeNode *, long> > q;
		q.push({root, 0});
		int ret = INT_MIN;

		while (!q.empty())
		{
			int len = q.size();
			int levelMin = q.front().second;
			long dif;

			for (int i = 0; i < len; ++i)
			{
				TreeNode *node = q.front().first;
				dif = q.front().second - levelMin;
				q.pop();

				if (node->left)
					q.push({node->left, 2 * dif + 1});
				if (node->right)
					q.push({node->right, 2 * dif + 2});
			}

			ret = max(ret, int(dif) + 1);
		}

		return ret;
	}
};