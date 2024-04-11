package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

/*
[Успешная посылка](https://contest.yandex.ru/contest/26133/run-report/111791581/)

Алгоритм распаковки:

Используем стек, в который можно сохранять символы. Итерируемся по строке, на каждый символ проверяем:
1. Если это не открывающая и не закрывающая квадратная скобка, добавляем символ в стек (в стек попадут в том числе и числа).
2. Открывающие скобки пропускаем (начало фрагмента, который выглядит как n[A] однозначно определяется наличием цифры вначале)
3. Если встречаем закрывающую скобку, то делаем следующее:
	3.1 Извлекаем из стека символы и кладём их в промежуточный массив body, до тех пор, пока не найдём число — это будет означать начало
		последовательности, берём число и сохраняем как число повторений.
	3.2 Переворачиваем массив body (строка будет перевёрнутая, т.к. извлекаем символы по одному из стека), и записываем его в стек, повторив нужное число раз
4. Процесс повторяется до окончания строки

Алгоритм поиска наибольшего общего префикса:

1. Итерируемся по символам первой строки
2. Итерируемся по всем оставшимся строкам и сравниваем текущий символ первой строки и символ другой строки с тем же индексом
(проверяя не вышли ли за длину другой строки). Если он не равен или вышли за пределы длины — возвращаем общий префикс.
В противном случае инкрементируем длину префикса.

Оценка сложности: распаковка: в худшем случае O(nm), n — длина строки, m — количество вложенных конструкций для распаковки
				  поиск наибольшего общего префикса: O(kl), k — количество строк для распаковки, l — минимальная длина строки

Пространственная сложность: распаковка O(nm), n — результирующая длина строки, m — количество вложенных последовательностей
                            поиск наибольшего общего префикса: O(1)
*/

type stack []rune

func (s *stack) empty() bool {
	if s == nil {
		return true
	}
	return len(*s) == 0
}

func (s *stack) push(r rune) {
	*s = append(*s, r)
}

func (s *stack) pushString(str string) {
	*s = append(*s, []rune(str)...)
}

func (s *stack) pop() (r rune) {
	if s.empty() || s == nil {
		return rune(0)
	}
	last := len(*s) - 1
	r = (*s)[last]
	*s = (*s)[:last]
	return r
}

func (s *stack) string() string {
	return string(*s)
}

const (
	beginClause = '['
	endClause   = ']'
)

func unpack(line string) string {
	var s stack

	for _, r := range line {
		switch r {
		case beginClause:
			continue
		case endClause:
			var (
				top  rune
				body []rune
			)
			for top = s.pop(); !unicode.IsDigit(top); top = s.pop() {
				body = append(body, top)
			}
			count := runeToInt(top)
			slices.Reverse(body)
			s.pushString(repeat(body, count))
		default:
			s.push(r)
		}
	}
	return s.string()
}

func longestCommonPrefix(a, b string) string {
	minlen := len(a)
	if len(b) < minlen {
		minlen = len(b)
	}
	var prefixlen int
	for i := 0; i < minlen; i++ {
		if a[i] != b[i] {
			return a[:prefixlen]
		}
		prefixlen++
	}
	return a[:prefixlen]
}

func solve(lines []string) string {
	a := unpack(lines[0])
	lcp := a
	for i := 1; i < len(lines); i++ {
		b := unpack(lines[i])
		lcp = longestCommonPrefix(lcp, b)
	}
	return lcp
}

func repeat(s []rune, c int) string {
	var sb strings.Builder
	for i := 0; i < c; i++ {
		sb.WriteString(string(s))

	}
	return sb.String()
}

func runeToInt(r rune) int {
	return int(r - '0')
}

func main() {
	scanner := makeScanner()
	count := readInt(scanner)
	lines := []string{}
	for i := 0; i < count; i++ {
		lines = append(lines, readString(scanner))
	}
	fmt.Println(solve(lines))
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
