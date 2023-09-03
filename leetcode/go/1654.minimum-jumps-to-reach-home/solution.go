// Created by Albert at 2023/09/01 21:03
// leetgo: 1.3.7
// https://leetcode.cn/problems/minimum-jumps-to-reach-home/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
// BFS
func minimumJumps(forbidden []int, a int, b int, x int) int {

	checkValidation := func(operationToCheck int, lastOperation int, nextPosition int) bool {
		if (operationToCheck == -1) && (nextPosition == -1) {
			return false
		}

		for _, num := range forbidden {
			if nextPosition == num {
				return false
			}
		}

		if nextPosition < 0 {
			return false
		}

		return true
	}

	queue := make([]int, 0)
	queue = append(queue, 0)

	jumpCount := 0
	lastOperation := 0
	for len(queue) > 0 {
		currentPosition := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

        if currentPosition == x {
            return currentPosition
        }
		leftPosition := currentPosition - b
		leftOperation := -1

		rightPosition := currentPosition + a
		rightOperation := 1

		if checkValidation(leftOperation, lastOperation, leftPosition) {
			queue = append(queue, leftPosition)
			lastOperation = -1
			jumpCount += 1
		}

		if checkValidation(rightOperation, lastOperation, rightPosition) {
			queue = append(queue, rightPosition)
			lastOperation = 1
			jumpCount += 1
		}

	}

	return jumpCount
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	forbidden := Deserialize[[]int](ReadLine(stdin))
	a := Deserialize[int](ReadLine(stdin))
	b := Deserialize[int](ReadLine(stdin))
	x := Deserialize[int](ReadLine(stdin))
	ans := minimumJumps(forbidden, a, b, x)

	fmt.Println("\noutput:", Serialize(ans))
}
