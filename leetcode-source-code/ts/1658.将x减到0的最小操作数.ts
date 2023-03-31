function minOperations(nums: number[], x: number): number {
  // sum = x + subArray
  // subArray = x - sum
  // subArray's length keep longest
  let numsSum = nums.reduce(
    (prev, curr) => prev + curr,
    0
  );
  let subTarget = numsSum - x;
  
  console.log(subTarget);
  if (subTarget < 0) {
    return -1;
  }

  let numsLength = nums.length;
  let subSum  = 0;
  let left = 0, right = 0;
  let maxSubLength = -1;
  for (right = 0; left <= right && right < numsLength; right++) {
    subSum += nums[right];
    while (subSum > subTarget) {
      subSum -= nums[left];
      left += 1;
    }
    if (subSum == subTarget) {
    //   console.log(left, right);
      maxSubLength = Math.max(maxSubLength, right - left + 1);  
    }
  }

  return (maxSubLength == -1 ? -1 : numsLength - maxSubLength);
}
//let ans1 = minoperations([1, 1, 4, 2, 3], 5)
//let ans2 = minoperations([3, 2, 20, 1, 1, 3], 10)
let testArr = [8828,9581,49,9818,9974,9869,9991,10000,10000,10000,9999,9993,9904,8819,1231,6309];
let x = 134365;
let ans3 = minOperations(testArr, x);
console.log(ans3);
