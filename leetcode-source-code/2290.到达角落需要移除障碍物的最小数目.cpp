/*
 * @lc app=leetcode.cn id=2290 lang=cpp
 *
 * [2290] 到达角落需要移除障碍物的最小数目
 */

#include <bits/stdc++.h>

#include "../headers/lc-all-header.h"
using namespace std;

// @lc code=start
class Solution {
   private:
    int rowCount = 0, colCount = 0;
    int directionX[4] = {1, 0, -1, 0};
    int directionY[4] = {0, 1, 0, -1};

   public:
    bool checkRange(int currentX, int currentY) {
        if (0 <= currentX && currentX < rowCount && 0 <= currentY && currentY < colCount) {
            return true;
        }

        return false;
    }
    int minimumObstacles(vector<vector<int>>& grid) {
        int obstaclesCount = 0;
        rowCount = grid.size();
        if (!rowCount) {
            return -1;
        }
        colCount = grid[0].size();
        if (!colCount) {
            return -1;
        }

        int startX = 0, startY = 0;
        deque<pair<int, int>> positionQueue;
        positionQueue.push_back({startX, startY});

        // 松弛操作
        vector<vector<int>> dist(rowCount, vector<int>(colCount, 0x3f3f3f3f));
        dist[startX][startY] = 0;

        while (!positionQueue.empty()) {
            int currentX = positionQueue.front().first;
            int currentY = positionQueue.front().second;
            positionQueue.pop_front();

            for (int i = 0; i < 4; i++) {
                int nextX = currentX + directionX[i];
                int nextY = currentY + directionY[i];
                if (!checkRange(nextX, nextY)) {
                    continue;
                }

                int nextVal = grid[nextX][nextY];
                if (!(dist[currentX][currentY] + nextVal < dist[nextX][nextY])) {
                    continue;
                }

                if (nextVal == 0) {
                    positionQueue.push_front({nextX, nextY});
                } else if (nextVal == 1){
                    positionQueue.push_back({nextX, nextY});
                } else {
                    return -1;
                }

                dist[nextX][nextY] = dist[currentX][currentY] + grid[nextX][nextY];
            }
        }

        return dist[rowCount - 1][colCount - 1];
    }
};
// @lc code=end

int main() {
    Solution s;
    vector<vector<int>> grid = {
        {0, 1, 0, 0, 0},
        {0, 1, 0, 1, 0},
        {0, 0, 0, 1, 0},
    };

    int cnt = s.minimumObstacles(grid);
    cout << cnt << endl;
    return 0;
}
