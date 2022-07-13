#include <vector>
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
	vector<vector<int> > levelOrder(TreeNode *root)
	{
		vector<vector<int> > result;
		queue<TreeNode *> queue;

		if (root == nullptr)
			return result;

		queue.push(root);

		while (!queue.empty())
		{
			result.push_back(vector<int>{});
			int n = queue.size();

			for (int i = 0; i < n; i++)
			{
				result.back().push_back(queue.front()->val);

				if (queue.front()->left != nullptr)
					queue.push(queue.front()->left);

				if (queue.front()->right != nullptr)
					queue.push(queue.front()->right);

				queue.pop();
			}
		}

		return result;
	}
};