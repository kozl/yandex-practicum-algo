package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func isPowerOfFour(number int) bool {
	powersOfFour := map[int]bool{}
	power := 0
	for i := 0; power <= 10000; i++ {
		power = int(math.Pow(4, float64(i)))
		powersOfFour[power] = true
	}
	_, ok := powersOfFour[number]
	return ok
}

func main() {
	scanner := makeScanner()
	number := readInt(scanner)
	if isPowerOfFour(number) {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}
