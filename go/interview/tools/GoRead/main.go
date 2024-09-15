package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	file, err := os.Open("./testcases.txt")
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	
	in := bufio.NewReader(file)

	var n, m int

	fmt.Fscan(in, &n, &m)	
	defer file.Close()
}