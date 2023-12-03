package main

/*
[Посылка](https://contest.yandex.ru/contest/23815/run-report/101692596/)

# Принцип работы

Агорит поиска такой же, как было описано в теории к спринту:
1. Рекурсивно делим массив на две части относительно опорного, который выбирается как случайный элемент массива
2. Левая часть содержит элементы меньше опорного, правая — не меньше
3. Рекурсия выполняется до тех пор, пока левая и правая части не состоят из одного элемента
4. Склеиваем левые и правые части между собой

Отличие состоит в двух вещах: нам нужно использовать кастомный ключ сортировки (используем функцию lt, она определяет
правила сравнения, которые указаны в условии) и мы не должны использовать дополнительную память.

Чтобы не использовать дополнительную память, в функции partition, которая отвечает за возврат левой и правой части
я использую два указателя. Внутри функции мы итерируемся до тех пор, пока указатели не встретятся. На каждой итерации
сдвигаем указатели так, чтобы они указывали на крайние элементы, которые не соответствуют условию. Условие такое:
все элементы левее левого указателя меньше опорного элемента, правее правого — не меньше. Как только больше сдвигать
указатели не можем — меняем местами элементы и сдвигаем их. Из функции возвращаем два слайса: [0:left) и (right:len(array)-1]

# Сложность

По времени O(nlogN), по памяти O(1)
*/

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type participant struct {
	name    string
	solved  int
	penalty int
}

func lt(a, b participant) bool {
	if a.solved != b.solved {
		return a.solved < b.solved
	}
	if a.penalty != b.penalty {
		return a.penalty > b.penalty
	}
	return a.name > b.name
}

func partition(array []participant, pivot participant, lt func(a, b participant) bool) ([]participant, []participant) {
	left := 0
	right := len(array) - 1

	for left <= right {
		if lt(array[left], pivot) {
			left++
			continue
		}
		if !lt(array[right], pivot) {
			right--
			continue
		}
		array[left], array[right] = array[right], array[left]
		left++
		right--
	}
	return array[:left], array[right+1:]
}

func quicksort(array []participant, cmpFn func(a, b participant) bool) []participant {
	if len(array) < 2 {
		return array
	}
	pivot := array[rand.Intn(len(array))]
	left, right := partition(array, pivot, cmpFn)
	return append(quicksort(left, cmpFn), quicksort(right, cmpFn)...)
}

func main() {
	scanner := makeScanner()
	count := readInt(scanner)
	array := make([]participant, count)
	for i := 0; i < count; i++ {
		s := readString(scanner)
		parts := strings.Split(s, " ")
		solved, _ := strconv.Atoi(parts[1])
		penalty, _ := strconv.Atoi(parts[2])
		array[i] = participant{
			name:    parts[0],
			solved:  solved,
			penalty: penalty,
		}
	}
	array = quicksort(array, lt)
	for i := len(array) - 1; i >= 0; i-- {
		fmt.Println(array[i].name)
	}
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

func readString(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}
