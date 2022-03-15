package main

import (
	"container/heap"
	"fmt"
)

type IntegerHeap []int

// Len returns the lenght of the integerHeap
func (iheap IntegerHeap) Len() int { return len(iheap) }

// Less checks if the element at i is less than the element at j
func (iheap IntegerHeap) Less(i int, j int) bool { return iheap[i] < iheap[j] }

// Swap interchanges the elements at i and j
func (iheap IntegerHeap) Swap(i int, j int) { iheap[i], iheap[j] = iheap[j], iheap[i] }

// Push appends the element to the heap
func (iheap *IntegerHeap) Push(heapintf interface{}) {
	*iheap = append(*iheap, heapintf.(int))
}

// Pop removes the last element in the heap
func (iheap *IntegerHeap) Pop() interface{} {
	var n int
	var x1 int
	var previous IntegerHeap = *iheap
	n = len(previous)
	x1 = previous[n-1]
	*iheap = previous[0 : n-1]
	return x1
}

func main() {
	var intHeap *IntegerHeap = &IntegerHeap{1, 4, 5}
	heap.Init(intHeap)
	heap.Push(intHeap, 2)
	for intHeap.Len() > 0 {
		fmt.Printf("minimum: %d\n", (*intHeap)[0])
		fmt.Printf("%d\n", heap.Pop(intHeap))
	}

}
