// Created by Albert at 2023/04/10 16:06
// https://leetcode.cn/problems/design-skiplist/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

type Skiplist struct {

}

func Constructor() Skiplist {

	return Skiplist{}
}

func (s *Skiplist) Search(target int) bool {

}

func (s *Skiplist) Add(num int)  {

}

func (s *Skiplist) Erase(num int) bool {

}

/**
 * Your Skiplist object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Search(target);
 * obj.Add(num);
 * param_3 := obj.Erase(num);
 */

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	ops := Deserialize[[]string](ReadLine(stdin))
	params := MustSplitArray(ReadLine(stdin))
	output := make([]string, 0, len(ops))
	output = append(output, "null")

	obj := Constructor()

	for i := 1; i < len(ops); i++ {
		switch ops[i] {
		case "search":
			methodParams := MustSplitArray(params[i])
			target := Deserialize[int](methodParams[0])
			ans := Serialize(obj.Search(target))
			output = append(output, ans)
		case "add":
			methodParams := MustSplitArray(params[i])
			num := Deserialize[int](methodParams[0])
			obj.Add(num)
			output = append(output, "null")
		case "erase":
			methodParams := MustSplitArray(params[i])
			num := Deserialize[int](methodParams[0])
			ans := Serialize(obj.Erase(num))
			output = append(output, ans)
		}
	}
	fmt.Println("output: " + JoinArray(output))
}
