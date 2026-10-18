package main

import (
	"fmt"
)

type IntHeap []int

func (this IntHeap) Len() int {
	return len(this)
}

func (this IntHeap) Less(i, j int) bool {
	return this[i] < this[j]
}

func (this IntHeap) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func main() {
	var (
		dishesCount int
		preferences []int
	)

	_, err := fmt.Scan(&dishesCount)
	if err != nil {
		fmt.Println("Invalid dishes count")

		return
	}

	preferences = make([]int, dishesCount)
	for index := range dishesCount {
		_, err = fmt.Scan(&preferences[index])
		if err != nil {
			fmt.Println("Invalid preferences")

			return
		}
	}
}
