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
	TreeNode *build(vector<int> &preorder, vector<int> &inorder, int stop, int &prePos, int &inPos)
	{
		if (prePos >= preorder.size())
			return nullptr;

		if (inorder[inPos] == stop)
		{
			inPos++;
			return nullptr;
		}

		TreeNode *node = new TreeNode(preorder[prePos++]);
		node->left = build(preorder, inorder, node->val, prePos, inPos);
		node->right = build(preorder, inorder, stop, prePos, inPos);

		return node;
	}

	TreeNode *buildTree(vector<int> &preorder, vector<int> &inorder)
	{
		int inPos = 0, prePos = 0;
		return build(preorder, inorder, INT_MIN, prePos, inPos);
	}
};