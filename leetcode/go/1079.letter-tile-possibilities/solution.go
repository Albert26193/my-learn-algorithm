// Created by Albert at 2023/04/13 20:46
// https://leetcode.cn/problems/letter-tile-possibilities/

package main

import (
	"fmt"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
// root
// 0:   	A 		B 		C
// 1:  	  B   C		C
// TODO
func backTrack(tiles string, combination string, ansString *[]string, used *[]int, startIndex int) {
	if startIndex > len(tiles) {
		return
	}

	if len(combination) > 0 {
		*ansString = append(*ansString, combination)
	}

	for i := startIndex; i < len(tiles); i++ {
		combination = combination + string(tiles[i])
		backTrack(tiles, combination, ansString, used, i+1)
		combination = combination[:len(combination)-1]
	}
}

func numTilePossibilities(tiles string) (ans int) {
	combination := ""
	used := make([]int, 0)
	ansString := make([]string, 0)
	backTrack(tiles, combination, &ansString, &used, 0)
	fmt.Println(ansString)
	return len(ansString)
}

// @lc code=end

func main() {
	// stdin := bufio.NewReader(os.Stdin)
	// tiles := Deserialize[string](ReadLine(stdin))
	tiles := "ABC"
	ans := numTilePossibilities(tiles)
	fmt.Println("output: " + Serialize(ans))
}
