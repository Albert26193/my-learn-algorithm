package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// TAGS: dfs
func main() {
	file, err := os.Open("./testcases.txt")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	in := bufio.NewReader(file)
	defer file.Close()

	var N, W int
	fmt.Fscan(in, &N, &W)

	var c = make([]int, N)
	cab := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Fscan(in, &c[i])
	}

	// fmt.Println(c)

	ans := N
	// cnt: count of containers

	var dfs func(currentCat int, cnt int)
	dfs = func(currentCat, cnt int) {
		if cnt >= ans {
			return
		}

		if currentCat == N {
			ans = minn(ans, cnt)
			return
		}

		// put current cat into deployed container
		for i := 0; i < cnt; i++ {
			if cab[i]+c[currentCat] <= W {
				cab[i] += c[currentCat]
				dfs(currentCat+1, cnt)
				cab[i] -= c[currentCat]
			}
		}

		// put current cat into a new container
		cab[cnt] = c[currentCat]
		dfs(currentCat+1, cnt+1)
		cab[cnt] = 0
	}

	// first put weight one into container
	// why? huge one is hard to transport
	// first transport huge ones will squeeze search space
	sort.Slice(c, func(i, j int) bool {
		return c[i] > c[j]
	})

	// fmt.Println(c)
	dfs(0, 0)
	fmt.Println(ans)
}

func minn(nums ...int) int {
	res := nums[0]
	for _, v := range nums {
		if v < res {
			res = v
		}
	}

	return res
}
