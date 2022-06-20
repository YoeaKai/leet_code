#include <vector>
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
	vector<int> findMode(TreeNode *root)
	{
		vector<int> ret;
		int curr = INT_MIN, times = 0, max = 0;
		inorderSearch(ret, root, curr, times, max);
		return ret;
	}

	void inorderSearch(vector<int> &ret, TreeNode *node, int &curr, int &times, int &max)
	{
		if (!node)
		{
			return;
		}

		inorderSearch(ret, node->left, curr, times, max);

		if (node->val > curr)
		{
			curr = node->val;
			times = 1;
		}
		else // node->val == curr
		{
			times++;
		}

		if (times > max)
		{
			max = times;
			if (ret[0] != curr)
			{
				ret.clear();
				ret.push_back(node->val);
			}
		}
		else if (times == max)
		{
			ret.push_back(node->val);
		}

		inorderSearch(ret, node->right, curr, times, max);
	}
};