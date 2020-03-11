package level_1

import (
	"testing"
)

func TestThreeSum(t *testing.T) {

	array := DistributeCandies(10, 3)
	t.Log("最终数组为:", array)

	nums:= []int{-1, 0, 1, 2, -1, -4}
	t.Log(threeSum(nums))
}