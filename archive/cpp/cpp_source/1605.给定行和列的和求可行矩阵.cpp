/*
 * @lc app=leetcode.cn id=1605 lang=cpp
 *
 * [1605] 给定行和列的和求可行矩阵
 */

#include <bits/stdc++.h>

#include "../cpp_headers/lc_headers.h"
using namespace std;

// @lc code=start
class Solution {
   public:
    vector<vector<int>> restoreMatrix(vector<int>& rowSum, vector<int>& colSum) {
        int rowCount = rowSum.size();
        int colCount = colSum.size();

        if (!rowCount || !colCount) {
            return {{}};
        }

        vector<vector<int>> resultMatrix(rowCount, vector<int>(colCount, -1));

        for (int x = 0; x < rowCount; x++) {
            for (int y = 0; y < colCount; y++) {
                resultMatrix[x][y] = min(rowSum[x], colSum[y]);
                rowSum[x] -= resultMatrix[x][y];
                colSum[y] -= resultMatrix[x][y];
            }
        }

        return resultMatrix;
    }
};
// @lc code=end

int main() {
    Solution s;
    vector<int> rowSum = {3, 8};
    vector<int> colSum = {4, 7};

    lcPrintMatrix(s.restoreMatrix(rowSum, colSum));

    return 0;
}
