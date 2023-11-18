package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func getSum(bigNumber []int, smallNumber int) []int {
	first := []int{}
	second := []int{}
	for i := len(bigNumber) - 1; i >= 0; i-- {
		first = append(first, bigNumber[i])
	}
	r := smallNumber
	for r != 0 {
		second = append(second, r%10)
		r = r / 10
	}

	if len(first) > len(second) {
		for i := len(second); i < len(first); i++ {
			second = append(second, 0)
		}
	}

	if len(second) > len(first) {
		for i := len(first); i < len(second); i++ {
			first = append(first, 0)
		}
	}

	overrun := 0
	result := []int{}
	for i := 0; i < len(first); i++ {
		sum := first[i] + second[i] + overrun
		if sum >= 10 {
			overrun = 1
			result = append(result, sum%10)
		} else {
			overrun = 0
			result = append(result, sum)
		}
	}
	if overrun == 1 {
		result = append(result, overrun)
	}

	for i := len(result)/2 - 1; i >= 0; i-- {
		opp := len(result) - 1 - i
		result[i], result[opp] = result[opp], result[i]
	}
	return result
}

func main() {
	scanner := makeScanner()
	_ = readInt(scanner)
	bigNumer := readArray(scanner)
	smallNumber := readInt(scanner)
	printArray(getSum(bigNumer, smallNumber))
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

func printArray(arr []int) {
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(arr); i++ {
		_, _ = writer.WriteString(strconv.Itoa(arr[i]))
		_, _ = writer.WriteString(" ")
	}
	writer.Flush()
}
