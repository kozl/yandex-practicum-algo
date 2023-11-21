package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var numLetters = map[rune][]rune{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

func generate(s string, nums []rune, result *[]string) {
	if len(nums) == 0 {
		*result = append(*result, s)
		return
	}

	for _, letter := range numLetters[nums[0]] {
		generate(s+string(letter), nums[1:], result)
	}
}

func result(line string) []string {
	result := []string{}
	generate("", []rune(line), &result)
	return result
}

func main() {
	scanner := makeScanner()
	line := readString(scanner)
	sequences := result(line)
	fmt.Println(strings.Join(sequences, " "))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 20 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readString(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}
