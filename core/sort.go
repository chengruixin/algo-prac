package core

import (
	"math/rand"
)

func swap(nums []int, a int, b int) {
	temp := nums[a]
	nums[a] = nums[b]
	nums[b] = temp
}

func BubbleSort(nums []int) {
	for size := len(nums); size > 0; size-- {
		for i := 0; i < size-1; i++ {
			if nums[i] > nums[i+1] {
				swap(nums, i, i+1)
			}
		}
	}
}

func InsertionSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		for j := i; j > 0; j-- {
			if nums[j] < nums[j-1] {
				swap(nums, j, j-1)
			} else {
				break
			}
		}
	}
}

func SelectionSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		smallestIdx := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[smallestIdx] {
				smallestIdx = j
			}
		}
		swap(nums, i, smallestIdx)
	}
}

func MergeSort(nums []int) {
	mSortProcess(nums, 0, len(nums)-1)
}

func mSortProcess(nums []int, l, r int) {
	if l == r {
		return
	}

	m := l + (r-l)>>1

	mSortProcess(nums, l, m)
	mSortProcess(nums, m+1, r)

	mergeSubs(nums, l, m, r)
}

func mergeSubs(nums []int, l, m, r int) {
	p1, p2 := l, m+1
	help := make([]int, r-l+1)
	i := 0
	for p1 <= m && p2 <= r {
		if nums[p1] < nums[p2] {
			help[i] = nums[p1]
			i++
			p1++
		} else {
			help[i] = nums[p2]
			i++
			p2++
		}
	}

	for p1 <= m {
		help[i] = nums[p1]
		p1++
		i++
	}

	for p2 <= r {
		help[i] = nums[p2]
		p2++
		i++
	}

	for i, v := range help {
		nums[i+l] = v
	}
}

func QuickSort1(nums []int) {
	quickSortProcess1(nums, 0, len(nums)-1)
}

func QuickSort2(nums []int) {
	quickSortProcess2(nums, 0, len(nums)-1)
}

func quickSortProcess1(nums []int, l, r int) {
	if l >= r {
		return
	}
	pivot := partition1(nums, l, r)
	quickSortProcess1(nums, l, pivot-1)
	quickSortProcess1(nums, pivot+1, r)
}

func quickSortProcess2(nums []int, l, r int) {
	if l >= r {
		return
	}

	pivotLeft, pivotRight := partition2(nums, l, r)

	quickSortProcess2(nums, l, pivotLeft-1)
	quickSortProcess2(nums, pivotRight+1, r)
}

func partition1(nums []int, l, r int) int {
	pivot := rand.Intn(r-l+1) + l
	wall := l
	for i := l; i <= r; i++ {
		if nums[i] < nums[pivot] {
			swap(nums, i, wall)

			if pivot == wall {
				pivot = i
			}

			wall++
		}
	}

	swap(nums, wall, pivot)

	return wall
}

func partition2(nums []int, l, r int) (lp, rp int) {
	lp, rp = l, r
	pivot := nums[rand.Intn(r-l+1)+l]
	for i := lp; i <= rp; {
		if nums[i] > pivot {
			swap(nums, i, rp)
			rp--
		} else if nums[i] == pivot {
			i++
		} else {
			swap(nums, i, lp)
			lp++
			i++
		}
	}
	return
}