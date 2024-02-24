package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./testcases.txt")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	in := bufio.NewReader(file)
	defer file.Close()

	var N, M int
	fmt.Fscan(in, &N, &M)

	graph := make([][]int, N+1)
	indegree := make([]int, N+1)

	record := make(map[string]bool)
	for i := 0; i < M; i++ {
		var x, y int

		fmt.Fscan(in, &x, &y)

		if x == y {
			fmt.Println(-1)
			return
		}

		xString := strconv.Itoa(x)
		yString := strconv.Itoa(y)
		combinedString := xString + "#" + yString
		if _, ok := record[combinedString]; ok {
			continue
		}

		record[combinedString] = true
		graph[x] = append(graph[x], y)
		indegree[y] += 1
	}

	// BFS
	queue := make([]int, 0)
	for i := 1; i <= N; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	ans := make([]int, 0)
	vis := make([]int, N+1)
	for len(queue) > 0 {
		head := queue[0]
		if vis[head] != 0 {
			fmt.Print(-1)
			return
		}

		vis[head] = 1
		queue = queue[1:]
		ans = append(ans, head)

		for _, nextNode := range graph[head] {
			indegree[nextNode] -= 1
			if indegree[nextNode] == 0 {
				queue = append(queue, nextNode)
			}
		}
	}

	for _, el := range indegree {
		if el > 0 {
			fmt.Print(-1)
			return
		}
	}

	for i, el := range ans {
		if i > 0 {
			fmt.Printf(" %d", el)
		} else {
			fmt.Printf("%d", el)
		}
	}
}
