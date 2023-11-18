package main

import (
	"fmt"
	"strconv"
)

func checkParity(a int, b int, c int) bool {
	if (((a%2 == 1) || (a%2 == -1)) && ((b%2 == 1) || (b%2 == -1)) && ((c%2 == 1) || (c%2 == -1))) || ((a%2 == 0) && (b%2 == 0) && (c%2 == 0)) {
		return true
	}
	return false
}

func main() {
	a := readInt()
	b := readInt()
	c := readInt()
	if checkParity(a, b, c) {
		fmt.Println("WIN")
	} else {
		fmt.Println("FAIL")
	}
}

func readInt() int {
	var aString string
	fmt.Scan(&aString)
	a, _ := strconv.Atoi(aString)
	return a
}
