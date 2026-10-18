package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (thisHeap *IntHeap) Len() int {
	return len(*thisHeap)
}

func (thisHeap *IntHeap) Less(i, j int) bool {
	return (*thisHeap)[i] < (*thisHeap)[j]
}

func (thisHeap *IntHeap) Swap(i, j int) {
	(*thisHeap)[i], (*thisHeap)[j] = (*thisHeap)[j], (*thisHeap)[i]
}

func (thisHeap *IntHeap) Push(item any) {
	value, ok := item.(int)
	if !ok {
		return
	}

	*thisHeap = append(*thisHeap, value)
}

func (thisHeap *IntHeap) Pop() any {
	var (
		oldHeap    = *thisHeap
		size       = len(oldHeap)
		poppedItem = oldHeap[size-1]
	)

	*thisHeap = oldHeap[0 : size-1]

	return poppedItem
}

func main() {
	var (
		dishesCount int
		targetIndex int
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

	_, err = fmt.Scan(&targetIndex)
	if err != nil {
		fmt.Println("Invalid target index")

		return
	}

	minHeap := &IntHeap{}
	heap.Init(minHeap)

	for _, rating := range preferences {
		if minHeap.Len() < targetIndex {
			heap.Push(minHeap, rating)
		} else if rating > (*minHeap)[0] {
			heap.Pop(minHeap)
			heap.Push(minHeap, rating)
		}
	}

	if minHeap.Len() > 0 {
		fmt.Println((*minHeap)[0])
	}
}
