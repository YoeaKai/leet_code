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
	void findRightNode(TreeNode *node, vector<int> &result, int level)
	{
		if (node == nullptr)
			return;

		if (result.size() < level)
			result.push_back(node->val);

		findRightNode(node->right, result, level + 1);
		findRightNode(node->left, result, level + 1);
	}

	vector<int> rightSideView(TreeNode *root)
	{
		vector<int> result;
		findRightNode(root, result, 1);
		return result;
	}
};