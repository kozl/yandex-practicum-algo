package main

/*
Успешная посылка: https://contest.yandex.ru/contest/24810/run-report/105007426/

Функция remove принимает ссылку на корень дерева и возвращает ссылку также на корень дерева, но уже с удалённым
элементом.
Задачу решал рекурсией: базовый случай когда нашли элемент для удаления,
в ином случае вызываем remove для одного из потомков.
Операция удаления узла из дерева делится на три случая:

1. Если удаляем лист, то ничего делать не нужно и remove может вернуть nil, элемент удалён.

2. Если удаляем узел, у которого только один из потомков, то достаточно
удалить его и сдвинуть поддерево на его место.

3. Если у узла есть оба потомка, то нам нужно удалить наименьший эелемент
из правого поддерева (можно наибольший из левого). Для этого используется функция popLeftMostNode,
которая извлекает целевой узел и возвращает ссылку на корень дерева в котором его уже нет.
Этот узел помещается на место удалённого узла.
*/

// <templatet>
type Node struct {
	value int
	left  *Node
	right *Node
}

// <template>

func remove(node *Node, key int) *Node {
	if node == nil {
		return node
	}
	if key == node.value {
		if node.left == nil && node.right == nil {
			return nil
		}
		if node.left == nil {
			return node.right
		}
		if node.right == nil {
			return node.left
		}
		var popped *Node
		node.right, popped = popLeftMostNode(node.right)
		node = popped
	} else if key > node.value {
		node.right = remove(node.right, key)
	} else {
		node.left = remove(node.left, key)
	}
	return node
}

func popLeftMostNode(node *Node) (*Node, *Node) {
	if node.left == nil {
		if node.right == nil {
			return nil, node
		}
		return node.right, node
	}
	var popped *Node
	node.left, popped = popLeftMostNode(node.left)
	return node, popped
}
