package rarray

import (
	"math/rand"
)

func RandomAccess(array []interface{}) (ele interface{}) {
	randomIndex := rand.Intn(len(array))

	ele = array[randomIndex]
	return
}

// only accept int array
func Insert(nums *[]int, num int, index int) {
	arr := *nums
	for i := index; i <= len(arr); i++ {
		if i == len(arr) {
			*nums = append(arr, num)
			return
		}
		temp := arr[i]
		arr[i] = num
		num = temp
	}
}

// Insert adds an integer `num` into the slice `nums` at the specified `index`.
func InsertV2(nums *[]int, num int, index int) {
	arr := *nums

	// Validate the index
	if index < 0 || index > len(arr) {
		panic("index out of range")
	}

	// Insert the element using slicing
	*nums = append(arr[:index], append([]int{num}, arr[index:]...)...)
}

func Delete[T any](nums *[]T, index int) {
	arr := *nums

	if index == 0 {
		*nums = arr[1:]
		return
	}

	if index == (len(arr) - 1) {
		*nums = arr[:index]
		return
	}

	*nums = append(arr[:index], arr[index+1:]...)
}

// find a number in nums, if number not exist in nums, return -1.
func Find(nums []int, target int) (index int) {
	index = -1

	for i := range len(nums) {
		if nums[i] == target {
			index = i
		}
	}

	return
}
