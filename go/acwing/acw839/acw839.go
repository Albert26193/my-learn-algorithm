package main

// TODO: Errors
import (
	"bufio"
	"fmt"
	"os"
)

var N = 100010

// heap
var h = make([]int, N)

// index to heap index, i2hi[k] = i
var i2hi = make([]int, N)

// heap index to index, hi2i[]
var hi2i = make([]int, N)

// heap size
var heapSize = 0

func hSwap(x int, y int) {
	h[x], h[y] = h[y], h[x]
	i2hi[x], i2hi[y] = i2hi[y], i2hi[x]
	hi2i[x], hi2i[y] = hi2i[y], hi2i[x]
}

func sink(x int) {
	minIndex := x
	left := 2*x + 1
	right := 2*x + 2

	if left < heapSize && h[left] < h[minIndex] {
		minIndex = left
	}

	if right < heapSize && h[right] < h[minIndex] {
		minIndex = right
	}

	if minIndex != x {
		hSwap(x, minIndex)

		sink(minIndex)
	}
}

func up(x int) {
	fa := (x - 1) / 2
	if fa > 0 && h[x] < h[fa] {
		hSwap(x, fa)
		up(fa)
	}
}

func main() {
	file, err := os.Open("./testcases.txt")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	in := bufio.NewReader(file)
	defer file.Close()

	var n int
	fmt.Fscan(in, &n)

	m := 0
	for i := 0; i < n; i++ {
		var op string
		fmt.Fscan(in, &op)

		switch op {
		case "I":
			var x int
			fmt.Fscan(in, &x)
			heapSize++
			m++

			// insert into tail
			h[heapSize-1] = x
			i2hi[m-1] = heapSize - 1
			hi2i[heapSize-1] = m - 1
			up(heapSize - 1)
		case "PM":
			fmt.Println(h[0])
		case "DM":
			var x int
			hSwap(x, heapSize-1)
			heapSize -= 1
			sink(0)
		case "D":
			var k int
			fmt.Fscan(in, &k)
			k -= 1
			index := i2hi[k]
			hSwap(index, heapSize-1)
			heapSize -= 1
			up(index)
			sink(index)
		case "C":
			var k, x int
			k -= 1
			fmt.Fscan(in, &k, &x)
			index := i2hi[k]
			h[index] = x
			up(index)
			sink(index)
		}
	}
}
