/*
 * @lc app=leetcode.cn id=875 lang=golang
 *
 * [875] 爱吃香蕉的珂珂
 */

package main

import "fmt"

//"fmt"
//"strings"

// @lc code=start
func checkLegality(piles []int, h int, speed int) bool {
	allTime := 0
	for _, pile := range piles {
		spareTime := 0
		if pile%speed != 0 {
			spareTime = 1
		}
		allTime += spareTime + (pile / speed)
	}

	return allTime <= h
}

func minEatingSpeed(piles []int, h int) int {
	left, right := 1, int(1e9)
	for left < right {
		mid := (left + right) / 2
		if checkLegality(piles, h, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return right
}

// @lc code=end

func main() {
	piles := []int{30, 11, 23, 4, 20}
	h := 6
	ans := minEatingSpeed(piles, h)
	fmt.Println(ans)
}
