#include <vector>
#include <cstring>

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode() : val(0), left(nullptr), right(nullptr) {}
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
    TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
};

class Solution {
public:
    TreeNode* sortedArrayToBST(std::vector<int>& nums) {
        int len = nums.size();

        if (len == 1){
            return new TreeNode(nums[0]);
        }
        
        int mid = len/2;

        TreeNode* node = new TreeNode(nums[mid]);

        vector<int> a (nums.begin(), nums.begin()+mid);
        vector<int> b (nums.begin()+mid+1, nums.end());

        node -> left = sortedArrayToBST(a);
        node -> right = sortedArrayToBST(b);

        return node;
    }
};
