#include <algorithm>
#include <iostream>
#include <sstream>
#include <vector>

using namespace std;

struct TreeNode {
    int val;
    TreeNode* left;
    TreeNode* right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

TreeNode* buildBST(vector<int>& nums, int start, int end) {
    if (start > end) return NULL;

    int mid = start + (end - start) / 2;
    TreeNode* root = new TreeNode(nums[mid]);

    root->left = buildBST(nums, start, mid - 1);
    root->right = buildBST(nums, mid + 1, end);

    return root;
}

void check(TreeNode* root, int target) {
    if (!root) {
        cout << "N";
        return;
    }

    if (target > root->val) {
        if (root->right) cout << "R";
        check(root->right, target);
    } else if (target == root->val) {
        cout << "Y";
    } else {
        if (root->left) cout << "L";
        check(root->left, target);
    }
}

int main() {
    vector<int> nums;
    string line;
    getline(cin, line);
    istringstream is(line);
    int num;
    while (is >> num) {
        nums.push_back(num);
    }

    int k;
    cin >> k;

    sort(nums.begin(), nums.end());

    TreeNode* root = buildBST(nums, 0, nums.size() - 1);

    cout << "S";
    check(root, k);

    return 0;
}
