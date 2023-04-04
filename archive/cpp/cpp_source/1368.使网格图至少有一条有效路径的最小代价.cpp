/*
 * @lc app=leetcode.cn id=1368 lang=cpp
 *
 * [1368] 使网格图至少有一条有效路径的最小代价
 */
#include <bits/stdc++.h>

#include "../cpp_headers/lc_headers.h"

using namespace std;

// @lc code=start
class Solution {
   private:
    int rowCount = 0, colCount = 0;
    int dx[4] = {0, 0, 1, -1};
    int dy[4] = {1, -1, 0, 0};
    bool checkRange(int m, int n) {
        if (0 <= m && m < rowCount && 0 <= n && n < colCount) {
            return true;
        }

        return false;
    }

   public:
    int minCost(vector<vector<int>>& grid) {
        rowCount = grid.size();
        if (rowCount == 0) {
            return -1;
        }

        colCount = grid[0].size();
        if (colCount == 0) {
            return -1;
        }

        deque<pair<int, int>> q;
        q.push_back({0, 0});

        vector<vector<int>> dist(rowCount, vector<int>(colCount, 0x3f3f3f3f));

        dist[0][0] = 0;

        while (!q.empty()) {
            int curX = q.front().first;
            int curY = q.front().second;
            q.pop_front();

            for (int i = 0; i < 4; i++) {
                int nextX = curX + dx[i];
                int nextY = curY + dy[i];

                int cost = (int)(grid[curX][curY] != (i + 1));

                if (!checkRange(nextX, nextY)) {
                    continue;
                }
                if (!(dist[curX][curY] + cost < dist[nextX][nextY])) {
                    continue;
                }

                if (cost == 0) {
                    q.push_front({nextX, nextY});
                } else {
                    q.push_back({nextX, nextY});
                }

                dist[nextX][nextY] = dist[curX][curY] + cost;
            }
        }

        return dist[rowCount - 1][colCount - 1];
    }
};
// @lc code=end

int main() {
    Solution s;
    vector<vector<int>> input = {
        {1, 1, 1, 1},
        {2, 2, 2, 2},
        {1, 1, 1, 1},
        {2, 2, 2, 2}};

    int cnt = s.minCost(input);
    return 0;
}
