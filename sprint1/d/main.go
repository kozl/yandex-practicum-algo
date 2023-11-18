package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getWeatherRandomness(temperatures []int) int {
	result := 0
	if len(temperatures) == 0 {
		return 0
	}
	if len(temperatures) == 1 {
		return 1
	}
	for i := range temperatures {
		if i == 0 {
			if temperatures[i] > temperatures[i+1] {
				result++
			}
			continue
		}
		if i == len(temperatures)-1 {
			if temperatures[i-1] < temperatures[i] {
				result++
			}
			continue
		}

		if temperatures[i-1] < temperatures[i] && temperatures[i] > temperatures[i+1] {
			result++
		}
	}
	return result
}

func main() {
	scanner := makeScanner()
	readInt(scanner)
	temperatures := readArray(scanner)
	fmt.Println(getWeatherRandomness(temperatures))
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
		writer.WriteString(strconv.Itoa(arr[i]))
		writer.WriteString(" ")
	}
	writer.Flush()
}
