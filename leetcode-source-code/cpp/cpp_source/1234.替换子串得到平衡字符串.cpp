/*
 * @lc app=leetcode.cn id=1234 lang=cpp
 *
 * [1234] 替换子串得到平衡字符串
 */

#include "../cpp_headers/lc_headers.h"

using namespace std;
// @lc code=start
class Solution {
   public:
    int balancedString(string s) {
        int len = s.size();
        if (len % 4) {
            return -1;
        }

        int subLen = len / 4;
        const string mapString = "QWER";
        vector<int> rec(4, 0);

        for (char ch : s) {
            rec[mapString.find(ch)] += 1;
        }

        if (rec[0] == subLen && rec[1] == subLen && rec[2] == subLen && rec[3] == subLen) {
            return 0;
        }

        int left = 0, right = 0;
        int ans = 1e9;
        for (right = 0; right < len; right++) {
            rec[mapString.find(s[right])] -= 1;
            while (left <= right && rec[1] <= subLen && rec[2] <= subLen && rec[3] <= subLen && rec[0] <= subLen) {
                ans = min(ans, right - left + 1);
                rec[mapString.find(s[left])] += 1;
                left += 1;
            }
        }

        return ans;
    }
};
// @lc code=end
