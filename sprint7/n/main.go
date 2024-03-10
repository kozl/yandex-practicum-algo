package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type state struct {
	sum      int
	freeDays []int
}

func (s state) isEmpty() bool {
	return s.sum == 0 && s.freeDays == nil
}

func solve(days []int) state {
	prev := make([]state, len(days)+1)
	if days[0] > 500 {
		prev[1] = state{sum: days[0], freeDays: []int{}}
	} else {
		prev[0] = state{sum: days[0], freeDays: []int{}}
	}

	// d stands for day
	for d := 2; d <= len(days); d++ {
		curr := make([]state, len(days)+1)
		// c stands for coupon count
		for c := 0; c <= len(days); c++ {
			var buy state
			if days[d-1] > 500 {
				if !get(prev, c-1).isEmpty() {
					fd := make([]int, len(prev[c-1].freeDays))
					copy(fd, prev[c-1].freeDays)
					buy = state{
						sum:      prev[c-1].sum + days[d-1],
						freeDays: fd,
					}
				}
			} else {
				if !get(prev, c).isEmpty() {
					fd := make([]int, len(prev[c].freeDays))
					copy(fd, prev[c].freeDays)
					buy = state{
						sum:      prev[c].sum + days[d-1],
						freeDays: fd,
					}
				}
			}
			var useCoupon state
			if !get(prev, c+1).isEmpty() {
				useCoupon = state{
					sum:      prev[c+1].sum,
					freeDays: append(prev[c+1].freeDays, d),
				}
			}
			curr[c] = minState(buy, useCoupon)
		}
		prev = curr
	}
	return findMin(prev)
}

func makeDP(d, c int) [][]state {
	dp := make([][]state, d+1)
	for i := 0; i <= d; i++ {
		dp[i] = make([]state, c+1)
	}
	return dp
}
func get(arr []state, c int) state {
	if c < 0 || c >= len(arr) {
		return state{}
	}
	return arr[c]
}

func findMin(arr []state) state {
	var min state
	for _, s := range arr {
		if s.isEmpty() {
			continue
		}
		if min.isEmpty() {
			min = s
			continue
		}
		if min.sum > s.sum {
			min = s
		}
	}
	return min
}

func minState(a, b state) state {
	if a.isEmpty() {
		return b
	}
	if b.isEmpty() {
		return a
	}
	if a.sum < b.sum {
		return a
	}
	return b
}

func main() {
	scanner := makeScanner()
	n := readInt(scanner)
	var days []int
	for i := 0; i < n; i++ {
		days = append(days, readInt(scanner))
	}

	result := solve(days)
	fmt.Printf("%d %d\n", result.sum, len(result.freeDays))
	printArray(result.freeDays)
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}

func printArray(arr []int) {
	s := []string{}
	for _, i := range arr {
		s = append(s, fmt.Sprint(i))
	}
	fmt.Println(strings.Join(s, " "))
}
