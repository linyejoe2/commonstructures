package rarray

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	Insert(&arr, 434, 4)

	fmt.Println("Insert: ", arr)

	InsertV2(&arr, 23, 1)
	fmt.Println("InsertV2: ", arr)

	Delete(&arr, 1)
	// Delete(&arr, 1)
	fmt.Println("Delete: ", arr)

	fmt.Println("Find: ", Find(arr, 1))
}

func TestRandomAccess(t *testing.T) {
	arr := []interface{}{1, 2, 3, 4, 5}
	result := RandomAccess(arr)

	fmt.Printf("%d\n", result)

	// Since RandomAccess returns a random element, we just need to ensure it is in the array
	found := false
	for _, v := range arr {
		if result == v {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("RandomAccess() returned %v, which is not in the array", result)
	}
}
