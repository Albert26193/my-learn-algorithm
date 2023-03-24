/*
 * @lc app=leetcode.cn id=1 lang=cpp
 *
 * [1] 两数之和
 */

#include <bits/stdc++.h>

using namespace std;

// @lc code=start
class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        unordered_map<int, int> mp;
        for (int i = 0, len = nums.size(); i < len; i++) {
            // if (mp.count(nums[i])) {
            //     continue;
            // }

            if (mp.count(target - nums[i])) {
                return {i, mp[target - nums[i]]};
            }

            mp[nums[i]] = i;
        }

        return {};
    }
};
// @lc code=end

