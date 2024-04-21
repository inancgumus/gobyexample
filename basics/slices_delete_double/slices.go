package main

import (
	"fmt"
	"slices"
)

func main() {
	numbers := []int{0, 0, 1, 2, 1, 5, 6, 71, 9, 6, 6, 0, 1}
	numbers = deleteDuplicate(numbers)
	fmt.Println(numbers)

}

func deleteDuplicate[T comparable](numbers []T) []T {
	count := len(numbers)
	for i := 0; i < count-1; i++ {

		index := slices.Index(numbers[i+1:], numbers[i])
		fmt.Print(numbers, index)
		for index > -1 {
			fmt.Print(",", index)

			numbers = slices.Delete(numbers, index+1+i, index+1+i+1)
			count = len(numbers)
			index = slices.Index(numbers[i+1:], numbers[i])

		}
		fmt.Println(",", numbers)

	}
	return numbers
}
