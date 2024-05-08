package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var str string
	var k int

	// read
	file, err := os.Open("testcases.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	in := bufio.NewReader(file)
	defer file.Close()

	// in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &str)
	fmt.Fscan(in, &k)
	n := len(str)
	fmt.Println(str, k)

	var nums []int

	for _, ch := range str {
		nums = append(nums, int(ch-'0'))
	}

	ans := getNum(nums, k)

	for len(ans) >= 1 && ans[0] == '0' {
		ans = ans[1:]
	}

	if len(ans) > n-k {
		ans = ans[:n-k]
	}

	if len(ans) == 0 {
		fmt.Println(0)
		return
	}
	fmt.Println(ans)
}

func getNum(nums []int, k int) string {
	sta := make([]int, 0)

	n := len(nums)
	for i := 0; i < n; i++ {
		cur := nums[i]

		for len(sta) > 0 && cur < sta[len(sta)-1] && k > 0 {
			sta = sta[:len(sta)-1]
			k--
		}

		sta = append(sta, nums[i])
	}

	ans := ""
	for _, el := range sta {
		ans = ans + strconv.Itoa(el)
	}

	return ans
}
