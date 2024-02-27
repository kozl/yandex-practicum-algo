package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func solve(capacity int, piles []goldenPile) int {
	sort.Slice(piles, func(i, j int) bool { return piles[i].pricePerKg > piles[j].pricePerKg })
	pilesInBackpack := []goldenPile{}

	free := capacity
	for _, pile := range piles {
		if free == 0 {
			break
		}
		if pile.weightKg <= free {
			pilesInBackpack = append(pilesInBackpack, pile)
			free -= pile.weightKg
		} else {
			pilesInBackpack = append(pilesInBackpack, goldenPile{pricePerKg: pile.pricePerKg, weightKg: free})
			free = 0
		}
	}

	totalPrice := 0
	for _, pile := range pilesInBackpack {
		totalPrice += pile.weightKg * pile.pricePerKg
	}
	return totalPrice
}

type goldenPile struct {
	weightKg   int
	pricePerKg int
}

func main() {
	scanner := makeScanner()
	m := readInt(scanner)
	pilesCount := readInt(scanner)
	piles := make([]goldenPile, pilesCount)
	for i := 0; i < pilesCount; i++ {
		raw := readArray(scanner)
		piles = append(piles, goldenPile{pricePerKg: raw[0], weightKg: raw[1]})
	}
	fmt.Println(solve(m, piles))
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
