package commonstructures

import (
	"math/rand"
)

func RandomAccess(nums []int) (randomNum int) {
	randomIndex := rand.Intn(len(nums))

	randomNum = nums[randomIndex]
	return
}

