function twoSum(numbers: number[], target: number): number[] {
  let leftPointer = 0;
  let numbersLength = numbers.length;
  let rightPointer = numbersLength - 1;
  
  while (leftPointer < rightPointer) {
    if (numbers[leftPointer] + numbers[rightPointer] == target) {
      return [leftPointer + 1, rightPointer + 1];
    } else if (numbers[leftPointer] + numbers[rightPointer] < target) {
      leftPointer += 1;
    } else if (numbers[leftPointer] + numbers[rightPointer] > target) {
      rightPointer -= 1;
    }
  }

  return [-1, -1];
};


let numbers = [2, 7, 11, 15];
let target = 9;

console.log(twoSum(numbers, target));
