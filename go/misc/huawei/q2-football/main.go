package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type comb struct {
	index  int
	scores string
}

type cp struct {
	total int
	seq   int
	lost  []int
}

func main() {
	// file, err := os.Open("./testcases.txt")
	// if err != nil {
	// 	fmt.Println("Error", err)
	// 	return
	// }
	//
	// in := bufio.NewReader(file)
	// defer file.Close()

	in := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(in, &n, &m)

	persons := make([]comb, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &persons[i].scores)
		persons[i].index = i + 1
	}

	sort.Slice(persons, func(i, j int) bool {
		tag1, tag2 := getTags(persons[i].scores), getTags(persons[j].scores)
		// 如果整体更多
		if tag1.total > tag2.total {
			return true
		}

		// 如果连续的数量更多
		if tag1.seq > tag2.seq {
			return true
		}

		// 如果失误更晚
		for index := range tag1.lost {
			if tag1.lost[index] > tag2.lost[index] {
				return true
			}
		}
		// 比较编号
		return i < j
	})

	for _, p := range persons {
		fmt.Printf("%d ", p.index)
	}
}

func getTags(str string) (ans cp) {
	// ans[0]: 1 的个数 cnt
	// ans[1]: 连续 1 的个数 maxSeq
	// ans[2]: 最早的 0 的位置 first

	n := len(str)
	cnt, maxSeq := 0, 0
	lost := make([]int, 0)
	seq := 0

	for i := 0; i < n; i++ {
		if str[i] == '1' {
			cnt++
			seq++
			if seq > maxSeq {
				maxSeq = seq
			}
		} else if str[i] == '0' {
			// 记录丢球
			lost = append(lost, i)

			// 连续序列最大长度清零
			seq = 0
		}
	}

	ans.total = cnt
	ans.seq = maxSeq
	ans.lost = lost
	return
}
