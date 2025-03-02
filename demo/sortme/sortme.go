package sortme

import "fmt"

func Bubble_sort(unsorted []int) bool {

	fmt.Printf("sortme:%p\n", &unsorted)

	for pos := 0; pos < len(unsorted); pos++ {
		for subpos := pos; subpos < len(unsorted); subpos++ {
			if unsorted[subpos] < unsorted[pos] {
				tmpVal := unsorted[subpos]
				unsorted[subpos] = unsorted[pos]
				unsorted[pos] = tmpVal
			}
		}
	}
	return true
}
