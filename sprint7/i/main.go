package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type move int

const (
	up = iota
	right
)

func solve(n, m int, field [][]int) (collected int, moves []move) {
	dp := make([][]int, n)
	dp[0] = make([]int, m)
	dp[0][0] = field[0][0]
	for x := 0; x < n; x++ {
		for y := 0; y < m; y++ {
			if dp[x] == nil {
				dp[x] = make([]int, m)
			}
			dp[x][y] = max(get(dp, x-1, y, 0), get(dp, x, y-1, 0)) + field[x][y]
		}
	}

	collected = dp[n-1][m-1]

	x, y := n-1, m-1
	for x != 0 || y != 0 {
		if get(dp, x-1, y, -1) > get(dp, x, y-1, -1) {
			moves = append(moves, up)
			x--
		} else {
			moves = append(moves, right)
			y--
		}
	}

	slices.Reverse(moves)

	return collected, moves
}

func get(dp [][]int, x, y, def int) int {
	if x < 0 || x >= len(dp) || y < 0 || y >= len(dp[x]) {
		return def
	}

	return dp[x][y]
}

func main() {
	scanner := makeScanner()
	raw := readArray(scanner)
	n, m := raw[0], raw[1]
	field := make([][]int, n)
	for i := 0; i < n; i++ {
		field[n-i-1] = readLineOfInts(scanner)
	}

	collected, moves := solve(n, m, field)
	fmt.Println(collected)
	for _, move := range moves {
		printMove(move)
	}
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

func readLineOfInts(scanner *bufio.Scanner) []int {
	scanner.Scan()
	runes := []rune(scanner.Text())
	arr := make([]int, len(runes))
	for i := 0; i < len(runes); i++ {
		arr[i], _ = strconv.Atoi(string(runes[i]))
	}
	return arr
}

func printMove(m move) {
	switch m {
	case up:
		fmt.Print("U")
	case right:
		fmt.Print("R")
	}
}
