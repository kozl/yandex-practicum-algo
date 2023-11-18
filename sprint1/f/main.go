package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func isPalindrome(line string) bool {
	i := 0
	j := len(line) - 1
	for (j >= 0) && (i < len(line)) {
		if !unicode.IsLetter(rune(line[i])) && !unicode.IsDigit(rune(line[i])) {
			i++
			continue
		}
		if !unicode.IsLetter(rune(line[j])) && !unicode.IsDigit(rune(line[j])) {
			j--
			continue
		}
		if unicode.ToLower(rune(line[i])) == unicode.ToLower(rune(line[j])) {
			i++
			j--
			continue
		}
		return false
	}
	return true
}

func main() {
	scanner := makeScanner()
	line := readLine(scanner)
	if isPalindrome(line) {
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
