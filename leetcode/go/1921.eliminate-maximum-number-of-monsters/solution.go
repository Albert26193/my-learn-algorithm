// Created by Albert at 2023/09/03 10:37
// leetgo: 1.3.7
// https://leetcode.cn/problems/eliminate-maximum-number-of-monsters/

package main

import (
	// "bufio"
	"fmt"
	// "os"
	"sort"

	. "github.com/j178/leetgo/testutils/go"
)

//TODO: 如何实现向上取整? 为什么需要向上取整?
// example: dist: 5  speed: 3 那么, 到达的时间就是 2
// @lc code=begin

func eliminateMaximum(dist []int, speed []int) (ans int) {
	distLength := len(dist)
	speedLength := len(speed)

	if distLength != speedLength {
		return -1
	}

	timeToCity := make([]int, distLength)

	for i := 0; i < distLength; i++ {
		timeToCity[i] = (dist[i]-1)/speed[i] + 1
	}

	sort.Ints(timeToCity)

	for i := 0; i < distLength; i++ {
		if timeToCity[i] <= i {
			return i
		}
	}

	fmt.Println(timeToCity)
	return distLength
}

// @lc code=end

func main() {
	// stdin := bufio.NewReader(os.Stdin)
	// dist := Deserialize[[]int](ReadLine(stdin))
	// speed := Deserialize[[]int](ReadLine(stdin))

	dist := []int{1, 3, 4}
	speed := []int{1, 1, 1}
	ans := eliminateMaximum(dist, speed)

	fmt.Println("\noutput:", Serialize(ans))
}
