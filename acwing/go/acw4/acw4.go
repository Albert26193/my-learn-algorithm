package main

import (
	"bufio"
	"fmt"
	"os"
)

// TAGS: dp
// f[i][j] ==> from [0, ... , i-1], volume is j, max value is f[i][j]
func main() {
	file, err := os.Open("./testcases.txt")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	in := bufio.NewReader(file)
	defer file.Close()

	var N, V int
	fmt.Fscan(in, &N, &V)

	var v = make([]int, N)
	var w = make([]int, N)
	var s = make([]int, N)

	allCount := 0
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &v[i], &w[i], &s[i])
		allCount += s[i]
	}

	f := make([]int, V+1)

	for i := 0; i < N; i++ {
		for j := 0; j < s[i]; j++ {
			currentVolume := v[i]
			for k := V; k >= currentVolume; k-- {
				f[k] = maxx(f[k], f[k-currentVolume]+w[i])
			}
		}
	}

	// for i := 0; i < N; i++ {
	// 	for j := V; j >= v[i]; j-- {
	// 		for k := 1; k <= s[i] && k*v[i] <= j; k++ {
	// 			f[j] = maxx(f[j], f[j-k*v[i]]+k*w[i])
	// 		}
	// 	}
	// }

	// fmt.Println(v, w, s)
	fmt.Println(f[V])

}

func maxx(nums ...int) int {
	res := nums[0]
	for _, v := range nums {
		if v > res {
			res = v
		}
	}

	return res
}
