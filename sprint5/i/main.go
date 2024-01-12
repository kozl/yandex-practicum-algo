package main

import (
	"fmt"
	"strconv"
)

func solve(n int) int {
	nums := []int{}
	for i := 1; i <= n; i++ {
		nums = append(nums, i)
	}
	return countTrees(nums)
}

func countTrees(nums []int) int {
	if len(nums) == 1 || len(nums) == 0 {
		return 1
	}
	count := 0
	for idx := range nums {
		leftct := countTrees(nums[:idx])
		rightct := countTrees(nums[idx+1:])
		count += leftct * rightct
	}
	return count
}

func main() {
	n := readInt()
	fmt.Println(solve(n))
}

func readInt() int {
	var aString string
	fmt.Scan(&aString)
	a, _ := strconv.Atoi(aString)
	return a
}
