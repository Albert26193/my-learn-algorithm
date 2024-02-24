// Created by Albert at 2023/11/12 23:08
// leetgo: 1.3.8
// https://leetcode.cn/problems/friends-of-appropriate-ages/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

// TODO: redo it with two pointers
func getEligibleCount(left int, right int, ages []int, cur int) int {
	lastIndex := right
	// [... x ... y cur]
	for left < right {
		mid := (left + right) / 2
		if cur >= ages[mid] && cur < 2*ages[mid]-14 {
			right = mid
		} else {
			left = mid + 1
		}
	}

	if cur >= ages[left] && cur < 2*ages[left]-14 {
		return lastIndex - left + 1
	}

	return 0
}

func numFriendRequests(ages []int) (ans int) {

	// n := len(ages)
	sort.Ints(ages)
	fmt.Println(ages)
	if len(ages) <= 1 {
		return 0
	}

	sameAgeCount := 0
	for i := 1; i < len(ages); i++ {
		if ages[i] <= 14 {
			continue
		}

		count := getEligibleCount(0, i-1, ages, ages[i])

		if ages[i] == ages[i-1] {
			sameAgeCount += 1
		} else {
			sameAgeCount = 0
		}

		ans += count
		ans += sameAgeCount
	}
	return
}

// @lc code=end

func main() {
	file, err := os.Open("./testcases.txt")
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	// skip first line
	ReadLine(reader)
	ages := Deserialize[[]int](ReadLine(reader))
	ans := numFriendRequests(ages)

	fmt.Println("\noutput:", Serialize(ans))
}
