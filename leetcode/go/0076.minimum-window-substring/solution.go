// Created by Albert at 2023/06/25 16:10
// leetgo: 1.3.2
// https://leetcode.cn/problems/minimum-window-substring/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func minWindow(s string, t string) string {
	ms, mt := make(map[byte]int), make(map[byte]int)
	ns, nt := len(s), len(t)
	l, r := 0, 0
	cnt := 0

	res := ""
	for r = 0; r < ns; r++ {
		ms[s[r]]++

		if mt[s[r]] >= ms[s[r]] {
			cnt++
		}

		for l < r && ms[s[l]] > mt[s[l]] {
			ms[s[l]]--
			l++
		}

		if cnt == nt {
			
		}

	}
	return ""
}

// @lc code=end

func main() {
	//stdin := bufio.NewReader(os.Stdin)
	file, err := os.Open("./testcases.txt")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	in := bufio.NewReader(file)
	defer file.Close()

	ReadLine(in)
	s := Deserialize[string](ReadLine(in))
	t := Deserialize[string](ReadLine(in))

	ans := minWindow(s, t)

	fmt.Println("\noutput:", Serialize(ans))
}
