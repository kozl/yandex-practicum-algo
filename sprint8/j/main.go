package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"unicode"
)

type node struct {
	values   []string
	children [26]*node
}

func (n *node) addValue(v string) {
	n.values = append(n.values, v)
}

func (n *node) subTrieValues() []string {
	result := []string{}
	result = append(result, n.values...)
	for i := 0; i < 26; i++ {
		if n.children[i] == nil {
			continue
		}
		result = append(result, n.children[i].subTrieValues()...)
	}
	return result
}

type camelCaseTrie struct {
	root *node
}

func (t camelCaseTrie) addWord(w string) {
	n := t.root
	for _, r := range w {
		if unicode.IsLower(r) {
			continue
		}
		if n.children[index(r)] == nil {
			n.children[index(r)] = &node{}
		}
		n = n.children[index(r)]
	}
	n.addValue(w)
}

func (t camelCaseTrie) find(pattern string) []string {
	n := t.root
	result := []string{}
	for _, r := range pattern {
		if n.children[index(r)] == nil {
			return result
		}
		n = n.children[index(r)]
	}
	result = append(result, n.subTrieValues()...)
	if len(result) == 0 {
		return []string{""}
	}
	sort.Strings(result)
	return result
}

func index(r rune) int {
	return int(unicode.ToLower(r)) - 'a'
}

func solve(lines, patterns []string) []string {
	trie := camelCaseTrie{
		root: &node{},
	}
	for _, line := range lines {
		trie.addWord(line)
	}

	result := []string{}
	for _, pattern := range patterns {
		result = append(result, trie.find(pattern)...)
	}
	return result
}

func main() {
	scanner := makeScanner()
	count := readInt(scanner)
	lines := make([]string, count)
	for i := 0; i < count; i++ {
		lines[i] = readLine(scanner)
	}

	count = readInt(scanner)
	patters := make([]string, count)
	for i := 0; i < count; i++ {
		patters[i] = readLine(scanner)
	}
	for _, line := range solve(lines, patters) {
		fmt.Println(line)
	}
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 50 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readLine(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}
