// Created by Albert's server at 2024/01/16 22:49
// leetgo: dev
// https://leetcode.cn/problems/similar-string-groups/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
type unionFind struct {
	fa []int
	r  []int
}

func initUF(n int) *unionFind {
	fa := make([]int, n)
	r := make([]int, n)
	for i := 0; i < n; i++ {
		fa[i] = i
		r[i] = 1
	}

	return &unionFind{
		fa: fa,
		r:  r,
	}
}

func (uf *unionFind) find(x int) int {
	if uf.fa[x] != x {
		uf.fa[x] = uf.find(uf.fa[x])
	}
	return uf.fa[x]
}

func (uf *unionFind) merge(x int, y int) {
	fx, fy := uf.find(x), uf.find(y)
	if fx == fy {
		return
	}

	if uf.r[fx] <= uf.r[fy] {
		uf.fa[fx] = uf.fa[fy]
	} else {
		uf.fa[fy] = uf.fa[fx]
	}

	if uf.r[fx] == uf.r[fy] {
		uf.r[fy] += 1
	}
}

func (uf *unionFind) getCnt(n int) int {
	cnt := 0
	for i := 0; i < n; i++ {
		if uf.fa[i] == i {
			cnt += 1
		}
	}

	return cnt
}

func check(s1 string, s2 string) bool {
	cnt := 0
	for i := range s1 {
		if s1[i] != s2[i] {
			cnt += 1
		}
	}

	return cnt == 2 || cnt == 0
}

func numSimilarGroups(strs []string) (ans int) {
	n := len(strs)
	uf := initUF(n)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if check(strs[i], strs[j]) {
				uf.merge(i, j)
			}
		}
	}

	ans = uf.getCnt(n)

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
	strs := Deserialize[[]string](ReadLine(in))
	ans := numSimilarGroups(strs)

	fmt.Println("\noutput:", Serialize(ans))
}
