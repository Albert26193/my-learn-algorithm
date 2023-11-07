// Created by Albert at 2023/11/07 16:36
// leetgo: 1.3.8
// https://leetcode.cn/problems/interleaving-string/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

/*
	f[i][j] ==> s1:i s2:j ==> merge into s3: i + j
	f[i][j] = f[i-1][j] && check(s3)
	f[i][j] = f[i][j-1] && check(s3)
*/
func isInterleave(s1 string, s2 string, s3 string) bool {
	s1Length := len(s1)
	s2Length := len(s2)
	if len(s3) != s1Length+s2Length {
		return false
	}

	f := make([][]bool, s1Length+1)
	for i := 0; i <= s1Length; i++ {
		f[i] = make([]bool, s2Length+1)
	}

	f[0][0] = true
	for i := 0; i <= s1Length; i++ {
		for j := 0; j <= s2Length; j++ {
			m := i + j
			if i > 0 {
				f[i][j] = f[i][j] || (f[i-1][j] && s1[i-1] == s3[m-1])
			}
			if j > 0 {
				f[i][j] = f[i][j] || (f[i][j-1] && s2[j-1] == s3[m-1])
			}
		}
	}

	return f[s1Length][s2Length]
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

	s1 := Deserialize[string](ReadLine(reader))
	s2 := Deserialize[string](ReadLine(reader))
	s3 := Deserialize[string](ReadLine(reader))
	ans := isInterleave(s1, s2, s3)

	fmt.Println("\noutput:", Serialize(ans))
}
