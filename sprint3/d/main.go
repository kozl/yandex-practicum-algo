package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func solve(children, cookies []int) int {
	sort.Ints(children)
	sort.Ints(cookies)

	result := 0
	var i, j int
	for i < len(children) && j < len(cookies) {
		if cookies[j] >= children[i] {
			result++
			i++
			j++
		} else {
			j++
		}
	}

	return result
}

func main() {
	scanner := makeScanner()
	_ = readInt(scanner)
	chilren := readArray(scanner)
	_ = readInt(scanner)
	cookies := readArray(scanner)
	fmt.Println(solve(chilren, cookies))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 20 * 1024 * 1024
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

func readArray(scanner *bufio.Scanner) []int {
	scanner.Scan()
	listString := strings.Split(scanner.Text(), " ")
	arr := make([]int, len(listString))
	for i := 0; i < len(listString); i++ {
		arr[i], _ = strconv.Atoi(listString[i])
	}
	return arr
}
