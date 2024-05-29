package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve(a, b []int) int {
	var maxlen int
	bmap := toMap(b)

	for apos, d := range a {
		if _, ok := bmap[d]; !ok {
			continue
		}
		for _, bpos := range bmap[d] {
			result := findCommonLen(a, b, apos, bpos)
			if result > maxlen {
				maxlen = result
			}
		}
	}
	return maxlen
}

func findCommonLen(a, b []int, apos, bpos int) int {
	var commonlen int
	for bpos < len(b) && apos < len(a) {
		if a[apos] != b[bpos] {
			return commonlen
		}
		commonlen++
		apos++
		bpos++
	}
	return commonlen
}

func main() {
	scanner := makeScanner()
	_ = readInt(scanner)
	a := readArray(scanner)
	_ = readInt(scanner)
	b := readArray(scanner)
	fmt.Println(solve(a, b))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
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

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}

func toMap(arr []int) map[int][]int {
	result := map[int][]int{}
	for idx, d := range arr {
		if _, ok := result[d]; !ok {
			result[d] = []int{}
		}
		result[d] = append(result[d], idx)
	}
	return result
}
