/*
 * @lc app=leetcode.cn id=851 lang=cpp
 *
 * [851] 喧闹和富有
 */

#include <bits/stdc++.h>

#include "../headers/lc-all-header.h"
using namespace std;

// @lc code=start
class Solution {
   public:
    vector<int> loudAndRich(vector<vector<int>>& richer, vector<int>& quiet) {
        int personCount = quiet.size();
        vector<int> inDegree(personCount, 0);
        vector<vector<int>> adj(personCount, vector<int>(0));
        for (int i = 0, len = richer.size(); i < len; i++) {
            adj[richer[i][0]].push_back(richer[i][1]);
            inDegree[richer[i][1]] += 1;
        }

        queue<int> personQueue;
        for (int i = 0; i < personCount; i++) {
            if (inDegree[i] == 0) {
                personQueue.push(i);
            }
        }

        vector<int> ans(personCount, 0);
        for (int i = 0; i < personCount; i++) {
            ans[i] = i;
        }

        while (!personQueue.empty()) {
            int curPerson = personQueue.front();
            personQueue.pop();

            for (int nextPerson : adj[curPerson]) {
                inDegree[nextPerson] -= 1;
                if (quiet[ans[curPerson]] < quiet[ans[nextPerson]]) {
                    ans[nextPerson] = ans[curPerson];
                }
                if (inDegree[nextPerson] == 0) {
                    personQueue.push(nextPerson);
                } 
            }
        }

        return ans;
    }
};
// @lc code=end
int main() {
    Solution s;
    vector<vector<int>> richer = {
        {1, 0},
        {2, 1},
        {3, 1},
        {3, 7},
        {4, 3},
        {5, 3},
        {6, 3},
    };
    vector<int> quiet = {3, 2, 5, 4, 6, 1, 7, 0};
    s.loudAndRich(richer, quiet);
    return 0;
}
