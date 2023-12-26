// Created by Albert's server at 2023/12/24 15:10
// leetgo: dev
// https://leetcode.cn/problems/minimum-garden-perimeter-to-collect-enough-apples/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
/*
   0 * * *(n,0)
   * * * *(n,1)
   * * * * ...
   * * * *(n, n)
  (0,n) (1,n)...  ==> 4 * (2*n*n + 2*sum + 2n - n ) = 8(n^2) + 8*sum+4n
    sum = (0 + (n-1)) * (n) / 2 = 0.5 * (n-1)*n
    all = 8n^2 + 4(n)(n-1) + 4n = 12n^2

    0 + 12 + 48 + ... 12n^2=12*(0 + 1 + 4 + ...n^2)
    0 * * (2,0)
    * * * (2,1)
    * * * (2,2)
    (0,2) (1,2) all = 10 + 4 - 2 = 12 => 48

    length: 4 * 2n = 8n
*/

func minimumPerimeter(neededApples int64) (ans int64) {
	if neededApples <= 12 {
		return 8
	}

	scale := 100
	if neededApples <= 1e5 {
		scale = 200
	}
	if neededApples <= 1e8 {
		scale = 700
	}
	if neededApples <= 1e12 {
		scale = 2000
	}
	if neededApples <= 1e15 {
		scale = 80000
	}

	left, right := int64(0), int64(scale)
	preMap := make([]int64, scale+1)

	for i := 1; i <= scale; i++ {
		preMap[i] = preMap[i-1] + int64(12*i*i)
	}

	for left < right {
		mid := (left + right) / 2
		if preMap[mid] >= neededApples {
			right = mid
		} else {
			left = mid + 1
		}
	}

	if preMap[left] >= neededApples {
		return 8 * left
	}

	return left
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
	neededApples := Deserialize[int64](ReadLine(in))
	ans := minimumPerimeter(neededApples)

	fmt.Println("\noutput:", Serialize(ans))
}
