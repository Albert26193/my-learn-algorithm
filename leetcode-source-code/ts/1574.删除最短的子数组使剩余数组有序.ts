function findStartIndex(base: number, targetArray: number[]): number {
  let left = 0, right = targetArray.length -1;
  while (left < right) {
    let mid = Math.floor((left + right + 1) / 2);
    if (targetArray[mid] < base) {
      left = mid;
    } else {
      right = mid - 1;
    }
  }

  if (targetArray[right] < base) {
    return right;
  } else {
    return -1;
  }
}

function findLengthOfShortestSubarray(arr: number[]): number {
  let prefixArray: number[] = [];
  let suffixArray: number[] = [];
  let arrLength = arr.length;
  for (let i = 0; i < arrLength; i++) {
    if (prefixArray.length == 0 || arr[i] >= prefixArray[prefixArray.length - 1]) {
      prefixArray.push(arr[i]);
    } else {
      break;
    }
  }
  
  let suffixStartIndex = 0;
  for (let i = arrLength - 1; i >= 0; i--) {
    if (suffixArray.length == 0 || arr[i] <= suffixArray[0]) {
      suffixArray.unshift(arr[i]);
      suffixStartIndex = i;
    } else {
      break;
    }
  }

  console.log(prefixArray, suffixArray);
  if (prefixArray.length + suffixArray.length > arrLength) {
    return 0;
  }
  let minLength = Math.min(arrLength - prefixArray.length, arrLength - suffixArray.length);
  prefixArray.forEach((el, index) => {
    let rightStartIndex = suffixStartIndex + findStartIndex(el, suffixArray);
    let subArrayLength = rightStartIndex - (index + 1) + 1;
    console.log(arr[index], arr[rightStartIndex], subArrayLength);
    minLength = Math.min(subArrayLength, minLength);
  })

  return minLength;
}

//         0  1 2 3  4 5 6 7  8 9 10
let arr = [16,10,0,3,22,1,14,7,1,12,15];
let minLen = findLengthOfShortestSubarray(arr);
console.log(minLen);
