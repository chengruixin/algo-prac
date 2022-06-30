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
		res[i] = rand.Int()
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
	arr := randGenArr(2e7)

	test0 := copySlice(arr)
	test1 := copySlice(arr)
	test2 := copySlice(arr)
	test3 := copySlice(arr)
	test4 := copySlice(arr)

	sg := sync.WaitGroup{}

	sg.Add(5)
	go func() {
		s0 := time.Now()
		//BubbleSort(test0)
		fmt.Println("bubble", time.Since(s0))
		sg.Done()
	}()

	go func() {
		s1 := time.Now()
		MergeSort(test1)
		fmt.Println("merge", time.Since(s1))
		sg.Done()

	}()

	go func() {
		s2 := time.Now()
		QuickSort1(test2)
		fmt.Println("q1", time.Since(s2))
		sg.Done()

	}()

	go func() {
		s3 := time.Now()
		QuickSort2(test3)
		fmt.Println("q2", time.Since(s3))
		sg.Done()

	}()

	go func() {
		s4 := time.Now()
		sort.Ints(test4)
		fmt.Println("builtin", time.Since(s4))
		sg.Done()

	}()

	sg.Wait()
	fmt.Println(areEquals(test3, test4), areEquals(test1, test2), areEquals(test3, test2), areEquals(test0, test2))
}
