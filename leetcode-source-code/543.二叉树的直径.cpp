/*
 * @lc app=leetcode.cn id=543 lang=cpp
 *
 * [543] 二叉树的直径
 */

#include <bits/stdc++.h>

#include "../cpp_headers/lc_headers.h"
using namespace std;

// @lc code=start
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
leetcode-source-code/851.喧闹和富有.cpp *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */

class Solution {
   public:
    int getDepth(TreeNode* root, int& ans) {
        if (!root) {
            return 0;
        }

        int leftDepth = getDepth(root->left, ans);
        int rightDepth = getDepth(root->right, ans);
        ans = max(ans, leftDepth + rightDepth + 1);
        return max(leftDepth, rightDepth) + 1;
    }

    int diameterOfBinaryTree(TreeNode* root) {
        int ans = 1;
        getDepth(root, ans);
        return ans - 1;
    }
};
// @lc code=end
