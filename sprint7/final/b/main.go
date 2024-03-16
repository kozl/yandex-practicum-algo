package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
[Успешная посылка](https://contest.yandex.ru/contest/25597/run-report/109726273/)

Для оптимизации пространственной сложности с O(nm) до O(m) (где n — кол-во чисел в последовательности, m — половина суммы последовательности)
храним не весь двумерный массив dp, а только значение для предыдущей строки. Двумерный массив хранит ответы на подзадачи для i и j, где
i подпоследовательность чисел с номерами от 1 до i, а j — величина суммы какого либо подмножества этих чисел. В массиве хранится ответ на вопрос,
есть ли такое подмножество в подпоследовательности, которое составляет сумму j.

Решением задачи будет значение в ячейках [i][sum/2]. Базовый случай — для j = 0, при любом i сумма есть.
Рекуррентная формула: имеем два случая — последнее число i входит в это подмножество или нет. В первом случае сумма есть, если она есть для подпоследовательности i-1
Во втором — сумма есть, если она есть для i-1 и суммы j из которой вычли последнее число.

Сложность по времени: O(len(scores)*s/2)
Пространственная сложность: O(s/2)
*/
func solve(scores []int) bool {
	s := sum(scores)
	if s%2 == 1 {
		return false
	}

	dp, nextdp := makeDPs(s / 2)
	for i := 1; i <= len(scores); i++ {
		for j := 1; j <= s/2; j++ {
			nextdp[j] = get(dp, j) || get(dp, j-scores[i-1])
		}
		copy(dp, nextdp)
	}

	return dp[s/2]
}

func get(dp []bool, i int) bool {
	if i < 0 || i > len(dp) {
		return false
	}
	return dp[i]
}

func sum(arr []int) (sum int) {
	for _, i := range arr {
		sum += i
	}
	return sum
}

// makeDPs инициализирует массивы dp и nextdp, которые соответствуют двум строкам (текущей и предыдущей)
// в двумерном массиве с индексами i и j.
// Двумерный массив хранит ответы на вопрос для подзадач: есть ли в множестве чисел с номерами от 1 до i
// такое подмножество, сумма которого равна j. Здесь же производится инициализация базового случая: для j = 0,
// то есть для нулевой суммы ответ для любого i положительный.
func makeDPs(n int) ([]bool, []bool) {
	dp, nextdp := make([]bool, n+1), make([]bool, n+1)
	dp[0], nextdp[0] = true, true
	return dp, nextdp
}

func main() {
	scanner := makeScanner()
	_ = readInt(scanner)
	scores := readArray(scanner)
	printBool(solve(scores))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
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

func printBool(b bool) {
	if b {
		fmt.Println("True")
		return
	}
	fmt.Println("False")
}
