// Created by Albert at 2023/11/10 15:49
// leetgo: 1.3.8
// https://leetcode.cn/problems/maximum-product-of-word-lengths/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

// key: a-string vs b-string ==> have no common char ==> O(1)
func stringToBit(s string) int {
	num := 0
	for _, ch := range s {
		offset := int(ch - 'a')
		num = (num | 1<<offset)
	}

	return num
}

func maxx(nums ...int) int {
	res := nums[0]

	for _, num := range nums {
		if num > res {
			res = num
		}
	}

	return res
}

func maxProduct(words []string) (ans int) {
	wordsCount := len(words)
	nums := make([]int, wordsCount)

	for i, word := range words {
		nums[i] = stringToBit(word)
	}

	for i := 0; i < wordsCount; i++ {
		for j := i + 1; j < wordsCount; j++ {
			if nums[i]&nums[j] == 0 {
				ans = maxx(ans, len(words[i])*len(words[j]))
			}
		}
	}
	return
}

// @lc code=end

func main() {
	file, err := os.Open("./testcases.txt")
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	// skip first line
	ReadLine(reader)

	words := Deserialize[[]string](ReadLine(reader))
	ans := maxProduct(words)

	fmt.Println("\noutput:", Serialize(ans))
}
