// Created by Albert at 2023/06/25 22:49
// leetgo: 1.3.2
// https://leetcode.cn/problems/circle-and-rectangle-overlapping/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func distanceTwoPoint(x1 int, y1 int, x2 int, y2 int) float64 {
	return math.Sqrt(float64((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)))
}

//TODO
func checkOverlap(radius int, xCenter int, yCenter int, x1 int, y1 int, x2 int, y2 int) bool {

}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	radius := Deserialize[int](ReadLine(stdin))
	xCenter := Deserialize[int](ReadLine(stdin))
	yCenter := Deserialize[int](ReadLine(stdin))
	x1 := Deserialize[int](ReadLine(stdin))
	y1 := Deserialize[int](ReadLine(stdin))
	x2 := Deserialize[int](ReadLine(stdin))
	y2 := Deserialize[int](ReadLine(stdin))
	ans := checkOverlap(radius, xCenter, yCenter, x1, y1, x2, y2)

	fmt.Println("\noutput:", Serialize(ans))
}
