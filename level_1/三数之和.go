package level_1

import (
	"log"
	"sort"
)

/*
给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。

例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]
*/

func threeSum(nums []int) [][]int {
	if 0 == len(nums) {
		return nil
	}

	sort.Ints(nums)

	prevNum := nums[0]
	var elements = []int{prevNum}

	for _, num := range nums {
		if num != prevNum {
			prevNum = num
			elements = append(elements, num)
		}
	}

	if 3 > len(elements) {
		return nil
	}

	log.Println(elements)

	var items [][]int
	for idx := 0; idx < len(elements) && idx+3 < len(elements); idx++ {
		if elements[idx]+elements[idx+1]+elements[idx+2] == 0 {
			item := []int{elements[idx], elements[idx+1], elements[idx+2]}
			items = append(items, item)
		}
	}

	return items
}

func ThreeSum(nums []int) [][]int {
	return threeSum(nums)
}
