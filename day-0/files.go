package main

import (
	"bufio"
	"fmt"
	"os"
)

func tryMe(e error) {
	if e != nil {
		fmt.Println("OH SNAP!")
		panic(e)
	}
}

func files() {
	dat, err := os.ReadFile("./data.txt")
	tryMe(err)

	fmt.Print(string(dat))
	fmt.Println("\n---")

	f, err := os.Open("./data.txt")
	tryMe(err)
	coupleBytes := make([]byte, 10)
	num, err := f.Read(coupleBytes)
	tryMe(err)
	fmt.Println(num, string(coupleBytes))

	reader := bufio.NewReader(f)
	rBuf, err := reader.Peek(5)
	tryMe(err)
	fmt.Println("buff reader", string(rBuf))

	f.Close()
}
