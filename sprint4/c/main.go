package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func precalculatePows(a, mod, l int) []int {
	result := make([]int, l+1)
	result[0] = 1
	for i := 1; i <= l; i++ {
		result[i] = (result[i-1] * a) % mod
	}
	return result
}

func precalculatePrefixHashes(s string, a, mod int) []int {
	result := make([]int, len(s)+1)
	result[0] = 0
	for i := 1; i <= len(s); i++ {
		result[i] = (result[i-1]*a + int(s[i-1])) % mod
	}
	return result
}

func polyHashSubstring(prefixHashes, pows []int, start, end int, a, mod int) int {
	res := prefixHashes[end] - (prefixHashes[start-1] * pows[end-start+1])
	if res < 0 {
		res %= mod
		res += mod
	}
	return res
}

func main() {
	scanner := makeScanner()
	a := readInt(scanner)
	mod := readInt(scanner)
	s := readString(scanner)
	pows := precalculatePows(a, mod, len(s))
	prefixHashes := precalculatePrefixHashes(s, a, mod)
	inputs := readInt(scanner)
	for i := 0; i < inputs; i++ {
		args := readArray(scanner)
		res := polyHashSubstring(prefixHashes, pows, args[0], args[1], a, mod)
		fmt.Println(res)
	}
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

func readString(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}
