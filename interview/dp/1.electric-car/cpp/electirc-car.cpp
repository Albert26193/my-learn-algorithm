#include <iostream>
#include <algorithm>
#include <vector>
#include <climits>

using namespace std;

int main() {
    int N;
    cin >> N;
    vector<vector<int>> lst(N, vector<int>(2));

    for (int i = 0; i < N; i++) {
        cin >> lst[i][0] >> lst[i][1];
    }

    int L, P;
    cin >> L >> P;

    vector<vector<int>> sortedLst(N + 1, vector<int>(2));
    for (int i = 0; i < N; i++) {
        sortedLst[i][0] = L - lst[i][0];
        sortedLst[i][1] = lst[i][1];
    }
    sortedLst[N][0] = L;
    sortedLst[N][1] = 0;

    sort(sortedLst.begin(), sortedLst.end(), [](const vector<int>& a, const vector<int>& b) {
        return a[0] < b[0];
    });

    vector<vector<int>> dp(N + 1, vector<int>(P + 1, INT_MAX));

    if (sortedLst[0][0] > P) {
        cout << -1 << endl;
    } else {
        dp[0][P - sortedLst[0][0]] = 0;

        for (int i = 0; i < N; i++) {
            int a = sortedLst[i][0];
            int b = sortedLst[i][1];
            for (int j = 0; j <= P; j++) {
                if (dp[i][j] == INT_MAX) {
                    continue;
                }

                int diff = sortedLst[i + 1][0] - a;

                int nxt = min(j + b, P) - diff;
                if (nxt >= 0) {
                    dp[i + 1][nxt] = min(dp[i][j] + 1, dp[i + 1][nxt]);
                }

                nxt = j - diff;
                if (nxt >= 0) {
                    dp[i + 1][nxt] = min(dp[i][j], dp[i + 1][nxt]);
                }
            }
        }

        int ans = *min_element(dp[N].begin(), dp[N].end());
        cout << (ans != INT_MAX ? ans : -1) << endl;
    }

    return 0;
}
