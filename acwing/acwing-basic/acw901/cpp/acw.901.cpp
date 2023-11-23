#include <fstream>
#include <iostream>

using namespace std;

const int N = 3e2 + 5;
int m[N][N];
int dp[N][N];

int dx[4] = {1, 0, -1, 0};
int dy[4] = {0, 1, 0, -1};

int max_len = 0;
int r, c;

bool valid(int x, int y, int row, int col) {
    if (x < 0 || x >= row) {
        return false;
    }
    if (y < 0 || y >= col) {
        return false;
    }
    return true;
}

int dfs(int x, int y) {
    // 防止重复读
    if (dp[x][y] != 0) {
        return dp[x][y];
    }

    for (int i = 0; i < 4; ++i) {
        int nx = x + dx[i], ny = y + dy[i];
        if (!valid(nx, ny, r, c)) {
            continue;
        }
        if (m[nx][ny] >= m[x][y]) {
            continue;
        }
        dp[x][y] = max(dfs(nx, ny) + 1, dp[x][y]);
    }

    return dp[x][y];
}

int main() {
    ifstream fin("../textcase.txt");
    fin >> r >> c;

    cout << r << c << endl;

    for (int i = 0; i < r; i++) {
        for (int j = 0; j < c; j++) {
            int temp = 0;
            fin >> temp;
            cout << temp << endl;
            m[i][j] = temp;
        }
    }

    int max_len = 0;
    for (int i = 0; i < r; ++i) {
        for (int j = 0; j < c; ++j) {
            int len = dfs(i, j);
            max_len = max(max_len, len);
        }
    }

    printf("%d\n", max_len + 1);
    return 0;
}