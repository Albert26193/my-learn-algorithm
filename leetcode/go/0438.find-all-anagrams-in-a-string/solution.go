// Created by Albert's server at 2024/01/03 17:11
// leetgo: dev
// https://leetcode.cn/problems/find-all-anagrams-in-a-string/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func findAnagrams(s string, p string) (ans []int) {
	ledger := make([]int, 26)

	for _, ch := range p {
		ledger[int(ch-'a')] += 1
	}

	rec := make([]int, 26)

	for left, right := 0, 0; right < len(s); right++ {
		rec[int(s[right]-'a')] += 1

		for left <= right && rec[int(s[right]-'a')] > ledger[int(s[right]-'a')] {
			rec[int(s[left]-'a')] -= 1
			left += 1
		}

		equal := true
		for i := 0; i < 26; i++ {
			if rec[i] != ledger[i] {
				equal = false
			}
		}

		if equal {
			ans = append(ans, left)
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

	ReadLine(in)
	s := Deserialize[string](ReadLine(in))
	p := Deserialize[string](ReadLine(in))
	ans := findAnagrams(s, p)

	fmt.Println("\noutput:", Serialize(ans))
}
