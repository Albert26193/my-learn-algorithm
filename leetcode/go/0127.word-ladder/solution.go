// Created by Albert's server at 2024/02/16 21:37
// leetgo: dev
// https://leetcode.cn/problems/word-ladder/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func ladderLength(beginWord string, endWord string, wordList []string) (ans int) {
	dict := make(map[string]bool)
	for _, v := range wordList {
		dict[v] = true
	}

	if !dict[endWord] {
		return 0
	}

	bq, eq := make([]string, 0), make([]string, 0)
	bq = append(bq, beginWord)
	eq = append(eq, endWord)

	visBegin, visEnd := make(map[string]bool), make(map[string]bool)
	visBegin[beginWord], visEnd[endWord] = true, true

	step := 0
	for len(bq) > 0 && len(eq) > 0 {
		step += 1

		if len(bq) > len(eq) {
			eq, bq = bq, eq
			visBegin, visEnd = visEnd, visBegin
		}

		temp := make([]string, 0)
		for _, word := range bq {
			for i := range word {
				for c := 'a'; c <= 'z'; c++ {
					nextWord := word[:i] + string(c) + word[i+1:]
					if visBegin[nextWord] {
						continue
					}

					if visEnd[nextWord] {
						return step + 1
					}

					if dict[nextWord] {
						temp = append(temp, nextWord)
						visBegin[nextWord] = true
					}
				}
			}
		}

		bq = temp
	}

	return 0
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

	ReadLine(in)
	beginWord := Deserialize[string](ReadLine(in))
	endWord := Deserialize[string](ReadLine(in))
	wordList := Deserialize[[]string](ReadLine(in))
	ans := ladderLength(beginWord, endWord, wordList)

	fmt.Println("\noutput:", Serialize(ans))
}
