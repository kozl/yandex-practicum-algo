package main

/*
[Посылка](https://contest.yandex.ru/contest/22781/problems/B/)

# Принцип работы

Функция eval занимается разбором и вычислением выражения, результат вычисления пишит в переданный Writer.
Разбор выражения происходит следующим образом:
1. Разбиваем выражения на токены (разделены пробелом)
2. Пробегаемся по массиву токенов
	2а. Если токен это операнд, то есть число, то помещаем его в стек
	2б. Если токен это операция, то извлекаем два операнда из стека (вверху находится второй операнд, за ним первый) и производим вычисление
		и кладём результат в стек
3. После того как массив токенов заканчивается наверху стека находится результат вычисления, извлекаем его и записываем во Writer

При вычислениях обращаем внимание, что при делении отрицательного числа в Go происходит математическое округление, а нам в этом случае нужно
вернуть наибольшее целое, меньшее результату от деления. Для этого используем функцию math.Floor

# Сложность

Временная O(n), пространственная O(1) — в стеке в каждый момент хранится не более двух элементов.
*/

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

type opFn func(int, int) int

var operations = map[string]opFn{
	"+": func(a, b int) int { return a + b },
	"-": func(a, b int) int { return a - b },
	"*": func(a, b int) int { return a * b },
	"/": func(a, b int) int { return int(math.Floor(float64(a) / float64(b))) },
}

type stack struct {
	data []int
}

func newStack() *stack {
	return &stack{
		data: []int{},
	}
}

func (s *stack) push(i int) {
	s.data = append(s.data, i)
}

func (s *stack) pop() int {
	if len(s.data) == 0 {
		panic("empty stack")
	}
	lastIdx := len(s.data) - 1
	item := s.data[lastIdx]
	s.data = s.data[:lastIdx]
	return item
}

func eval(s *stack, expr string, w io.Writer) {
	tokens := strings.Split(expr, " ")
	for _, token := range tokens {
		if _, ok := operations[token]; ok {
			b, a := s.pop(), s.pop()
			result := operations[token](a, b)
			s.push(result)
			continue
		}
		number, _ := strconv.Atoi(token)
		s.push(number)
	}
	result := s.pop()
	fmt.Fprintln(w, result)
}

func main() {
	scanner := makeScanner()
	expr := readLine(scanner)
	s := newStack()
	eval(s, expr, os.Stdout)
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
