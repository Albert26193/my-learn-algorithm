package main

import (
    "math"
    "fmt"
)

func findRightLegalMax(rightArr []int, current int) int {
    rightLegalMax := -1
    rightLegalMaxIndex := -1
    for  index, value := range rightArr {
        if value < current {
            rightLegalMax = int(math.Max(float64(rightLegalMax), float64(value)))
            if value == rightLegalMax {
                rightLegalMaxIndex = index
            }
        }
    }

    return rightLegalMaxIndex
}

func prevPermOpt1(arr []int) []int {
    rightLegalMaxIndex := -1
    for index, value := range arr {
        rightLegalMaxIndex = findRightLegalMax(arr[index:], value) + index 
        if rightLegalMaxIndex != -1 {
            arr[index], arr[rightLegalMaxIndex] = arr[rightLegalMaxIndex], arr[index]
            break;
        }
    }
    
    return arr
}

// TODO：低位高位看反了，重新做一下有空的时候
func main() {
  arr1 := []int{3, 2, 1}
  arr2 := []int {1, 1, 5}
  arr3 := []int {1, 9, 4, 6, 7}
  ans1 := prevPermOpt1(arr1)
  ans2 := prevPermOpt1(arr2)
  ans3 := prevPermOpt1(arr3)
  fmt.Println(ans1, ans2, ans3)
}
