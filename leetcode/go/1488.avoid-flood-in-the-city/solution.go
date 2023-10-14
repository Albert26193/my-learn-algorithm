// Created by Albert at 2023/10/13 23:15
// leetgo: 1.3.8
// https://leetcode.cn/problems/avoid-flood-in-the-city/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
// TODO: redo it
func avoidFlood(rains []int) []int {
	sunnyDays := make([]int, 0)
	recordRainyDays := make(map[int]int)
	ans := make([]int, len(rains))

	for i := 0; i < len(ans); i++ {
		ans[i] = 1
	}

	for i, lake := range rains {
		if lake == 0 {
			sunnyDays = append(sunnyDays, i)
		} else {
			ans[i] = -1

			if rainyDay, exist := recordRainyDays[lake]; exist {
				it := sort.SearchInts(sunnyDays, rainyDay)
				if it == len(sunnyDays) {
					return []int{}
				}

				ans[sunnyDays[it]] = lake
				copy(sunnyDays[it:len(sunnyDays)-1], sunnyDays[it+1:])
				sunnyDays = sunnyDays[:len(sunnyDays)-1]
			}
			recordRainyDays[lake] = i
		}

	}

	return ans
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
	rains := Deserialize[[]int](ReadLine(reader))
	ans := avoidFlood(rains)

	fmt.Println("\noutput:", Serialize(ans))
}
