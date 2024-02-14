// Created by Albert's server at 2024/01/07 02:12
// leetgo: dev
// https://leetcode.cn/problems/decode-string/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func decodeString(s string) string {

}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := decodeString(s)

	fmt.Println("\noutput:", Serialize(ans))
}
