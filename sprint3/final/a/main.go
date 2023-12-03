package main

/*
[Посылка](https://contest.yandex.ru/contest/23815/run-report/101288967/)

# Принцип работы

Для решения я использовал следующий алгоритм:
1. Методом бинарного поиска находим номер элемента, который был наименьшим (стартовым) в кольцевом буфере.
Функция сравнения для поиска такого элемента: сравниваем текущий элемент с первым. Если он больше, значит искомый
элемент находится в правой части массива. Если меньше — в левой. Если выбранный элемент меньше предыдущего, значит
мы его нашли. Пусть номер элемента — start.
2. Вводим функцию трансляции адреса упорядоченного массива, в адрес массива, который был сдвинут на start позиций.
3. Чтобы найти элемент бинарным поиском, при доступе к массиву мы транслируем его адрес при помощи функции трансляции.
4. Полученный адрес элемента — искомый, но это адрес массива без сдвига, используя функцию трансляции получаем адрес
этого элемента в сдвинутом массиве

# Сложность

Для поиска элемента мы делаем два раза бинарный поиск, следовательно сложность по времени O(logN)
Сложность по памяти — константная
*/

func brokenSearch(array []int, n int) int {
	left, right := 0, len(array)-1
	start := binarySearch(array, left, right, func(array []int, i int) int {
		if i-1 >= 0 && array[i] < array[i-1] {
			return 0
		}
		if array[i] >= array[0] {
			return 1
		}
		return -1
	})

	pos := binarySearch(array, 0, len(array), func(array []int, i int) int {
		if array[ringPos(start, i, len(array))] == n {
			return 0
		}
		if array[ringPos(start, i, len(array))] < n {
			return 1
		}
		return -1
	})
	if array[ringPos(start, pos, len(array))] != n {
		return -1
	}
	return ringPos(start, pos, len(array))
}

// binarySearch реализует бинарный поиск, в качестве критерия поиска принимает функцию fn.
// fn должна возвращать 0 если искомый элемент найден
// 1, если следует взять правую часть половины массива
// -1 — если левую
func binarySearch(array []int, left, right int, fn func(array []int, i int) int) int {
	for left < right {
		mid := (left + right) / 2
		switch fn(array, mid) {
		case 0:
			return mid
		case 1:
			left = mid + 1
		case -1:
			right = mid
		}
	}
	return left
}

// rungPos транслирует адрес элемента в отсортированном массиве длины len без сдвига
// в адрес элемента в массиве, в котором был совершен сдвиг кольцевого массива на start позиций
func ringPos(start, pos, len int) int {
	return (pos + start) % len
}
