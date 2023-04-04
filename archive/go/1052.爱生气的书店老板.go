/*
 * @lc app=leetcode.cn id=1052 lang=golang
 *
 * [1052] 爱生气的书店老板
 */
package main

import (
	"fmt"
	// "string"
	"math"
)

// @lc code=start
func maxSatisfied(customers []int, grumpy []int, minutes int) int {
    sales := make([]int, len(grumpy))
    satisfied := 0
    maxSatisfied := 0
    for i := 0; i < len(grumpy); i++ {
        if grumpy[i] == 0 {
            satisfied += customers[i]
        } else {
            sales[i] = customers[i]
        }
    }
    for i := 0; i < minutes; i++ {
        maxSatisfied += sales[i]
    }
    tempSatisfied := maxSatisfied
    for i := minutes; i < len(grumpy); i++ {
        tempSatisfied = tempSatisfied - sales[i-minutes] + sales[i]
        maxSatisfied = int(math.Max(float64(maxSatisfied), float64(tempSatisfied)))
    }
    return satisfied + maxSatisfied
}
// @lc code=end

func main() {
	customers := []int{1, 0, 1, 2, 1, 1, 7, 5}
	grumpy := []int{0, 1, 0, 1, 0, 1, 0, 1}
	minutes := 3
	ans := maxSatisfied(customers, grumpy, minutes)
	fmt.Println(ans)
}
