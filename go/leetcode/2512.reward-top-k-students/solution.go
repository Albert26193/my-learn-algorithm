// Created by Albert at 2023/10/11 11:14
// leetgo: 1.3.8
// https://leetcode.cn/problems/reward-top-k-students/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
type student struct {
	id    int
	score int
}

func topStudents(positive_feedback []string, negative_feedback []string, report []string, student_id []int, k int) (ans []int) {
	feedback_map := make(map[string]int)

	for _, s := range positive_feedback {
		feedback_map[s] = 3
	}

	for _, s := range negative_feedback {
		feedback_map[s] = -1
	}

	student_list := make([]student, 0)
	for i, r := range report {
		student_score := 0
		words := strings.Split(r, " ")
		for _, word := range words {
			score, exist := feedback_map[word]

			if exist {
				student_score += score
			}
		}

		current_student := student{id: student_id[i], score: student_score}
		student_list = append(student_list, current_student)
	}

	sort.Slice(student_list, func(i int, j int) bool {
		if student_list[i].score > student_list[j].score {
			return true
		} else if student_list[i].score == student_list[j].score {
			return student_list[i].id < student_list[j].id
		} else {
			return false
		}
	})

	for i := 0; i < k; i++ {
		ans = append(ans, student_list[i].id)
	}
	return
}

// @lc code=end

func main() {
	file, err := os.Open("./testcases.txt")
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	// skip first line
	ReadLine(reader)

	positive_feedback := Deserialize[[]string](ReadLine(reader))
	negative_feedback := Deserialize[[]string](ReadLine(reader))
	report := Deserialize[[]string](ReadLine(reader))
	student_id := Deserialize[[]int](ReadLine(reader))
	k := Deserialize[int](ReadLine(reader))
	ans := topStudents(positive_feedback, negative_feedback, report, student_id, k)

	fmt.Println("\noutput:", Serialize(ans))
}
