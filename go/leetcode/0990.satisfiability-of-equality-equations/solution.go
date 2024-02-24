// Created by Albert's server at 2024/01/21 00:29
// leetgo: dev
// https://leetcode.cn/problems/satisfiability-of-equality-equations/

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
}

func (uf *unionFind) find(x int) int {
	if uf.fa[x] != x {
		uf.fa[x] = uf.find(uf.fa[x])
	}

	return uf.fa[x]
}

func (uf *unionFind) merge(x int, y int) {
	fx, fy := uf.find(x), uf.find(y)
	if fx != fy {
		uf.fa[fx] = fy
	}
}

func initUF() *unionFind {
	fa := make([]int, 26)
	for i := 0; i < 26; i++ {
		fa[i] = i
	}

	return &unionFind{
		fa: fa,
	}
}

func equationsPossible(equations []string) bool {
	uf := initUF()

	for _, eq := range equations {
		op := eq[1]
		n1, n2 := int(eq[0]-'a'), int(eq[3]-'a')

		if op == '=' {
			uf.merge(n1, n2)
		}
	}

	// fmt.Println(uf.fa)
	for _, eq := range equations {
		op := eq[1]
		n1, n2 := int(eq[0]-'a'), int(eq[3]-'a')

		if op == '!' {
			if uf.find(n1) == uf.find(n2) {
				return false
			}
		}
	}

	return true
}

// @lc code=end

func main() {
	file, err := os.Open("./testcases.txt")
	if err != nil {
		fmt.Println("error :", err)
		return
	}
	in := bufio.NewReader(file)
	defer file.Close()

	ReadLine(in)
	equations := Deserialize[[]string](ReadLine(in))
	ans := equationsPossible(equations)

	fmt.Println("\noutput:", Serialize(ans))
}
