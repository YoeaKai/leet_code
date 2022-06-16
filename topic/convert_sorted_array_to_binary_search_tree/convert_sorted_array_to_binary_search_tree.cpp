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
    using iter = vector<int>::const_iterator;

    TreeNode *sortedArrayToBST(vector<int> &nums)
    {
        return buildBST(nums.begin(), nums.end());
    }

    TreeNode *buildBST(iter b, iter e)
    {
        if (b >= e)
            return nullptr;

        iter m = b + (e - b) / 2;

        TreeNode *node = new TreeNode(*m);
        node->left = buildBST(b, m);
        node->right = buildBST(m + 1, e);

        return node;
    }

    TreeNode *sortedArrayToBST2(vector<int> &nums)
    {
        int len = nums.size();

        if (len == 0)
        {
            return NULL;
        }

        if (len == 1)
        {
            return new TreeNode(nums[0]);
        }

        int mid = len / 2;

        TreeNode *node = new TreeNode(nums[mid]);
        vector<int> left(nums.begin(), nums.begin() + mid);
        vector<int> right(nums.begin() + mid + 1, nums.end());
        node->left = sortedArrayToBST2(left);
        node->right = sortedArrayToBST2(right);

        return node;
    }
};
