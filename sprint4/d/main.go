package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func uniq(array []string) []string {
	result := []string{}
	m := map[string]bool{}
	for _, val := range array {
		if !m[val] {
			result = append(result, val)
		}
		m[val] = true
	}
	return result

}

func main() {
	scanner := makeScanner()
	c := readInt(scanner)
	array := make([]string, c)
	for i := 0; i < c; i++ {
		array[i] = readString(scanner)
	}
	result := uniq(array)
	for _, val := range result {
		fmt.Println(val)
	}
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readString(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}
