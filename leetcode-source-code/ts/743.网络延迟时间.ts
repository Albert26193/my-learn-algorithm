/*
 * @lc app=leetcode.cn id=743 lang=typescript
 *
 * [743] 网络延迟时间
 */

// @lc code=start
function networkDelayTime(times: number[][], n: number, k: number): number {
  if (n == 0 || n == 1) {
    return 0;
  }

  let N = 100 + 5, M = 6010;

  let weightMatrix: Number[][] = new Array(N).fill(null).map(i => new Array(N));

  weightMatrix[1][0] = 100;


  return 0;
};
// @lc code=end

