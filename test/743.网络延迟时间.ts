/*
 * @lc app=leetcode.cn id=743 lang=typescript
 *
 * [743] 网络延迟时间
 */

import test from "node:test";

// @lc code=start
function floyd(n: number, weightMatrix: number[][]) {
  for (let p = 1; p <= n; p += 1) {
    for (let i = 1; i <= n; i++) {
      for (let j = 1; j <= n; j++) {
        weightMatrix[i][j] = Math.min(weightMatrix[i][j], weightMatrix[i][p] + weightMatrix[p][j]);
      }
    }
  }

  // console.log(weightMatrix);
}

function networkDelayTime(times: number[][], n: number, k: number): number {
  if (n == 0 || n == 1) {
    return 0;
  }

  let N = 100 + 5, M = 6010;

  let weightMatrix: number[][] = new Array(N).fill(null).map(i => []);

  for (let i = 1; i <= n; i++) {
    for (let j = 1; j <= n; j++) {
      weightMatrix[i][j] = weightMatrix[j][i] = (i == j ? 0 : 0x3f3f3f3f);
    }
  }
  
  for (const time of times) {
    weightMatrix[time[0]][time[1]] = time[2];
  }

  floyd(n, weightMatrix);  
  let ans = 0;

  for (let i = 1; i <= n; i++) {
    ans = Math.max(ans, weightMatrix[k][i]);
  }

  return ans == 0x3f3f3f3f ? -1 : ans;
};
// @lc code=end

(() => {
  let times = [
    [2, 1, 1],
    [2, 3, 1],
    [3, 4, 1],
  ];
  let n = 4, k = 2;
  let res = networkDelayTime(times, n , k);
  console.log(res);
})();
