// Created by Albert at 2023/09/03 17:14
// leetgo: 1.3.7
// https://leetcode.cn/problems/minimum-score-of-a-path-between-two-cities/

package main

import (
	// "bufio"
	"fmt"
	"math"
	// "os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
type node struct {
	to int
	d  int
}

func minNum(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func minScore(n int, roads [][]int) (ans int) {
	cityGraph := make([][]node, n)
	minDistance := math.MaxInt64

	for _, city := range roads {
		startCity, endCity, distance := city[0]-1, city[1]-1, city[2]
		cityGraph[startCity] = append(cityGraph[startCity], node{endCity, distance})
		cityGraph[endCity] = append(cityGraph[endCity], node{startCity, distance})
		minDistance = minNum(city[2], minDistance)
	}
	// fmt.Println(minDistance)

	vis := make([]bool, n)

	queue := make([]int, 0)

	queue = append(queue, 0)
	vis[0] = true

	ans = math.MaxInt64
	for len(queue) > 0 {
		currentCity := queue[0]
		queue = queue[1:]

		nextCitys := cityGraph[currentCity]

		for _, nextCity := range nextCitys {
			ans = minNum(ans, nextCity.d)
			// fmt.Println(ans)
			if ans == minDistance {
				return ans
			}

			if vis[nextCity.to] {
				continue
			}
			queue = append(queue, nextCity.to)
			vis[nextCity.to] = true
		}

	}
	return
}

// @lc code=end

func main() {
	// stdin := bufio.NewReader(os.Stdin)
	// n := Deserialize[int](ReadLine(stdin))
	// roads := Deserialize[[][]int](ReadLine(stdin))

	n := 4
	roads := [][]int{{1, 2, 9}, {2, 3, 6}, {2, 4, 5}, {1, 4, 7}}
	ans := minScore(n, roads)

	fmt.Println("\noutput:", Serialize(ans))
}
