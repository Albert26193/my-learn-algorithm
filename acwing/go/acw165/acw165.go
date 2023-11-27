package main

import (
	"bufio"
	"fmt"
	"os"
)

var N, W int

// TAGS: dp
func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &N)
	fmt.Fscan(in, &W)
}
