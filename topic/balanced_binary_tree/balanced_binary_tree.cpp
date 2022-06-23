#include <iostream>
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
	int depth(TreeNode *root)
	{
		if (root == nullptr)
			return 0;

		int leftDepth = depth(root->left);
		if (leftDepth == -1)
			return -1;

		int rightDepth = depth(root->right);
		if (rightDepth == -1)
			return -1;

		if (abs(leftDepth - rightDepth) > 1)
			return -1;

		return max(leftDepth, rightDepth) + 1;
	}

	bool isBalanced(TreeNode *root)
	{
		return depth(root) != -1;
	}
};