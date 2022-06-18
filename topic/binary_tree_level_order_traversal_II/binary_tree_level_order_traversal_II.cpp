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
	int depth(TreeNode *node)
	{
		return !node ? 0 : max(depth(node->left), depth(node->right)) + 1;
	}
	void levelOrder(vector<vector<int> > &ret, TreeNode *node, int level)
	{
		if (!node)
		{
			return;
		}
		ret[level].push_back(node->val);
		levelOrder(ret, node->left, level - 1);
		levelOrder(ret, node->right, level - 1);
	}
	vector<vector<int> > levelOrderBottom(TreeNode *root)
	{
		int d = depth(root);
		vector<vector<int> > ret(d, vector<int>{});
		levelOrder(ret, root, d - 1);
		return ret;
	}
};