// Created by Albert at 2023/11/11 00:30
// leetgo: 1.3.8
// https://leetcode.cn/problems/repeated-dna-sequences/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

// [start, end]
func findRepeatedDnaSequences(s string) (ans []string) {
	SLength := len(s)
	if SLength < 10 {
		return ans
	}

	mp := make(map[string]int)
	currentString := s[0:10]
	mp[currentString] = 1

	for start, end := 1, 10; end < SLength; start, end = start+1, end+1 {
		currentString = currentString[1:10] + string(s[end])
		mp[currentString] += 1
		if mp[currentString] == 2 {
			ans = append(ans, currentString)
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
	ans := findRepeatedDnaSequences(s)

	fmt.Println("\noutput:", Serialize(ans))
}
