package main

func merge_sort(arr []int, lf int, rg int) {
	if rg-lf <= 1 {
		return
	}
	merge_sort(arr, lf, (rg+lf)/2)
	merge_sort(arr, (rg+lf)/2, rg)

	res := merge(arr, lf, (rg+lf)/2, rg)
	idx := 0
	for i := lf; i < rg; i++ {
		arr[i] = res[idx]
		idx++
	}
}

func merge(arr []int, left int, mid int, right int) (result []int) {
	first := arr[left:mid]
	second := arr[mid:right]

	result = make([]int, len(first)+len(second))
	f, s, r := 0, 0, 0
	for f < len(first) && s < len(second) {
		if first[f] <= second[s] {
			result[r] = first[f]
			f++
		} else {
			result[r] = second[s]
			s++
		}
		r++
	}
	for f < len(first) {
		result[r] = first[f]
		f++
		r++
	}
	for s < len(second) {
		result[r] = second[s]
		s++
		r++
	}
	return result
}
