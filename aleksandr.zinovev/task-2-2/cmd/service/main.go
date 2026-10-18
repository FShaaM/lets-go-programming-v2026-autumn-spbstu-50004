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

func (this *IntHeap) Push(item any) {
	value, ok := item.(int)
	if !ok {
		return
	}

	*this = append(*this, value)
}

func (this *IntHeap) Pop() any {
	var (
		oldHeap    = *this
		heapLength = len(oldHeap)
		poppedItem = oldHeap[heapLength-1]
	)

	*this = oldHeap[0 : heapLength-1]

	return poppedItem
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
