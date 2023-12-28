package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func solve(arr []int, s int) [][4]int {
	sort.Ints(arr)
	quadruples := map[[4]int]bool{}
	sumPairs := map[int][][2]int{}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			sum := arr[i] + arr[j]
			sumPairs[sum] = append(sumPairs[sum], [2]int{i, j})
		}
	}

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			sum := s - arr[i] - arr[j]
			if _, ok := sumPairs[sum]; ok {
				for _, indexes := range sumPairs[sum] {
					if i != indexes[0] && j != indexes[0] && i != indexes[1] && j != indexes[1] {
						res := [4]int{arr[i], arr[j], arr[indexes[0]], arr[indexes[1]]}
						sort.Ints(res[:])
						quadruples[res] = true
					}
				}
			}
		}
	}
	result := [][4]int{}
	for k := range quadruples {
		result = append(result, k)
	}
	sort.Slice(result, func(i, j int) bool {
		if result[i][0] != result[j][0] {
			return result[i][0] < result[j][0]
		}
		if result[i][1] != result[j][1] {
			return result[i][1] < result[j][1]
		}
		if result[i][2] != result[j][2] {
			return result[i][2] < result[j][2]
		}
		if result[i][3] != result[j][3] {
			return result[i][3] < result[j][3]
		}
		return true
	})
	return result
}

func main() {
	scanner := makeScanner()
	_ = readInt(scanner)
	s := readInt(scanner)
	arr := readArray(scanner)
	result := solve(arr, s)
	fmt.Println(len(result))
	for _, ar := range result {
		printIntArray(ar[:])
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

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}

func printIntArray(array []int) {
	s := make([]string, len(array))
	for idx, i := range array {
		s[idx] = fmt.Sprint(i)
	}
	fmt.Println(strings.Join(s, " "))
}
