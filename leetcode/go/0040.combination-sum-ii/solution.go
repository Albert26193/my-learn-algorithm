// Created by Albert at 2023/04/19 01:55
// https://leetcode.cn/problems/combination-sum-ii/

package main

import (
	// "bufio"
	"fmt"
	"sort"

	// "os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
// TODO 需要写题解
// During a traversal, duplicate elements are not allowed on the same level of the tree.
func backTrack(candidates []int, target int, combination []int, ans *[][]int, startIndex int, currentSum int) {
	if startIndex > len(candidates) {
		return
	}

	if currentSum > target {
		return
	}

	if currentSum == target {
		*ans = append(*ans, append([]int(nil), combination...))
	}

	for i := startIndex; i < len(candidates); i++ {
		if i > startIndex && candidates[i] == candidates[i-1] {
			continue
		}
		combination = append(combination, candidates[i])
		currentSum += candidates[i]
		backTrack(candidates, target, combination, ans, i+1, currentSum)
		currentSum -= candidates[i]
		combination = combination[:len(combination)-1]
	}

}

func combinationSum2(candidates []int, target int) (ans [][]int) {
	sort.Ints(candidates)
	backTrack(candidates, target, []int{}, &ans, 0, 0)
	return
}

// @lc code=end

func main() {
	// stdin := bufio.NewReader(os.Stdin)
	// candidates := Deserialize[[]int](ReadLine(stdin))
	candidates := []int{2, 5, 2, 1, 2}
	target := 5
	ans := combinationSum2(candidates, target)
	fmt.Println("output: " + Serialize(ans))
}
