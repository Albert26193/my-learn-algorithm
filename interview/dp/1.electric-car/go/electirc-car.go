package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("../textcase.txt")
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	// stdin := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(reader, &N)
	// fmt.Fscan(stdin, &N)

	cityDistance := make([]int, N+1)
	maxInput := make([]int, N+1)

	allDistance := 0
	allVolume := 0

	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &cityDistance[i])
		fmt.Fscan(reader, &maxInput[i])
	}

	fmt.Fscan(reader, &allDistance)
	fmt.Fscan(reader, &allVolume)

	fmt.Println(cityDistance)
	fmt.Println(maxInput)
	fmt.Println(allDistance, allVolume)

	return
}
