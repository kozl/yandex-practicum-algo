package main

import (
	"fmt"
	"strconv"
)

func evaluateFunction(a int, b int, c int, x int) int {
	return a*x*x + b*x + c
}

func main() {
	a := readInt()
	x := readInt()
	b := readInt()
	c := readInt()
	fmt.Println(evaluateFunction(a, b, c, x))
}

func readInt() int {
	var aString string
	fmt.Scan(&aString)
	a, _ := strconv.Atoi(aString)
	return a
}
