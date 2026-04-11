package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sortPartition(part []int, ch chan []int) {
	fmt.Println("Sorting partition:", part)
	sort.Ints(part)
	ch <- part
}

func main() {
	fmt.Println("Enter a series of integers separated by spaces:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	strs := strings.Fields(input)
	var nums []int
	for _, s := range strs {
		n, err := strconv.Atoi(s)
		if err == nil {
			nums = append(nums, n)
		}
	}

	if len(nums) < 4 {
		fmt.Println("Please enter at least 4 integers.")
		return
	}

	size := len(nums) / 4
	remainder := len(nums) % 4
	partitions := make([][]int, 4)
	start := 0
	for i := 0; i < 4; i++ {
		end := start + size
		if i < remainder {
			end++
		}
		partitions[i] = nums[start:end]
		start = end
	}

	ch := make(chan []int, 4)
	for _, part := range partitions {
		go sortPartition(part, ch)
	}

	var sortedParts [][]int
	for i := 0; i < 4; i++ {
		sortedParts = append(sortedParts, <-ch)
	}

	result := mergeSortedArrays(sortedParts[0], sortedParts[1])
	result = mergeSortedArrays(result, sortedParts[2])
	result = mergeSortedArrays(result, sortedParts[3])

	fmt.Println("Sorted array:", result)
}

func mergeSortedArrays(a, b []int) []int {
	result := make([]int, 0, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}
	result = append(result, a[i:]...)
	result = append(result, b[j:]...)
	return result
}
