#include <string>
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
	string ret = "";

	void dfs(TreeNode *root, string curr)
	{
		if (!root)
			return;

		curr.insert(curr.begin(), 'a' + root->val);

		if (!root->left && !root->right)
		{
			ret = ret.empty() ? curr : min(ret, curr);
		}

		dfs(root->left, curr);
		dfs(root->right, curr);

		curr.erase(curr.begin());
	}

	string smallestFromLeaf(TreeNode *root)
	{
		dfs(root, "");
		return ret;
	}

	string smallestFromLeaf2(TreeNode *root, string s = "")
	{
		s = string(1, 'a' + root->val) + s;
		// root->left == root->right represent nullptr.
		return root->left == root->right ? s : min(root->left ? smallestFromLeaf2(root->left, s) : "|", root->right ? smallestFromLeaf2(root->right, s) : "|");
	}
};