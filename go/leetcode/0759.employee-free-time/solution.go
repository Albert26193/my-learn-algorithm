// Created by Albert at 2023/09/12 14:52
// leetgo: 1.3.7
// https://leetcode.cn/problems/employee-free-time/

package main

import (
	// "bufio"
	"fmt"
	// "os"
	"sort"
	// "sort"
	// . "github.com/j178/leetgo/testutils/go"
)

type Interval struct {
	Start int
	End   int
}

// @lc code=begin

func employeeFreeTime(schedule [][]*Interval) (ans []*Interval) {
	sch := make([]*Interval, 0)
	for i := 0; i < len(schedule); i++ {
		sch = append(sch, schedule[i]...)
	}

	sort.Slice(sch, func(i, j int) bool {
		if sch[i].Start == sch[j].Start {
			return sch[i].End < sch[j].End
		}

		return sch[i].Start < sch[j].Start
	})

	end := sch[0].End
	ans = make([]*Interval, 0)

	for i := 1; i < len(sch); i++ {
		if sch[i].Start <= end {
			if sch[i].End > end {
				end = sch[i].End
			}
		} else {
			ans = append(ans, &Interval{Start: end, End: sch[i].Start})
			end = sch[i].End
		}
	}
	return ans
}

// @lc code=end

// Warning: this is a manual question, the generated test code may be incorrect.
func main() {
	// stdin := bufio.NewReader(os.Stdin)
	// schedule := Deserialize[int](ReadLine(stdin))
	schedule := [][]*Interval{
		{{1, 2}, {5, 6}},
		{{1, 3}},
		{{4, 10}},
	}

	ans := employeeFreeTime(schedule)
	for _, interval := range ans {
		fmt.Println(interval.Start, interval.End)
	}
}
