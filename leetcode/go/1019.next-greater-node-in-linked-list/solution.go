// https://leetcode.cn/problems/next-greater-node-in-linked-list/

package main

import (
	"fmt"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func nextLargerNodes(head *ListNode) (ans []int) {
	currentNode := head
	// sta[x][0]: index of head
	// sta[x][1]: val in head
	sta := [][]int{}
	currentIndex := -1

	for currentNode != nil {
		currentIndex += 1
		currentVal := currentNode.Val
		ans = append(ans, 0)

		for len(sta) > 0 && currentVal > sta[len(sta)-1][1] {
			ans[sta[len(sta)-1][0]] = currentVal
			sta = sta[:len(sta)-1]
		}

		sta = append(sta, []int{currentIndex, currentVal})
		currentNode = currentNode.Next
	}

	return ans
}

// @lc code=end

func main() {
	// stdin := bufio.NewReader(os.Stdin)
	stdin := "[2, 1, 5]"
	// stdin = "[2, 7, 4, 3, 5]"
	head := Deserialize[*ListNode](stdin)
	ans := nextLargerNodes(head)
	fmt.Println("output: " + Serialize(ans))
}
