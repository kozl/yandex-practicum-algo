package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
[Успешная посылка](https://contest.yandex.ru/contest/25597/run-report/109638151/)

Решение задачи для подстроки t, в которую был добавлен новый символ складывается
из трёх элементов:
 1. Подстрока t может быть сведена к подстроке s путём удаления символа из t,
    в этом случае ответом будет расстояние левенштейна для подстроки предыдущего размера плюс одна операция
 2. Подстрока t может быть сведена к подстроке s путём удаления символа из s — аналогично п. 1
 3. Подстрока t может быть сведена к подстроке s путем замены последнего символа одной из подстрок.
    В том случае, если символы равны, замена не требуется и ответ равен расстроянию левенштейна для строк без последних символов

Сложность по времени: O(len(s)*len(t))
Пространственная сложность: O(len(s)*len(t))
*/
func levenstein(s, t string) int {
	prevdp, dp := makeDPs(t)
	for i := 1; i <= len(s); i++ {
		dp[0] = i
		for j := 1; j <= len(t); j++ {
			dp[j] = min(prevdp[j]+1, dp[j-1]+1, prevdp[j-1]+m(get(s, i), get(t, j)))
		}
		dp, prevdp = prevdp, dp
	}
	return prevdp[len(t)]
}

func m(a, b string) int {
	if a == b {
		return 0
	}
	return 1
}

func get(s string, i int) string {
	return string(s[i-1])
}

// makeDPs инициализирует массив dp, в котором индексы i и j соответствуют
// решению задачи для 0..i и 0..j подстрок s и t соответственно.
// При инициализации массива задаём также базовые случаи: для пустой строки s или t
//
// Оптимизация: инициализируем только 0 и 1 строку двумерного массива
func makeDPs(t string) ([]int, []int) {
	i0, i1 := make([]int, len(t)+1), make([]int, len(t)+1)
	for idx := range i0 {
		i0[idx] = idx
	}
	return i0, i1
}

func main() {
	scanner := makeScanner()
	s := readString(scanner)
	t := readString(scanner)
	fmt.Println(levenstein(s, t))
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
