package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type area struct {
	start int
	end   int
}

func (a area) overlaps(another area) bool {
	if a.start <= another.start && a.end >= another.start {
		return true
	}
	if a.start > another.start && a.start <= another.end {
		return true
	}
	return false
}

func sortAreas(areas []area) {
	sort.Slice(areas, func(i, j int) bool {
		if areas[i].start != areas[j].start {
			return areas[i].start < areas[j].start
		}
		return areas[i].end < areas[j].end
	})
}

func mergeAreas(areas []area) []area {
	result := []area{areas[0]}
	for i := 1; i < len(areas); i++ {
		last := result[len(result)-1]
		if last.overlaps(areas[i]) {
			result[len(result)-1] = area{start: min(last.start, areas[i].start), end: max(last.end, areas[i].end)}
		} else {
			result = append(result, areas[i])
		}
	}
	return result
}

func main() {
	scanner := makeScanner()
	lines := readInt(scanner)
	areas := make([]area, lines)
	for i := 0; i < lines; i++ {
		array := readArray(scanner)
		areas[i] = area{array[0], array[1]}
	}
	sortAreas(areas)
	merged := mergeAreas(areas)
	for _, area := range merged {
		fmt.Printf("%d %d\n", area.start, area.end)
	}
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 20 * 1024 * 1024
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

func readArray(scanner *bufio.Scanner) []int {
	scanner.Scan()
	listString := strings.Split(scanner.Text(), " ")
	arr := make([]int, len(listString))
	for i := 0; i < len(listString); i++ {
		arr[i], _ = strconv.Atoi(listString[i])
	}
	return arr
}
