/*
 * @lc app=leetcode.cn id=475 lang=cpp
 *
 * [475] 供暖器
 */

#include <bits/stdc++.h>

#include "../cpp_headers/lc_headers.h"
using namespace std;

// @lc code=start
class Solution {
   private:
    const int LEFT = 1;
    const int RIGHT = 2;

   public:
    int findFirstHeaterPosition(int direction, int currentHouse, vector<int>& heaters) {
        int heatersCount = heaters.size();
        if (direction == LEFT) {
            int left = 0, right = heatersCount - 1;
            while (left < right) {
                int mid = (left + right + 1) / 2;
                if (heaters[mid] <= currentHouse) {
                    left = mid;
                } else {
                    right = mid - 1;
                }
            }

            return (heaters[right] <= currentHouse) ? (heaters[right]) : (heaters[0]);
        } else if (direction == RIGHT) {
            int left = 0, right = heatersCount - 1;
            while (left < right) {
                int mid = (left + right) / 2;
                if (heaters[mid] >= currentHouse) {
                    right = mid;
                } else {
                    left = mid + 1;
                }
            }

            return (heaters[left] >= currentHouse) ? (heaters[left]) : (heaters[heatersCount - 1]);
        }

        return -1;
    };

    int findRadius(vector<int>& houses, vector<int>& heaters) {
        sort(houses.begin(), houses.end());
        sort(heaters.begin(), heaters.end());
        int minRadius = -1;
        for (int currentHouse : houses) {
            int firstLeftHeaterPosition = findFirstHeaterPosition(LEFT, currentHouse, heaters);
            int firstRightHeaterPosition = findFirstHeaterPosition(RIGHT, currentHouse, heaters);
            int lessRadius = min(abs(currentHouse - firstRightHeaterPosition), abs(currentHouse - firstLeftHeaterPosition));
            minRadius = max(minRadius, lessRadius);
        }

        return minRadius;
    };
};
// @lc code=end

int main() {
    Solution s;
    vector<int> houses = {1, 2, 3, 4};
    vector<int> heaters = {1, 4};
    return 0;
}