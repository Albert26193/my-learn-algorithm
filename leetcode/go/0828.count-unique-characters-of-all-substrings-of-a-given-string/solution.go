// Created by Albert at 2023/06/29 10:59
// leetgo: 1.3.2
// https://leetcode.cn/problems/count-unique-characters-of-all-substrings-of-a-given-string/

package main

import (
	"fmt"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func uniqueLetterString(s string) int {
	// recordLeft[s[i]]: the index of the first char the same as s[i](left)
	// recordRight[s[i]]: the index of the first char the same as s[i](right)
	// recordIndex[s[i]]: record the index to supply data for recrodLeft and recordRight

	sLength := len(s)
	recordLeft := make([]int, sLength)
	recordRight := make([]int, sLength)
	recordIndex := make([]int, 26)

	// record Left
	for i := range recordIndex {
		recordIndex[i] = -1
	}

	for i := 0; i < sLength; i += 1 {
		currentCharCode := int(s[i] - 'A')
		// from recordIndex[] get the first left index
		recordLeft[i] = recordIndex[currentCharCode]
		// fill current index to recordIndex[]
		recordIndex[currentCharCode] = i
	}

	// record Right
	for i := range recordIndex {
		recordIndex[i] = sLength
	}

	for i := sLength - 1; i >= 0; i -= 1 {
		currentCharCode := int(s[i] - 'A')
		recordRight[i] = recordIndex[currentCharCode]
		recordIndex[currentCharCode] = i
	}

	ans := 0
	for i := 0; i < sLength; i++ {
		// possible value of LEFT endpoint: (i - recordLeft[i])
		// possible value of RIGHT endpoint: (recordRight[i] - i)
		ans += (i - recordLeft[i]) * (recordRight[i] - i)
	}
	return ans
}

// @lc code=end

func main() {
	// stdin := bufio.NewReader(os.Stdin)
	// s := Deserialize[string](ReadLine(stdin))

	s := "ABC"
	ans := uniqueLetterString(s)

	fmt.Println("\noutput:", Serialize(ans))
}
