package ttr

import "container/heap"

type HeuristicItem struct {
	value interface{}
	cost  int
	index int
	h Heuristic
}

type HeuristicHeapArray []*HeuristicItem

type HeuristicHeap struct {
	items HeuristicHeapArray
}

func NewHeuristicHeap() *HeuristicHeap {
	hh := new(HeuristicHeap)
	hh.items = make(HeuristicHeapArray, 0)
	heap.Init(&hh.items)
	return hh
}

func (hh HeuristicHeap) Len() int {
	return len(hh.items)
}

func (hha HeuristicHeapArray) Len() int {
	return len(hha)
}

func (hha HeuristicHeapArray) Less(i, j int) bool {
	h := hha[i].h
	return h.Less(hha[i].cost, hha[j].cost)
}

func (hha HeuristicHeapArray) Swap(i, j int) {
	hha[i], hha[j] = hha[j], hha[i]
	hha[i].index = i
	hha[j].index = j
}

func (hh *HeuristicHeap) Push(x interface{}) {
	heap.Push(&hh.items, x)
}

func (hha *HeuristicHeapArray) Push (x interface{}) {
	n := len(*hha)
	item := x.(*HeuristicItem)
	item.index = n
	*hha = append(*hha, item)
}

func (hh *HeuristicHeap) Pop() interface{} {
	return heap.Pop(&hh.items)
}

func (hha *HeuristicHeapArray) Pop() interface{} {
	old := *hha
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*hha = old[0:n-1]
	return item
}

func (hha *HeuristicHeapArray) HeapUpdate(item *HeuristicItem, value *Track, cost int) {
	item.value = value
	item.cost = cost
	heap.Fix(hha, item.index)
}