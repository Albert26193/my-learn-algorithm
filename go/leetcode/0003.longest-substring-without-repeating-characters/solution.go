// Created by Albert at 2023/11/11 00:00
// leetgo: 1.3.8
// https://leetcode.cn/problems/longest-substring-without-repeating-characters/

package main

import (
	"bufio"

	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func lengthOfLongestSubstring(s string) (ans int) {
	sLength := len(s)
	mp := make(map[byte]int)

	fast, slow := 0, 0
	for fast = 0; fast < sLength; fast++ {
		currentChar := s[fast]
		mp[currentChar] += 1
		for mp[currentChar] > 1 && slow <= fast {
			mp[s[slow]] -= 1
			slow += 1
		}

		currentLength := fast - slow + 1
		if currentLength > ans {
			ans = currentLength
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

	s := Deserialize[string](ReadLine(reader))
	ans := lengthOfLongestSubstring(s)

	fmt.Println("\noutput:", Serialize(ans))
}
