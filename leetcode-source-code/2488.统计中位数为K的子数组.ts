function countSubarrays(nums: number[], k: number): number {
  if (!nums.includes(k)) {
    return -1
  };

  let numsLength: number = nums.length;
  let prefixSum: number[] = Array(numsLength + 1).fill(0);
  
  nums.reduce((acc, curr, i) => {
    acc += (curr > k ? 1 : (curr < k ? -1 : 0));
    prefixSum[i + 1] = acc;
    return acc;
  }, 0)
  
  let recordCounts = new Map();
  let kIndex = nums.indexOf(k);
  let ans = 0;
  //recordCounts.set(0, 1);
  
  //prefixSum.length == nums.length + 1
  for (let i = 0; i < kIndex + 1; i++) {
    if (!recordCounts.get(prefixSum[i])) {
      recordCounts.set(prefixSum[i], 1);
    } else {
      recordCounts.set(prefixSum[i], recordCounts.get(prefixSum[i]) + 1);
    }
  }

  // console.log(recordCounts);
  // console.log(prefixSum);

  for (let i = kIndex + 1; i < numsLength + 1; i++) {
    ans += (recordCounts.get(prefixSum[i]) || 0);
    ans += (recordCounts.get(prefixSum[i] - 1) || 0);
  }

  return ans;
}

// 4 5 6 7 1 2 3
// *
// 0 1 2 3 2 1 0
console.log(countSubarrays([3, 2, 1, 4, 5], 4))
console.log(countSubarrays([4, 1, 3, 2], 1))