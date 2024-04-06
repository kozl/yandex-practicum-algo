package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type node struct {
	value    *string
	children [26]*node
}

func (n *node) addValue(v string) {
	n.value = &v
}

type trie struct {
	root *node
}

func newTrie() trie {
	return trie{
		root: &node{},
	}
}

func (t trie) addWord(w string) {
	n := t.root
	for _, r := range w {
		if n.children[index(r)] == nil {
			n.children[index(r)] = &node{}
		}
		n = n.children[index(r)]
	}
	n.addValue(w)
}

func (t trie) traverseLine(line string) bool {
	n := t.root
	for _, r := range line {
		if n.value != nil {
			n = t.root
		}
		if n.children[index(r)] == nil {
			return false
		}
		n = n.children[index(r)]
	}
	return n.value != nil
}

func index(r rune) int {
	return int(r) - 'a'
}

func solve(line string, words []string) bool {
	t := newTrie()
	for _, word := range words {
		t.addWord(word)
	}
	return t.traverseLine(line)
}

func main() {
	scanner := makeScanner()
	line := readString(scanner)
	count := readInt(scanner)
	words := []string{}
	for i := 0; i < count; i++ {
		words = append(words, readString(scanner))
	}
	printBool(solve(line, words))
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

func printBool(b bool) {
	if b {
		fmt.Println("YES")
		return
	}
	fmt.Println("NO")
}
