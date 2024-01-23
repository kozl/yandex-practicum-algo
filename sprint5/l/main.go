package main

import "errors"

func siftDown(heap []int, idx int) int {
	for {
		maxS, err := maxSibling(idx, heap)
		if err != nil {
			return idx
		}

		if heap[idx] > heap[maxS] {
			return idx
		}
		heap[idx], heap[maxS] = heap[maxS], heap[idx]
		idx = maxS
	}
}

var noSiblingsErr = errors.New("no siblings")

func maxSibling(idx int, heap []int) (int, error) {
	left, right := idx*2, idx*2+1
	if left > len(heap)-1 && right > len(heap)-1 {
		return 0, noSiblingsErr
	}
	if left > len(heap)-1 {
		return right, nil
	}

	if right > len(heap)-1 {
		return left, nil
	}

	maxS := left
	if heap[right] > heap[left] {
		maxS = right
	}
	return maxS, nil
}
