/*
 * @lc app=leetcode.cn id=1590 lang=cpp
 *
 * [1590] 使数组和能被 P 整除
 */

#include <bits/stdc++.h>

#include "../headers/lc-all-header.h"
using namespace std;

// @lc code=start
class Solution {
   public:
    int minSubarray(vector<int>& nums, int p) {
        int numsLength = nums.size();

        vector<int> prefixSum(numsLength + 1, 0);
        for (int i = 1; i < numsLength + 1; i++) {
            prefixSum[i] = (prefixSum[i - 1] + nums[i - 1]) % p;
        }

        int targetSubSum = prefixSum[numsLength] % p;
        if (targetSubSum == 0) {
            return 0;
        }

        int minLength = 0x3f3f3f3f;
        unordered_map<int, int> posRecord;
        for (int i = 0; i < numsLength + 1; i++) {
            posRecord[prefixSum[i]] = i;
            if (posRecord.count((prefixSum[i] - targetSubSum + p) % p)) {
                minLength = min(minLength, i - posRecord[(prefixSum[i] - targetSubSum + p) % p]);
                ;
            }
        }

        // cout << minLength << endl;
        return minLength < numsLength ? minLength : -1;
    }
};
// @lc code=end

int main() {
    Solution s;
    vector<int> nums = {6, 3, 5, 2};
    int p = 9;
    int len = s.minSubarray(nums, p);
    cout << len << endl;
    return 0;
}