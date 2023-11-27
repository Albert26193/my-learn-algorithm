package main

// TAGS: BFS
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func findIndex(s string) int {
	for i, ch := range s {
		if ch == 'x' {
			return i
		}
	}

	return -1
}

func checkLegality(index int, nextIndex int) bool {
	if index%3 == 0 && nextIndex%3 == 2 {
		return false
	}

	if index%3 == 2 && nextIndex%3 == 0 {
		return false
	}

	return true
}

func main() {
	// file, err := os.Open("./testcases.txt")
	// if err != nil {
	// 	fmt.Println("Error", err)
	// 	return
	// }

	// in := bufio.NewReader(file)
	// defer file.Close()

	in := bufio.NewReader(os.Stdin)

	var d = [4]int{1, -1, 3, -3}
	var s = make([]string, 9)
	var target = "12345678x"

	for i := 0; i < 9; i++ {
		_, err := fmt.Fscan(in, &s[i])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}

	ss := strings.Join(s, "")
	// fmt.Println(ss)
	queue := make([]string, 0)
	vis := make(map[string]int)

	// BFS
	queue = append(queue, ss)
	vis[ss] = 1

	steps := 0
	for len(queue) > 0 {

		levelSize := len(queue)
		for j := 0; j < levelSize; j++ {
			// pop
			head := queue[0]
			queue = queue[1:]

			// check
			if head == target {
				fmt.Printf("%d\n", steps)
				return
			}

			// push
			xIndex := findIndex(head)
			for i := 0; i < 4; i++ {
				nextIndex := xIndex + d[i]
				if !(0 <= nextIndex && nextIndex <= 8) {
					continue
				}

				if !checkLegality(xIndex, nextIndex) {
					continue
				}

				runes := []rune(head)
				runes[nextIndex], runes[xIndex] = runes[xIndex], runes[nextIndex]
				nextState := string(runes)

				if vis[nextState] == 1 {
					continue
				}

				queue = append(queue, nextState)
				vis[nextState] = 1
			}
		}
		steps += 1
	}

	fmt.Printf("%d", -1)
}
