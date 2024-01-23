package main

import "errors"

func siftUp(heap []int, idx int) int {
	for {
		parentIdx, err := parent(idx)
		if err != nil {
			return idx
		}
		if heap[parentIdx] < heap[idx] {
			heap[parentIdx], heap[idx] = heap[idx], heap[parentIdx]
			idx = parentIdx
		} else {
			return idx
		}
	}
}

var noParentErr = errors.New("no parent")

func parent(idx int) (int, error) {
	if idx == 1 {
		return 0, noParentErr
	}
	return idx / 2, nil
}
