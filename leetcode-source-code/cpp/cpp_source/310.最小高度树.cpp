/*
 * @lc app=leetcode.cn id=310 lang=cpp
 *
 * [310] 最小高度树
 */

#include <bits/stdc++.h>

#include "../cpp_headers/lc_headers.h"

using namespace std;

// @lc code=start
class Solution {
   public:
    void dfs(int curPoint, vector<int>& fa, vector<vector<int>>& adj, vector<int>& depth) {
        for (int i = 0, len = adj[curPoint].size(); i < len; i++) {
            int nextPoint = adj[curPoint][i];

            // 如果这个点被遍历过，那么就需要放弃
            // 遍历到最远的叶子节点，就不可能往回走，因此，最后一条DFS路径，就是最长的一条路径
            if (depth[nextPoint] > -1) {
                continue;
            }

            depth[nextPoint] = depth[curPoint] + 1;
            // cout << "ffffff" << endl;
            fa[nextPoint] = curPoint;
            dfs(nextPoint, fa, adj, depth);
        }
    }

    int findLongestPoint(int startPoint, vector<vector<int>>& adj, vector<int>& fa) {
        // LC::PrintVector(fa);
        int nodeCount = adj.size();
        vector<int> depth(nodeCount, -1);

        fa = vector<int>(nodeCount, -1);
        depth[startPoint] = 0;
        dfs(startPoint, fa, adj, depth);
        int maxDepth = 0;
        int furthestPoint = -1;
        for (int i = 0; i < nodeCount; i++) {
            if (depth[i] > maxDepth) {
                furthestPoint = i;
                maxDepth = depth[i];
            }
        }

        return furthestPoint;
    }

    vector<int> findMinHeightTrees(int n, vector<vector<int>>& edges) {
        if (n == 2) {
            return {1};
        }
        int randomPoint = rand() % n;
        vector<vector<int>> adj(n, vector<int>(0));
        vector<int> fa(n, -1);
        for (int i = 0, len = edges.size(); i < len; i++) {
            adj[edges[i][0]].push_back(edges[i][1]);
            adj[edges[i][1]].push_back(edges[i][0]);
        }

        int xNode = findLongestPoint(randomPoint, adj, fa);
        int yNode = findLongestPoint(xNode, adj, fa);

        int fatherNode = fa[yNode];
        vector<int> path(1, yNode);
        while (fatherNode != -1) {
            // cout << fatherNode << endl;
            path.push_back(fatherNode);
            fatherNode = fa[fatherNode];
        }

        // LC::PrintVector(path);
        // LC::PrintMatrix(adj);

        int pathLength = path.size();

        if (pathLength % 2) {
            return {path[pathLength / 2]};
        } else {
            return {{path[pathLength / 2 - 1], path[pathLength / 2]}};
        }
    }
};

// @lc code=end
