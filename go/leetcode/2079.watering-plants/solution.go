// Created by Albert's server at 2024/05/08 22:20
// leetgo: dev
// https://leetcode.cn/problems/watering-plants/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func wateringPlants(plants []int, capacity int) (ans int) {
	n := len(plants)
	cur, cnt := capacity, 0
	for i := 0; i < n; i++ {
		if cur < plants[i] {
			cur = capacity - plants[i]
			cnt += i
			cnt += i + 1
		} else {
			cur -= plants[i]
			cnt++
		}
	}
	return cnt
}

// @lc code=end

func main() {
	file, err := os.Open("./testcases.txt")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	in := bufio.NewReader(file)
	defer file.Close()

	// stdin := bufio.NewReader(os.Stdin)
	ReadLine(in)
	plants := Deserialize[[]int](ReadLine(in))
	capacity := Deserialize[int](ReadLine(in))
	ans := wateringPlants(plants, capacity)

	fmt.Println("\noutput:", Serialize(ans))
}
