package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func searchDay(array []int, number int, left int, right int, pos *int) {
	if left >= right {
		return
	}
	mid := (right + left) / 2
	if array[mid] >= number {
		*pos = mid
		searchDay(array, number, left, mid, pos)
	} else {
		searchDay(array, number, mid+1, right, pos)
	}
}

func result(array []int, price int) (int, int) {
	var bicycle, doubleBicycle = -1, -1
	searchDay(array, price, 0, len(array), &bicycle)
	searchDay(array, price*2, 0, len(array), &doubleBicycle)
	if bicycle != -1 {
		bicycle++
	}
	if doubleBicycle != -1 {
		doubleBicycle++
	}
	return bicycle, doubleBicycle
}

func main() {
	scanner := makeScanner()
	readInt(scanner)
	array := readArray(scanner)
	targetPrice := readInt(scanner)
	bicycle, doubleBicycle := result(array, targetPrice)
	fmt.Printf("%d %d", bicycle, doubleBicycle)
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 20 * 1024 * 1024
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
