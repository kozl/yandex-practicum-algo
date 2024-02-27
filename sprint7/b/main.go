package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func solve(lessons []lesson) []lesson {
	sort.Slice(lessons, func(i, j int) bool {
		if lessons[i].end == lessons[j].end {
			return lessons[i].start < lessons[j].start
		}
		return lessons[i].end < lessons[j].end
	})

	result := []lesson{lessons[0]}
	for _, lesson := range lessons[1:] {
		if result[len(result)-1].end <= lesson.start {
			result = append(result, lesson)
		}
	}
	return result
}

type lesson struct {
	start int
	end   int
}

func newLesson(start, end string) lesson {
	return lesson{
		start: toMinutes(start),
		end:   toMinutes(end),
	}
}

func (l lesson) String() string {
	sHours, sMinutes := l.start/60, l.start%60
	start := fmt.Sprint(sHours)
	if sMinutes != 0 {
		start += fmt.Sprintf(".%d", sMinutes)
	}
	eHours, eMinutes := l.end/60, l.end%60
	end := fmt.Sprint(eHours)
	if eMinutes != 0 {
		end += fmt.Sprintf(".%d", eMinutes)
	}
	return fmt.Sprintf("%s %s", start, end)
}

func toMinutes(s string) int {
	parts := strings.Split(s, ".")
	if len(parts) == 1 {
		hours, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		return hours * 60
	}
	hours, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}
	minutes, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}

	return hours*60 + minutes
}

func main() {
	run(os.Stdin, os.Stdout)
}

func run(r io.Reader, w io.Writer) {
	scanner := makeScanner(r)
	lessonCount := readInt(scanner)

	var lessons []lesson
	for i := 0; i < lessonCount; i++ {
		raw := readStringArray(scanner)
		lessons = append(lessons, newLesson(raw[0], raw[1]))
	}
	schedule := solve(lessons)
	fmt.Fprintln(w, len(schedule))
	for _, lesson := range schedule {
		fmt.Fprintln(w, lesson)
	}
}

func makeScanner(r io.Reader) *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(r)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readStringArray(scanner *bufio.Scanner) []string {
	scanner.Scan()
	return strings.Split(scanner.Text(), " ")
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}
