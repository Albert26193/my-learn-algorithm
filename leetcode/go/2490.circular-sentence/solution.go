// Created by Albert at 2023/06/30 08:43
// leetgo: 1.3.2
// https://leetcode.cn/problems/circular-sentence/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func isCircularSentence(sentence string) bool {
	firstChar := sentence[0]
	lastChar := sentence[len(sentence)-1]
	if firstChar != lastChar {
		return false
	}

	for i := 0; i < len(sentence); i += 1 {
		if sentence[i] == ' ' {
			if i+1 < len(sentence) {
				firstChar = sentence[i+1]
			}

			if i-1 >= 0 {
				lastChar = sentence[i-1]
			}

			if firstChar != lastChar {
				return false
			}
		}
	}

	return true
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	sentence := Deserialize[string](ReadLine(stdin))
	ans := isCircularSentence(sentence)

	fmt.Println("\noutput:", Serialize(ans))
}
