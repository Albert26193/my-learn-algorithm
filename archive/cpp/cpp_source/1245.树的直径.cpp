#include <bits/stdc++.h>

#include "../cpp_headers/lc_headers.h"

using namespace std;

class Solution {
   public:
    void dfs(vector<vector<int>>& adj, int currNode, vector<int>& fa, vector<int>& depth) {
        for (int i = 0, len = adj[currNode].size(); i < len; i++) {
            int nextNode = adj[currNode][i];
            if (depth[nextNode] > -1) {
                continue;
            }

            depth[nextNode] = depth[currNode] + 1;
            fa[nextNode] = currNode;
            dfs(adj, nextNode, fa, depth);
        }
    }

    int findFurthestNode(vector<vector<int>>& adj, int currNode, vector<int>& fa) {
        int nodeCount = adj.size();
        vector<int> depth(nodeCount, -1);
        int maxDepth = -1;
        int furthestNode = -1;

        fa = vector<int>(nodeCount, -1);
        depth[currNode] = 0;
        dfs(adj, currNode, fa, depth);
        for (int i = 0; i < nodeCount; i++) {
            if (depth[i] > maxDepth) {
                maxDepth = depth[i];
                furthestNode = i;
            }
        }

        return furthestNode;
    }

    int treeDiameter(vector<vector<int>>& edges) {
        int nodeCount = edges.size() + 1;
        if (nodeCount == 1) {
            return 0;
        }
        vector<vector<int>> adj(nodeCount, vector<int>(0));
        for (int i = 0, len = edges.size(); i < len; i++) {
            adj[edges[i][1]].push_back(edges[i][0]);
            adj[edges[i][0]].push_back(edges[i][1]);
        }

        int randomNode = rand() % nodeCount;

        vector<int> fa(nodeCount, -1);

        int xNode = findFurthestNode(adj, randomNode, fa);
        fa = vector<int>(nodeCount, -1);
        int yNode = findFurthestNode(adj, xNode, fa);

        int fatherNode = fa[yNode];
        vector<int> path = {yNode};
        while (fatherNode != -1) {
            // cout << fatherNode << endl;
            path.push_back(fatherNode);
            fatherNode = fa[fatherNode];
        }

        return path.size() - 1;
    }
};