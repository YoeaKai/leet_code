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
	// DFS
	int minDepth(TreeNode *root)
	{
		if (!root)
			return 0;
		int L = minDepth(root->left), R = minDepth(root->right);
		return L < R && L || !R ? 1 + L : 1 + R;
	}

	// DFS 2
	int minDepth2(TreeNode *root)
	{
		if (!root)
			return 0;
		int left = minDepth(root->left);
		int right = minDepth(root->right);
		if (!root->left || !root->right)
			return !root->left ? right + 1 : left + 1;
		root->left = NULL;
		root->right = NULL;
		return min(left, right) + 1;
	}

	// BFS
	int minDepth3(TreeNode *root)
	{
		if (!root)
			return 0;

		int depth = 1;
		queue<TreeNode *> q;
		q.push(root);
		q.push(nullptr);
		TreeNode *tmp;

		while (q.size() > 1)
		{
			tmp = q.front();

			if (tmp == nullptr)
			{
				depth++;
				q.push(nullptr);
			}
			else
			{
				if (tmp->left)
					q.push(tmp->left);
				if (tmp->right)
					q.push(tmp->right);
				if (!tmp->left && !tmp->right)
					return depth;
			}

			q.pop();
		}

		return depth;
	}
};