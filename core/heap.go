package core

type HeapStruct struct {
	arr        []int
	comparator func(a, b int) int
}

func NewHeap(arr []int, comparator func(a, b int) int) *HeapStruct {
	hs := &HeapStruct{
		arr:        arr,
		comparator: comparator,
	}

	lastEffective := (len(arr) - 2) / 2

	for i := lastEffective; i >= 0; i-- {
		hs.heapify(i)
	}

	return hs
}

func (hs *HeapStruct) Size() int {
	return len(hs.arr)
}

func (hs *HeapStruct) Sort() []int {
	// original := []int{}
	// copy(original, hs.arr)

	res := []int{}

	for hs.Size() > 0 {
		top := hs.Pop()
		res = append(res, top)
	}

	// hs.arr = original

	return res
}

func (hs *HeapStruct) Peek() int {
	return hs.arr[0]
}

func (hs *HeapStruct) Pop() int {
	res := hs.Peek()
	swap(hs.arr, 0, hs.Size()-1)
	hs.arr = hs.arr[:hs.Size()-1]
	hs.heapify(0)
	return res
}

func (hs *HeapStruct) Push(node int) {
	hs.arr = append(hs.arr, node)
	hs.heapInsert(hs.Size() - 1)
}

func (hs *HeapStruct) heapify(node int) {
	for l, r := node*2+1, node*2+2; l < hs.Size(); l, r = node*2+1, node*2+2 {
		var best int
		if r < hs.Size() {
			if hs.comparator(hs.arr[node], hs.arr[l]) <= 0 && hs.comparator(hs.arr[node], hs.arr[r]) <= 0 {
				break
			}
			if hs.comparator(hs.arr[l], hs.arr[node]) <= 0 && hs.comparator(hs.arr[l], hs.arr[r]) <= 0 {
				best = l
			} else {
				best = r
			}
		} else {
			if hs.comparator(hs.arr[node], hs.arr[l]) <= 0 {
				break
			}
			best = l
		}
		swap(hs.arr, node, best)
		node = best
	}
}

func (hs *HeapStruct) heapInsert(node int) {
	for hs.comparator(hs.arr[node], hs.arr[(node-1)/2]) < 0 {
		swap(hs.arr, node, (node-1)/2)
		node = (node - 1) / 2
	}
}
