package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
ПО ЭТОЙ ЗАДАЧЕ НУЖНА ПОМОЩЬ
Успешной посылки нет, некоторые тесты не проходят по TL: https://contest.yandex.ru/contest/26133/run-report/111504678/4
Как оптимизировать идей что-то нет :(

Алгоритм решения задачи следующий:
1. Формируем префиксное дерево из всех возможных слов, которые могут встретиться в тексте
2. Алгоритм рекурсивный, для рекурсии используется стек. В стек кладём позицию символа, с которого начнём
   очередную проверку того, что строка состоит из слов
3. В стек добавляем индекс в случае если в префиксном дереве находим узел, содержащий значение
4. Если доходим до конца строки, а последний узел префиксного дерева содержит значение — значит строка состоит из корректных слов
5. Итерируемся до момента, пока не опустеет стек. Если до этого момента не вышли из функции, значит так и не дошли до конца строки
6. Используется мемоизация, массив checked хранит стартовые префиксы строк, которые уже проверяли

*/

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
		ridx := index(r)
		if n.children[ridx] == nil {
			n.children[ridx] = &node{}
		}
		n = n.children[ridx]
	}
	n.addValue(w)
}

type stack []int

func (s *stack) empty() bool {
	if s == nil {
		return true
	}
	return len(*s) == 0
}

func (s *stack) push(i int) {
	*s = append(*s, i)
}

func (s *stack) pop() (i int) {
	if s.empty() || s == nil {
		panic("unexpected: empty or nil stack")
	}
	last := len(*s) - 1
	i = (*s)[last]
	*s = (*s)[:last]
	return i
}

const firstCharacter = 'a'

func index(r rune) int {
	return int(r) - firstCharacter
}

func containValidWords(t trie, line string) bool {
	s := stack{}
	checked := make([]bool, len(line))
	s.push(0)
	for !s.empty() {
		start := s.pop()
		if checked[start] {
			continue
		}
		node := t.root
		var idx int
		for idx = start; idx < len(line); idx++ {
			if node.value != nil {
				s.push(idx)
			}
			ridx := index(rune(line[idx]))
			if node.children[ridx] == nil {
				break
			}
			node = node.children[ridx]
		}
		if idx == len(line) {
			if node.value != nil {
				return true
			}
			checked[start] = true
		}
	}
	return false
}

func solve(line string, words []string) bool {
	t := newTrie()
	for _, word := range words {
		t.addWord(word)
	}
	return containValidWords(t, line)
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
