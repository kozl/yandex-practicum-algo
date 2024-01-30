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

Вычислительная сложность: O(h), в случае сбалансированного дерева O(log(n))
Сложность по памяти: O(n)
*/

// <templatet>
type Node struct {
	value int
	left  *Node
	right *Node
}

// <template>

func remove(root *Node, key int) *Node {
	if root == nil {
		return root
	}
	var (
		parent *Node
		node   *Node = root
	)
	for node != nil && key != node.value {
		if key > node.value {
			parent = node
			node = node.right
		} else {
			parent = node
			node = node.left
		}
	}
	if node == nil {
		return root
	}

	if node.left == nil && node.right == nil {
		replaceParentChildWith(parent, node, nil)
		return root
	}

	if node.left == nil {
		replaceParentChildWith(parent, node, node.right)
		return root
	}

	if node.right == nil {
		replaceParentChildWith(parent, node, node.left)
		return root
	}

	rightSubtree, popped := popLeftMostNode(node.right)
	popped.right = rightSubtree
	popped.left = node.left
	replaceParentChildWith(parent, node, popped)

	return root
}

func popLeftMostNode(root *Node) (*Node, *Node) {
	var (
		parent *Node
		node   *Node = root
	)
	for node.left != nil {
		parent = node
		node = node.left
	}

	replaceParentChildWith(parent, node, node.right)
	return root, node
}

func replaceParentChildWith(parent, child, new *Node) {
	if parent == nil {
		*child = *new
		return
	}
	if parent.left == child {
		parent.left = new
	} else {
		parent.right = new
	}
}
