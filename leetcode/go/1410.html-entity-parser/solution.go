// Created by Albert at 2023/11/23 14:13
// leetgo: 1.3.8
// https://leetcode.cn/problems/html-entity-parser/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func entityParser(text string) string {
	var htmlMap = map[string]string{
		"&quot;":  "\"",
		"&apos;":  "'",
		"&amp;":   "&",
		"&gt;":    ">",
		"&lt;":    "<",
		"&frasl;": "/",
	}

	ans := make([]string, 0)

	isEntity := false
	n := len(text)
	for i := 0; i < n; {
		isEntity = false
		if text[i] == '&' {
			for k, v := range htmlMap {
				if i+len(k) <= n && k == text[i:i+len(k)] {
					// fmt.Println(k, i, v, len(k))
					ans = append(ans, v)
					// fmt.Println(ans)
					i += len(k)
					isEntity = true
					break
				}
			}
		}

		if !isEntity {
			ans = append(ans, text[i:i+1])
			i += 1
		}
	}

	// fmt.Println(ans)
	return strings.Join(ans, "")
}

// @lc code=end

func main() {
	file, err := os.Open("./testcases.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	reader := bufio.NewReader(file)
	defer file.Close()

	ReadLine(reader)
	text := Deserialize[string](ReadLine(reader))
	ans := entityParser(text)

	fmt.Println("\noutput:", Serialize(ans))
}
