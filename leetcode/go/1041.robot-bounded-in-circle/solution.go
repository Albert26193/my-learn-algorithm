// Created by Albert at 2023/04/12 17:05
// https://leetcode.cn/problems/robot-bounded-in-circle/

package main

import (
	"fmt"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
/*
	true:
		1. not int start point
		2. direction is not north
	fit the all of the 2 conditions is TRUE
*/
func isRobotBounded(instructions string) bool {
	/*
		directionsMap := map[int] string {
			0: "NORTH",
			1: "EAST",
			2: "SOUTH",
			3: "WEST",
		}
	*/

	mileInFourDirections := make([]int, 4)
	currentDirection := 0
	for _, ch := range instructions {
		if ch == 'G' {
			mileInFourDirections[currentDirection] += 1
		} else if ch == 'R' {
			currentDirection += 1
			currentDirection = (currentDirection + 4) % 4
		} else if ch == 'L' {
			currentDirection -= 1
			currentDirection = (currentDirection + 4) % 4
		} else {
			return false
		}
	}

	xMileSum := mileInFourDirections[1] - mileInFourDirections[3]
	yMileSum := mileInFourDirections[0] - mileInFourDirections[2]

	if (xMileSum == 0 && yMileSum == 0) || (currentDirection != 0) {
		return true
	}

	return false
}

// @lc code=end

func main() {
	// stdin := bufio.NewReader(os.Stdin)
	// stdin := "GGLLGG"
	// instructions := Deserialize[string](stdin)
	instructions := "GLRLLGLL"
	ans := isRobotBounded(instructions)
	fmt.Println("output: " + Serialize(ans))
}
