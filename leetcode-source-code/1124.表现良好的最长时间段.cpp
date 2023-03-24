/*
 * @lc app=leetcode.cn id=1124 lang=cpp
 *
 * [1124] 表现良好的最长时间段
 */
#include <bits/stdc++.h>

#include "../headers/lc-all-header.h"
using namespace std;

// @lc code=start
class Solution {
   public:
    int longestWPI(vector<int>& hours) {
        int n = hours.size();
        vector<int> summ(n + 1, 0);
        stack<int> sta;
        sta.push(0);

        for (int i = 1; i <= n; i++) {
            summ[i] = summ[i - 1] + (hours[i - 1] > 8 ? 1 : -1);
            if (summ[sta.top()] > summ[i]) {
                sta.push(i);
            }
        }

        int res = 0;

        for (int r = n; r >= 1; r--) {
            while (sta.size() && summ[sta.top()] < summ[r]) {
                res = max(res, r - sta.top());
                sta.pop();
            }
        }

        return res;
    }
};
// @lc code=end

int main() {
    Solution s;
    vector<int> hours = {6, 9, 9};
    cout << s.longestWPI(hours) << endl;
    return 0;
}
