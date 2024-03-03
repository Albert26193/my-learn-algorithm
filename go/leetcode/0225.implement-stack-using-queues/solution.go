// Created by Albert's server at 2024/03/03 12:05
// leetgo: dev
// https://leetcode.cn/problems/implement-stack-using-queues/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

type MyStack struct {
	helper, master []int
}

func Constructor() MyStack {
	var helper, master []int
	return MyStack{
		helper: helper,
		master: master,
	}
}

func (m *MyStack) Push(x int) {
	m.master = append(m.master, x)

	// put helper content to master
	for len(m.helper) > 0 {
		m.master = append(m.master, m.helper[0])
		m.helper = m.helper[1:]
	}

	m.helper, m.master = m.master, m.helper
}

func (m *MyStack) Pop() (ans int) {
	v := m.helper[0]

	m.helper = m.helper[1:]
	return v
}

func (m *MyStack) Top() (ans int) {
	return m.helper[0]
}

func (m *MyStack) Empty() bool {
	return len(m.helper) == 0
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
	ops := Deserialize[[]string](ReadLine(in))
	params := MustSplitArray(ReadLine(in))
	output := make([]string, 0, len(ops))
	output = append(output, "null")

	obj := Constructor()

	for i := 1; i < len(ops); i++ {
		switch ops[i] {
		case "push":
			methodParams := MustSplitArray(params[i])
			x := Deserialize[int](methodParams[0])
			obj.Push(x)
			output = append(output, "null")
		case "pop":
			ans := Serialize(obj.Pop())
			output = append(output, ans)
		case "top":
			ans := Serialize(obj.Top())
			output = append(output, ans)
		case "empty":
			ans := Serialize(obj.Empty())
			output = append(output, ans)
		}
	}
	fmt.Println("\noutput:", JoinArray(output))
}
