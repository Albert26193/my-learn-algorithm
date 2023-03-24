/*
 * @lc app=leetcode.cn id=994 lang=cpp
 *
 * [994] 腐烂的橘子
 */

#include <bits/stdc++.h>

#include "../headers/lc-all-header.h"
using namespace std;

// @lc code=start
class Solution {
   private:
    int dx[4] = {1, 0, -1, 0};
    int dy[4] = {0, 1, 0, -1};
    int row = 0, col = 0;

   public:
    bool checkRange(int x, int y) {
        return (x >= 0 && x < row && y >= 0 && y < col);
    }

    int orangesRotting(vector<vector<int>>& grid) {
        row = grid.size();
        if (!row) {
            return 0;
        }

        col = grid[0].size();
        if (!col) {
            return 0;
        }

        queue<pair<int, int>> q;
        for (int i = 0; i < row; i++) {
            for (int j = 0; j < col; j++) {
                if (grid[i][j] == 2) {
                    q.push({i, j});
                }
            }
        }

        if (q.size() == 0) {
            for (int i = 0; i < row; i++) {
                for (int j = 0; j < col; j++) {
                    if (grid[i][j] == 1) {
                        return -1;
                    }
                }
            }
            return 0;
        }

        int minutes = -1;
        while (!q.empty()) {
            int queueLen = q.size();
            minutes += 1;
            for (int i = 0; i < queueLen; i++) {
                pair<int, int> curPos = q.front();
                q.pop();
                for (int i = 0; i < 4; i++) {
                    int nextX = curPos.first + dx[i];
                    int nextY = curPos.second + dy[i];
                    if (checkRange(nextX, nextY) && grid[nextX][nextY] == 1) {
                        q.push({nextX, nextY});
                        grid[nextX][nextY] = 2;
                    }
                }
            }
        }

        for (int i = 0; i < row; i++) {
            for (int j = 0; j < col; j++) {
                if (grid[i][j] == 1) {
                    return -1;
                }
            }
        }

        return minutes;
    }
};
// @lc code=end

int main() {
    Solution s;
    vector<vector<int>> grid = {
        {2, 1, 1},
        {1, 1, 0},
        {0, 1, 1},
    };

    int res = s.orangesRotting(grid);
    cout << res << endl;
    return 0;
}
