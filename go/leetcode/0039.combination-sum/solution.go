// Created by Albert at 2023/04/06 14:36
// https://leetcode.cn/problems/combination-sum/

package main

import (
	"fmt"
	"sort"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
// root
// 0:   2			 3 6 7
// 1: 	2		 3 6 7   	 ...
// 2:	2 		3 6 7 ...
func backTrack(candidates []int, target int, currentSum int, startIndex int, combination []int, ans *[][]int) {
	if currentSum == target {
		copyCombination := make([]int, len(combination))
		copy(copyCombination, combination)
		*ans = append(*ans, copyCombination)
		return
	}

	for i := startIndex; i < len(candidates); i++ {
		newSum := currentSum + candidates[i]
		if newSum > target {
			break
		}
		combination = append(combination, candidates[i])
		backTrack(candidates, target, newSum, i, combination, ans)
		combination = combination[:len(combination)-1]
	}
}

func combinationSum(candidates []int, target int) (ans [][]int) {
	sort.Ints(candidates)
	currentSum := 0
	combination := make([]int, 0)
	backTrack(candidates, target, currentSum, 0, combination, &ans)
	return
}

// @lc code=end

func main() {
	// stdin := bufio.NewReader(os.Stdin)
	// candidates := Deserialize[[]int](ReadLine(stdin))
	candidates := []int{2, 3, 6, 7}
	// target := Deserialize[int](ReadLine(stdin))
	target := 7
	ans := combinationSum(candidates, target)
	fmt.Println("output: " + Serialize(ans))
}
