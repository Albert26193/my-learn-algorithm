/*
 * @lc app=leetcode.cn id=962 lang=cpp
 *
 * [962] 最大宽度坡
 */
#include <bits/stdc++.h>

#include "../cpp_headers/lc_headers.h"

using namespace std;

// @lc code=start
class Solution {
   public:
    int maxWidthRamp(vector<int>& nums) {
        int numsLen = nums.size();
        vector<int> sta;

        for (int i = 0; i < numsLen; i++) {
            if (sta.empty() || nums[i] < nums[sta.back()]) {
                sta.push_back(i);
            }
        }

        int maxRange = 0;
        for (int i = numsLen - 1; i >= 0; i--) {
            while (sta.size() && nums[i] >= nums[sta.back()]) {
                maxRange = max(i - sta.back(), maxRange);
                sta.pop_back();
            }
        }

        return maxRange;
    }
};
// @lc code=end

int main() {
    Solution s;
    return 0;
}
