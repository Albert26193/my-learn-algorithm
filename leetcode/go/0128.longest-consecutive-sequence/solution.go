// Created by Albert's server at 2024/01/03 13:19
// leetgo: dev
// https://leetcode.cn/problems/longest-consecutive-sequence/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

// union-find
// type unionFind struct {
// 	fa   map[int]int
// 	size map[int]int
// }
//
// func (uf *unionFind) find(x int) int {
// 	if uf.fa[x] != x {
// 		uf.fa[x] = uf.find(uf.fa[x])
// 	}
//
// 	return uf.fa[x]
// }
//
// func (uf *unionFind) merge(x int, y int) {
// 	fax := uf.find(x)
// 	fay := uf.find(y)
// 	if fax != fay {
// 		uf.fa[y] = fax
// 		uf.size[fax] += uf.size[fay]
// 	}
//
// }
//
// func longestConsecutive(nums []int) (ans int) {
// 	uf := &unionFind{
// 		fa:   make(map[int]int),
// 		size: make(map[int]int),
// 	}
//
// 	for _, num := range nums {
// 		uf.fa[num] = num
// 		uf.size[num] = 1
// 	}
//
// 	for _, num := range nums {
// 		if _, ok := uf.fa[num+1]; ok {
// 			uf.merge(num, num+1)
// 		}
// 	}
//
// 	ans = 0
// 	for _, num := range nums {
// 		if uf.size[num] > ans {
// 			ans = uf.size[num]
// 		}
// 	}
// 	return
// }
//

// hash map

func longestConsecutive(nums []int) (ans int) {

	// key: index; value: range length
	mp := make(map[int]bool)

	// for every left endpoint
	for _, num := range nums {
		mp[num] = true
	}

	for num := range mp {
		if mp[num-1] {
			continue
		}

		currentLength := 1
		for mp[num+1] {
			num += 1
			currentLength += 1
		}

		ans = maxx(ans, currentLength)
	}

	return
}

func maxx(nums ...int) int {
	res := nums[0]
	for _, v := range nums {
		if v > res {
			res = v
		}
	}

	return res
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
	nums := Deserialize[[]int](ReadLine(in))
	ans := longestConsecutive(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
