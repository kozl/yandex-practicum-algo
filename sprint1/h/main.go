package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getSum(firstNumber string, secondNumber string) string {
	first := []int{}
	second := []int{}
	for i := len(firstNumber) - 1; i >= 0; i-- {
		num, _ := strconv.Atoi(string(firstNumber[i]))
		first = append(first, num)
	}
	for i := len(secondNumber) - 1; i >= 0; i-- {
		num, _ := strconv.Atoi(string(secondNumber[i]))
		second = append(second, num)
	}
	if len(first) > len(second) {
		for i := len(second); i < len(first); i++ {
			second = append(second, 0)
		}
	} else if len(first) < len(second) {
		for i := len(first); i < len(second); i++ {
			first = append(first, 0)
		}
	}

	overrun := 0
	result := []int{}
	for i := 0; i < len(first); i++ {
		sum := first[i] + second[i] + overrun
		if sum >= 2 {
			overrun = 1
			result = append(result, sum%2)
		} else {
			overrun = 0
			result = append(result, sum)
		}
	}
	if overrun == 1 {
		result = append(result, overrun)
	}

	resString := ""
	for i := len(result) - 1; i >= 0; i-- {
		resString += strconv.Itoa(result[i])
	}
	return resString
}

func main() {
	scanner := makeScanner()
	firstNumber := readLine(scanner)
	secondNumber := readLine(scanner)
	sum := getSum(firstNumber, secondNumber)
	fmt.Println(sum)
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readLine(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}
