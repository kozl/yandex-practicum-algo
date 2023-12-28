package main

/*

[Посылка](https://contest.yandex.ru/contest/24414/run-report/102922851/)

# Принцип работы

Для поиска по документам формируем индекс. Индекс это ассоциативный массив, отображающий слова на набор их вхождений в документы.
Каждое вхождение характеризуется номером документа и количеством вхождений этого слова в определённый документ.

Для того чтобы составить индекс нужно пробежаться по всем документам, в каждом документе по каждому слову и заполнить массив.
При заполнении обращаем внимание на то, что если слово встречается впервые — в массиве будет либо nil, либо массив из попаданий в другие документы.

При поиске запрос делим на набор уникальный слов, пробегаемся по кажому слову и кладём в общий массив все попадания.
Нам нужно сгруппировать попадания по документам, суммировав количество попаданий слова и отсортировать полученный массив документов и попаданий
по количеству попаданий.

*/

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func search(index map[string][]occurence, query string) []int {
	words := uniq(strings.Split(query, " "))
	var queryWordsOccurencies []occurence
	for _, word := range words {
		if occurences, ok := index[word]; ok {
			queryWordsOccurencies = append(queryWordsOccurencies, occurences...)
		}
	}
	scoreMap := map[int]int{}
	for _, o := range queryWordsOccurencies {
		scoreMap[o.docIdx] += o.count
	}

	scores := [][2]int{}
	for k, v := range scoreMap {
		scores = append(scores, [2]int{k, v})
	}

	bubbleSortNTimes(scores, func(i, j int) bool {
		if scores[i][1] != scores[j][1] {
			return scores[i][1] < scores[j][1]
		}
		return scores[i][0] > scores[j][0]
	}, 5)

	if len(scores) <= 5 {
		return docNumber(scores)
	}
	return docNumber(scores[:5])
}

func bubbleSortNTimes(scores [][2]int, lessFn func(i, j int) bool, n int) {
	if len(scores) < n {
		n = len(scores)
	}

	for i := 0; i < n; i++ {
		for idx := len(scores) - 1; idx >= 1+i; idx-- {
			if !lessFn(idx, idx-1) {
				scores[idx], scores[idx-1] = scores[idx-1], scores[idx]
			}
		}
	}
}

func uniq(arr []string) (result []string) {
	m := map[string]bool{}
	for _, s := range arr {
		if _, ok := m[s]; !ok {
			m[s] = true
			result = append(result, s)
		}
	}
	return result
}

func docNumber(scores [][2]int) []int {
	result := make([]int, len(scores))
	for idx, score := range scores {
		result[idx] = score[0] + 1 // нужен номер документа, а не индекс в массиве
	}
	return result
}

type occurence struct {
	docIdx int
	count  int
}

func buildIndex(docs []string) map[string][]occurence {
	index := map[string][]occurence{}
	for docIdx, doc := range docs {
		words := strings.Split(doc, " ")
		for _, word := range words {
			occurences, found := index[word]
			if !found || occurences[len(occurences)-1].docIdx != docIdx {
				index[word] = append(index[word], occurence{docIdx: docIdx, count: 1})
			} else {
				occurences[len(occurences)-1].count++
			}
		}
	}
	return index
}

func main() {
	scanner := makeScanner()
	numDocs := readInt(scanner)
	docs := make([]string, numDocs)
	for i := 0; i < numDocs; i++ {
		doc := readString(scanner)
		docs[i] = doc
	}
	index := buildIndex(docs)
	numQueries := readInt(scanner)
	for i := 0; i < numQueries; i++ {
		query := readString(scanner)
		result := search(index, query)
		printIntArray(result)
	}
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func printIntArray(array []int) {
	s := make([]string, len(array))
	for idx, i := range array {
		s[idx] = fmt.Sprint(i)
	}
	fmt.Println(strings.Join(s, " "))
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}

func readString(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}
