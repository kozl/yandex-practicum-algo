package main

/*
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
	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case "+", "-", "/", "*":
			operand2, operand1 := s.pop(), s.pop()
			result := calc(operand1, operand2, tokens[i])
			s.push(result)
		default:
			val, _ := strconv.Atoi(tokens[i])
			s.push(val)
		}
	}
	result := s.pop()
	fmt.Fprintln(w, result)
}

func calc(operand1, operand2 int, operation string) int {
	switch operation {
	case "+":
		return operand1 + operand2
	case "-":
		return operand1 - operand2
	case "*":
		return operand1 * operand2
	case "/":
		return int(math.Floor(float64(operand1) / float64(operand2)))
	default:
		panic(fmt.Sprintf("unknown operation: %s", operation))
	}
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
