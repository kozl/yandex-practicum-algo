package main

import (
	"bufio"
	"math"
	"os"
	"sort"
	"strconv"
)

func factorize(number int) []int {
	result := []int{}
	if number == 1 {
		return result
	}
	delim := int(math.Round(math.Sqrt(float64(number))))
	for i := 2; i <= delim; i++ {
		if number%i == 0 {
			result = append(result, factorize(number/i)...)
			result = append(result, i)
			break
		}
	}
	if len(result) == 0 {
		result = append(result, number)
	}
	return result
}

func main() {
	scanner := makeScanner()
	number := readInt(scanner)
	factorization := factorize(number)
	sort.Sort(sort.IntSlice(factorization))
	printArray(factorization)
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

func printArray(arr []int) {
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(arr); i++ {
		writer.WriteString(strconv.Itoa(arr[i]))
		writer.WriteString(" ")
	}
	writer.Flush()
}
