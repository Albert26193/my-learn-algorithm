// Created by Albert at 2023/09/03 17:14
// leetgo: 1.3.7
// https://leetcode.cn/problems/minimum-score-of-a-path-between-two-cities/

package main

import (
	// "bufio"
	"fmt"
	"math"

	// "os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
// TODO: 1. DFS    2. Union-Find      3. BFS

/* BFS */
/*
type node struct {
	to int
	d  int
}

func minNum(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func minScore(n int, roads [][]int) (ans int) {
	cityGraph := make([][]node, n)
	minDistance := math.MaxInt64

	for _, city := range roads {
		startCity, endCity, distance := city[0]-1, city[1]-1, city[2]
		cityGraph[startCity] = append(cityGraph[startCity], node{endCity, distance})
		cityGraph[endCity] = append(cityGraph[endCity], node{startCity, distance})
		minDistance = minNum(city[2], minDistance)
	}
	// fmt.Println(minDistance)

	vis := make([]bool, n)

	queue := make([]int, 0)

	queue = append(queue, 0)
	vis[0] = true

	ans = math.MaxInt64
	for len(queue) > 0 {
		currentCity := queue[0]
		queue = queue[1:]

		nextCitys := cityGraph[currentCity]

		for _, nextCity := range nextCitys {
			ans = minNum(ans, nextCity.d)
			// fmt.Println(ans)
			if ans == minDistance {
				return ans
			}

			if vis[nextCity.to] {
				continue
			}
			queue = append(queue, nextCity.to)
			vis[nextCity.to] = true
		}

	}
	return
}
*/

/*
type unionFind struct {
	parent []int
	size   []int
	count  int
}

func newUnionFind(n int) *unionFind {
	parent := make([]int, n)
	size := make([]int, n)

	for i := range parent {
		parent[i] = i
		size[i] = 1
	}

	count := n

	return &unionFind{
		parent: parent,
		size:   size,
		count:  count,
	}
}

func (uf *unionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}

	return uf.parent[x]
}

func (uf *unionFind) union(x int, y int) bool {
	xRoot := uf.find(x)
	yRoot := uf.find(y)

	if xRoot == yRoot {
		return false
	}

	if uf.size[xRoot] < uf.size[yRoot] {
		xRoot, yRoot = yRoot, xRoot
	}

	uf.parent[yRoot] = xRoot
	uf.size[xRoot] += uf.size[yRoot]
	uf.count--
	return true
}

func (uf *unionFind) connected(x int, y int) bool {
	x = uf.find(x)
	y = uf.find(y)
	return x == y
}

func minScore(n int, roads [][]int) (ans int) {
	uf := newUnionFind(n + 1)

	for _, road := range roads {
		uf.union(road[0], road[1])
	}

	ans = math.MaxInt64

	for _, road := range roads {
		if road[2] < ans && uf.connected(1, road[0]) && uf.connected(n, road[0]) {
			ans = road[2]
		}
	}
	return
}
*/

type node struct {
	to int
	d  int
}

func minNum(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func minScore(n int, roads [][]int) (ans int) {
	cityGraph := make([][]node, n)
	minDistance := math.MaxInt64

	for _, city := range roads {
		startCity, endCity, distance := city[0]-1, city[1]-1, city[2]
		cityGraph[startCity] = append(cityGraph[startCity], node{endCity, distance})
		cityGraph[endCity] = append(cityGraph[endCity], node{startCity, distance})
		minDistance = minNum(city[2], minDistance)
	}
	// fmt.Println(minDistance)

	vis := make([]bool, n)

	queue := make([]int, 0)

	queue = append(queue, 0)
	vis[0] = true

	ans = math.MaxInt64

	var dfs func(int)

	dfs = func(x int) {
		vis[x] = true

		for _, city := range cityGraph[x] {
			ans = minNum(ans, city.d)
			if !vis[city.to] {
				dfs(city.to)
			}
		}
	}

	dfs(0)
	return
}

// @lc code=end
func main() {
	// stdin := bufio.NewReader(os.Stdin)
	// n := Deserialize[int](ReadLine(stdin))
	// roads := Deserialize[[][]int](ReadLine(stdin))

	n := 4
	roads := [][]int{{1, 2, 9}, {2, 3, 6}, {2, 4, 5}, {1, 4, 7}}
	ans := minScore(n, roads)

	fmt.Println("\noutput:", Serialize(ans))
}
