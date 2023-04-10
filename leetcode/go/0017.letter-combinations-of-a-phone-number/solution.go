// Created by Albert at 2023/04/06 18:31
// https://leetcode.cn/problems/letter-combinations-of-a-phone-number/

package main

import (
	// "bufio"
	"fmt"
	// "os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func backTrack(digits string, numberToString map[rune]string, curDepth int, ans *[]string, combiantion string) {
    if (curDepth == len(digits)) {
        *ans = append(*ans, combiantion)
        return
    } else {
        digit := digits[curDepth]
        letters := numberToString[rune(digit)]
        for i := 0; i < len(letters); i++ {
            combiantion += string(letters[i])
            backTrack(digits, numberToString, curDepth + 1, ans, combiantion)
            combiantion = combiantion[:len(combiantion) - 1]
        }
    }
    
}

func letterCombinations(digits string) (ans []string) {
    if (digits == "") {
        return []string{}
    }

    numberToString := map[rune] string {
        '2': "abc",
        '3': "def",
        '4': "ghi",
        '5': "jkl",
        '6': "mno",
        '7': "pqrs",
        '8': "tuv",
        '9': "wxyz",
    }
    
    backTrack(digits, numberToString, 0, &ans, "")
    return 
}

// @lc code=end

func main() {
    // stdin := bufio.NewReader(os.Stdin)
    // digits := Deserialize[string](ReadLine(stdin))
    digits := "234"
    digits = "4"
    ans := letterCombinations(digits)
    fmt.Println("output: " + Serialize(ans))
}
