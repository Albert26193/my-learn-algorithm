// Created by Albert's server at 2024/01/14 15:20
// leetgo: dev
// https://leetcode.cn/problems/validate-ip-address/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func getIPV4(ip string) string {
	var check = func(s string) bool {
		if len(s) >= 4 {
			return false
		}

		if len(s) >= 2 && s[0] == '0' {
			return false
		}

		num, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("Error:", err)
			return false
		}

		if num < 0 || num > 255 {
			return false
		}
		return true
	}

	res := "IPv4"
	ss := strings.Split(ip, ".")
	valid := true
	if len(ss) != 4 {
		valid = false
	}

	for _, s := range ss {
		if !check(s) {
			valid = false
			break
		}
	}

	if !valid {
		res = "Neither"
	}
	return res
}

func getIPV6(ip string) string {
	ss := strings.Split(ip, ":")
	valid := true
	res := "IPv6"

	if len(ss) != 8 {
		valid = false
	}

	mp := make(map[rune]bool)
	for i := 0; i < 9; i++ {
		ch := rune('0' + i)
		mp[ch] = true
	}
	for i := 0; i < 6; i++ {
		ch1 := rune('a' + i)
		ch2 := rune('A' + i)
		mp[ch1] = true
		mp[ch2] = true
	}

	var check = func(s string) bool {
		if len(s) == 0 || len(s) >= 5 {
			return false
		}

		for _, ch := range s {
			if !mp[ch] {
				return false
			}
		}

		return true
	}

	for _, s := range ss {
		if !check(s) {
			valid = false
			break
		}
	}

	if !valid {
		res = "Neither"
	}
	return res
}

func validIPAddress(queryIP string) string {
	ans := "Neither"
	for _, ch := range queryIP {
		if ch == '.' {
			ans = getIPV4(queryIP)
			return ans
		}

		if ch == ':' {
			ans = getIPV6(queryIP)
			return ans
		}
	}

	return ans
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
	queryIP := Deserialize[string](ReadLine(in))
	ans := validIPAddress(queryIP)

	fmt.Println("\noutput:", Serialize(ans))
}
