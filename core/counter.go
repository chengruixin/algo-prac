package core

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

func copySlice(slice []int) []int {
	res := make([]int, len(slice))
	for i, num := range slice {
		res[i] = num
	}
	return res
}

func randGenArr(arrLen int) []int {
	res := make([]int, arrLen)
	for i := range res {
		res[i] = rand.Intn(1e6)
	}
	return res
}

func areEquals(arr1, arr2 []int) bool {
	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}
func Test() {
	arr := randGenArr(3e7)
	test0 := copySlice(arr)
	test1 := copySlice(arr)
	test2 := copySlice(arr)

	sg := sync.WaitGroup{}

	sg.Add(3)

	go func() {
		s := time.Now()
		sort.Ints(test0)
		fmt.Println("builtin", time.Since(s))
		sg.Done()
	}()

	go func() {
		s := time.Now()
		RadixSort(test1)
		fmt.Println("radix", time.Since(s))
		sg.Done()
	}()

	go func() {
		s := time.Now()
		test2 = NewHeap(test2, func(a, b int) int {
			return a - b
		}).Sort()
		fmt.Println("heap", time.Since(s))
		sg.Done()
	}()

	sg.Wait()
	fmt.Println(areEquals(test0, test1), areEquals(test1, test2))
}

func HeapTest() {
	hps := NewHeap([]int{45, 23, 12, 34, 21, 45}, func(a, b int) int {
		return a - b
	})

	hps.Push(23)
	hps.Push(85)
	hps.Push(255)
	hps.Push(431)
	hps.Push(1)

	fmt.Println(hps.Sort())

}
