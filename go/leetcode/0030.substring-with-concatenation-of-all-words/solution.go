// Created by Albert at 2023/06/25 13:03
// https://leetcode.cn/problems/substring-with-concatenation-of-all-words/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func findSubstring(s string, words []string) (ans []int) {
	ls, m, n := len(s), len(words), len(words[0])
	for i := 0; i < n && i+m*n-1 < ls; i++ {
		differ := map[string]int{}

		// j : cur word index
		for j := 0; j < m; j++ {
			begin, end := i+j*n, i+(j+1)*n-1
			differ[s[begin:end+1]]++
		}

		for _, word := range words {
			differ[word]--
			if differ[word] == 0 {
				delete(differ, word)
			}
		}

		// begin slide word
		for start := i; start+m*n-1 < ls; start += n {
			if start != i {
				begin, end := start+(m-1)*n, start+m*n-1
				word := s[begin : end+1]
				differ[word]++
				if differ[word] == 0 {
					delete(differ, word)
				}
				word = s[start-n : start]
				differ[word]--
				if differ[word] == 0 {
					delete(differ, word)
				}
			}
			if len(differ) == 0 {
				ans = append(ans, start)
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

	in := bufio.NewReader(file)
	defer file.Close()

	// stdin := bufio.NewReader(os.Stdin)
	ReadLine(in)
	s := Deserialize[string](ReadLine(in))
	words := Deserialize[[]string](ReadLine(in))

	ans := findSubstring(s, words)

	fmt.Println("\noutput:", Serialize(ans))
}
