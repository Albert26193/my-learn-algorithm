// Created by Albert at 2023/04/12 19:22
// https://leetcode.cn/problems/replace-the-substring-for-balanced-string/

package main

import (
	"fmt"
	"math"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
// inner string + external string = 4 * n
// because inner string can be replaced by any string with the same length
// the only one condition is:
// 		keep the external string: every element is less than n
func balancedString(s string) (ans int) {
	sLen := len(s)
	if sLen%4 != 0 {
		return -1
	}

	n := sLen / 4
	left, right := 0, 0
	cnt := map[rune]int{
		'Q': 0,
		'W': 0,
		'E': 0,
		'R': 0,
	}

	for i := range s {
		cnt[rune(s[i])] += 1
	}

	ans = 0x3f3f3f3f
	for right = 0; right < sLen; right++ {
		rightChar := rune(s[right])
		cnt[rightChar] -= 1
		for (left < sLen) && (cnt['Q'] <= n) && (cnt['W'] <= n) && (cnt['E'] <= n) && (cnt['R'] <= n) {
			leftChar := rune(s[left])
			cnt[leftChar] += 1
			ans = int(math.Min(float64(right-left+1), float64(ans)))
			left += 1
		}
	}
	return ans
}

// @lc code=end

func main() {
	// stdin := bufio.NewReader(os.Stdin)
	// s := Deserialize[string](ReadLine(stdin))
	s := "QQQW"
	s = "QWER"
	ans := balancedString(s)
	fmt.Println("output: " + Serialize(ans))
}
