package main

/*

Алгоритмическая сложность решения:
Пространственная — O(n), где n количество ячеек на доске (в этой задаче фиксированное)
Вычислительная — O(n)

*/

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func calculateScore(board [4][4]int, k int) int {
	digits := [10]int{}

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			digits[board[i][j]]++
		}
	}

	score := 0
	for p := 1; p < 10; p++ {
		if digits[p] != 0 && digits[p] <= k*2 {
			score++
		}
	}
	return score
}

func main() {
	scanner := makeScanner()
	k := readInt(scanner)
	board := readBoard(scanner)
	score := calculateScore(board, k)
	fmt.Println(score)
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 5 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readBoard(scanner *bufio.Scanner) [4][4]int {
	board := [4][4]int{}
	for i := 0; i < 4; i++ {
		scanner.Scan()
		line := scanner.Text()
		for j := 0; j < 4; j++ {
			if j == '.' {
				board[i][j] = 0
				continue
			}
			res, _ := strconv.Atoi(string(line[j]))
			board[i][j] = res
		}

	}
	return board
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}
