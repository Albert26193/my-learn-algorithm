#include "../cpp_headers/lc_headers.h"

using namespace std;

class Solution {
   public:
    int findLengthOfShortestSubarray(vector<int>& arr) {
        int n = arr.size();
        vector<int> prefix;
        vector<int> suffix;
        for (int i = 0; i < n; i++) {
            if (i == 0 || arr[i - 1] <= arr[i]) {
                prefix.push_back(arr[i]);
            } else {
                break;
            }
        }

        int suffixStartIndex = 0;
        for (int i = n - 1; i >= 0; i--) {
            if (i == n - 1 || arr[i] <= arr[i + 1]) {
                suffix.push_back(arr[i]);
                suffixStartIndex = i;
            } else {
                break;
            }
        }
        reverse(suffix.begin(), suffix.end());
        int ans = min(n - prefix.size(), n - suffix.size());
        int prefixPtr = 0, suffixPtr = suffixStartIndex;
        while (prefixPtr < prefix.size() && prefixPtr < suffixPtr && suffixPtr < n) {
            if (prefix[prefixPtr] <= arr[suffixPtr]) {
                ans = min(ans, (int)(suffixPtr - prefixPtr - 1));
                // cout << prefixPtr << " " << suffixPtr << endl;
                // cout << suffixPtr - prefixPtr - 1 << endl;
                prefixPtr += 1;
            } else {
                suffixPtr += 1;
            }
        }
        return ans;
    }
};

int main() {
    Solution s;
    // {1, 2, 3, 10, 4, 2, 3, 5}
    //  0  1  2   3  4  5  6  7
    //  *               *
    //     *            *
    vector<int> arr = {1, 2, 3, 10, 4, 2, 3, 5};
    cout << s.findLengthOfShortestSubarray(arr) << endl;
    return 0;
}
