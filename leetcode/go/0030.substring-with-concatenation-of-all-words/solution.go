// Created by Albert at 2023/06/25 13:03
// https://leetcode.cn/problems/substring-with-concatenation-of-all-words/

package main

import (
	"fmt"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func checkWindowContent(subString string, singleWordLength int, wordCount int, dictionary map[string]int) bool {
	currentDictionary := make(map[string]int)
	left, right := 0, singleWordLength-1
	for i := 0; i < wordCount; i += 1 {
		currentWord := subString[left : right+1]
		currentDictionary[currentWord] += 1
		left += singleWordLength
		right += singleWordLength
	}

	for k := range dictionary {
		if currentDictionary[k] != dictionary[k] {
			return false
		}
	}

	return true
}

func findSubstring(s string, words []string) (ans []int) {
	dictionary := make(map[string]int)
	singleWordLength := len(words[0])
	WordCount := len(words)
	windowSize := singleWordLength * WordCount

	for _, word := range words {
		dictionary[word] += 1
	}

	startIndex, endIndex := 0, windowSize-1
	for endIndex < len(s) {
		subString := s[startIndex : endIndex+1]
		if checkWindowContent(subString, singleWordLength, WordCount, dictionary) {
			ans = append(ans, startIndex)
		}
		startIndex += 1
		endIndex += 1
	}

	return ans
}

// @lc code=end

func main() {
	// stdin := bufio.NewReader(os.Stdin)
	// s := Deserialize[string](ReadLine(stdin))
	s := "barfoothefoobarman"
	// s = "wordgoodgoodgoodbestword"
	// words := Deserialize[[]string](ReadLine(stdin))
	words := []string{"foo", "bar"}
	// words = []string{"word", "good", "best", "word"}

	ans := findSubstring(s, words)

	fmt.Println("\noutput:", Serialize(ans))
}
