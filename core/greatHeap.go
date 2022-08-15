package core

import "fmt"

type wrappedObj struct {
	a int
}
type GreatHeap struct {
	arr         []*wrappedObj
	reversedMap map[*wrappedObj]int
	comparator  func(a, b *wrappedObj) int
}

func NewGreatHeap(arr []*wrappedObj, comparator func(a, b *wrappedObj) int) *GreatHeap {
	indexMap := make(map[*wrappedObj]int)
	for k, v := range arr {
		indexMap[v] = k
	}

	hp := &GreatHeap{
		arr:         arr,
		comparator:  comparator,
		reversedMap: indexMap,
	}

	lastEffective := (len(arr) - 2) / 2

	for i := lastEffective; i >= 0; i-- {
		hp.heapify(i)
	}

	return hp
}

func (hp *GreatHeap) swap(a, b int) {
	targetA := hp.arr[a]
	targetB := hp.arr[b]

	hp.reversedMap[targetA] = b
	hp.reversedMap[targetB] = a

	hp.arr[a] = hp.arr[b]
	hp.arr[b] = targetA
}

func (hp *GreatHeap) heapInsert(index int) {
	for hp.comparator(hp.arr[index], hp.arr[(index-1)/2]) < 0 {
		hp.swap(index, (index-1)/2)
		index = (index - 1) / 2
	}
}

func (hp *GreatHeap) heapify(index int) {
	l := index*2 + 1

	for l < hp.Size() {
		var best int = l

		if l+1 < hp.Size() && hp.comparator(hp.arr[l], hp.arr[l+1]) >= 0 {
			best = l + 1
		}

		if hp.comparator(hp.arr[index], hp.arr[best]) < 0 {
			best = index
		}

		if best == index {
			break
		}

		hp.swap(best, index)
		index = best
		l = index*2 + 1
	}
}

func (hp *GreatHeap) Pop() *wrappedObj {
	top := hp.arr[0]

	hp.swap(0, hp.Size()-1)
	hp.arr = hp.arr[:hp.Size()-1]

	delete(hp.reversedMap, top)

	hp.heapify(0)

	return top
}

func (hp *GreatHeap) Push(obj *wrappedObj) {
	hp.arr = append(hp.arr, obj)
	hp.reversedMap[obj] = hp.Size() - 1
	hp.heapInsert(hp.Size() - 1)
}

func (hp *GreatHeap) Peek() *wrappedObj {
	return hp.arr[0]
}

func (hp *GreatHeap) Remove(obj *wrappedObj) {
	idxOfRemoved, ok := hp.reversedMap[obj]

	if !ok {
		return
	}
	lastEl := hp.arr[hp.Size()-1]

	hp.swap(hp.Size()-1, idxOfRemoved)

	hp.arr = hp.arr[:hp.Size()-1]
	delete(hp.reversedMap, obj)

	if lastEl == obj {
		return
	}

	hp.resign(idxOfRemoved)

}

func (hp *GreatHeap) Size() int {
	return len(hp.arr)
}

func (hp *GreatHeap) resign(index int) {
	hp.heapInsert(index)
	hp.heapify(index)
}

func (hp *GreatHeap) Print() {
	for _, v := range hp.arr {
		fmt.Println(v)
	}
}
