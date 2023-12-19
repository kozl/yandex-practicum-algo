package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyz"
	a        = 1000
	mod      = 123_987_123
)

func solve(words []string) [][]int {
	set := map[int][]int{}
	pows := precalculateAPows(a, mod)
	for idx, word := range words {
		h := myHash(word, mod, pows)
		set[h] = append(set[h], idx)
	}
	groups := [][]int{}
	for _, v := range set {
		groups = append(groups, v)
	}
	sort.Slice(groups, func(i, j int) bool {
		return groups[i][0] < groups[j][0]
	})
	return groups
}

func myHash(word string, mod int, pows map[rune]int) int {
	hash := 0
	for _, r := range word {
		hash += (int(r) * pows[r]) % mod
	}
	return hash
}

func precalculateAPows(a, mod int) map[rune]int {
	m := map[rune]int{}
	for _, l := range alphabet {
		pow := 1
		times := int(l) - int('a')
		for i := 0; i < times; i++ {
			pow = (pow * a) % mod
		}
		m[l] = pow
	}
	return m
}

func main() {
	scanner := makeScanner()
	_ = readInt(scanner)
	s := readString(scanner)
	groups := solve(strings.Split(s, " "))
	for _, group := range groups {
		printArray(group)
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

func printArray(array []int) {
	s := make([]string, len(array))
	for i := 0; i < len(array); i++ {
		s[i] = strconv.Itoa(array[i])
	}
	fmt.Println(strings.Join(s, ""))
}
