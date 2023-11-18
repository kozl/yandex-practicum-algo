package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getLongestWord(line string) string {
	line = strings.TrimSpace(line)
	parts := strings.Split(line, " ")
	var longest string
	for _, word := range parts {
		if len(word) > len(longest) {
			longest = word
		}
	}
	return longest
}

func main() {
	scanner := makeScanner()
	readInt(scanner)
	line := readLine(scanner)
	longestWord := getLongestWord(line)
	fmt.Println(longestWord)
	fmt.Println(len(longestWord))
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
