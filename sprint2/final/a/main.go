package main

/*
[Посылка](https://contest.yandex.ru/contest/22781/problems/)

# Принцип работы

Я реализовал очередь на базе кольцевого буфера, который представляет из себя слайс заданного capacity.
Capacity слайса, то есть размер буфера задаётся при инициализации структуры. Структура включает в себя указатели
на голову и хвост очереди, которые двигаются в зависимости от операции.

Кроме этого структура включает в себя количество уже содержащихся в очереди элементов для обработки ситуации переполнения очереди
или попытки извлечь что-то из пустой.

Указатель на голову очереди содержит индекс элемента в слайсе, который содержит первое значение очереди.
Указатель на хвост — индекс элемента в слайсе, который следует за последним значением.

При добавлении элементов следим за тем, чтобы head или tail не превысил размер слайса и берём остаток от деления на его размер.
При удалении — тоже следим чтобы индекс не вышел за пределы слайса слева, в этом случае при переходе через 0 мы должны изменить индекс,
сложив его с размером слайса.

# Доказательство корректности

Самое сложное, пожалуй, здесь убедиться в том, что мы не выйдем за пределы нижележащего массива.

При добавлении элемента в дек (в операциях push*) индексы нужно брать остатком от деления на длину массива,
при удалении (pop*) — убедиться что указатель будет переведён на конец массива.

Также, нужно обращать внимание на то, что head указывает именно на первый элемент, а tail на ближайший с конца свободный элемент.
Поэтому, например, при операциях добавления с головы очереди нам нужно сначала передвинуть указатель, а потом записать значение.
А при добавлении с хвоста — сначала записать, а потом передвинуть. Те же отличия будут при удалении.

# Сложность

Все операции с деком происходят за константное время, пространственная сложность O(n)
*/

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var (
	errFull  = errors.New("dequeue full")
	errEmpty = errors.New("dequeue empty")
)

type dequeue struct {
	data []int
	head int
	size int
	tail int
}

func newDeq(size int) *dequeue {
	return &dequeue{
		data: make([]int, size),
		head: 0,
		size: 0,
		tail: 0,
	}
}

func (d *dequeue) pushBack(i int) error {
	if d.size >= cap(d.data) {
		return errFull
	}
	d.data[d.tail] = i
	d.tail = toRight(d.tail, cap(d.data))
	d.size++
	return nil
}

func (d *dequeue) pushFront(i int) error {
	if d.size >= cap(d.data) {
		return errFull
	}
	d.head = toLeft(d.head, cap(d.data))
	d.data[d.head] = i
	d.size++
	return nil
}

func (d *dequeue) popFront() (int, error) {
	if d.size == 0 {
		return 0, errEmpty
	}
	elem := d.data[d.head]
	d.head = toRight(d.head, cap(d.data))
	d.size--
	return elem, nil
}

func (d *dequeue) popBack() (int, error) {
	if d.size == 0 {
		return 0, errEmpty
	}
	d.tail = toLeft(d.tail, cap(d.data))
	d.size--
	return d.data[d.tail], nil
}

func toLeft(pos, cap int) int {
	if pos == 0 {
		return cap - 1
	}
	return pos - 1
}

func toRight(pos, cap int) int {
	return (pos + 1) % cap
}

func eval(d *dequeue, cmd string, w io.Writer) {
	parts := strings.Split(cmd, " ")
	pushFuncs := map[string]func(int) error{
		"push_back":  d.pushBack,
		"push_front": d.pushFront,
	}

	popFuncs := map[string]func() (int, error){
		"pop_front": d.popFront,
		"pop_back":  d.popBack,
	}

	switch parts[0] {
	case "push_back", "push_front":
		i, _ := strconv.Atoi(parts[1])
		if err := pushFuncs[parts[0]](i); err != nil {
			fmt.Fprintln(w, "error")
		}
	case "pop_back", "pop_front":
		value, err := popFuncs[parts[0]]()
		if err != nil {
			fmt.Fprintln(w, "error")
			break
		}
		fmt.Fprintln(w, value)
	}
}

func main() {
	scanner := makeScanner()
	commands := readInt(scanner)
	deqSize := readInt(scanner)
	D := newDeq(deqSize)
	for i := 0; i < commands; i++ {
		command := readLine(scanner)
		eval(D, command, os.Stdout)
	}
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readLine(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}
