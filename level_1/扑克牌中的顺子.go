package level_1

import (
	"sort"
)

/*
从扑克牌中随机抽5张牌，判断是不是一个顺子，即这5张牌是不是连续的。2～10为数字本身，A为1，J为11，Q为12，K为13，而大、小王为 0 ，可以看成任意数字。A 不能视为 14。

示例 1:

输入: [1,2,3,4,5]
输出: True
 

示例 2:

输入: [0,0,1,2,5]
输出: True
 

限制：

数组长度为 5 

数组的数取值为 [0, 13] .
*/

/*
   使用排序方式
   连续５张牌，意味着不能出现重复的牌
   最大值和最小值的差值一定小于５
*/
func isStraight(nums []int) bool {
	// 对牌进行排序
	sort.Ints(nums)

	count := 0

	for i := 0; i < 4; i++ {
		if 0 == nums[i] { // 大小王可以充当任何数字，直接跳过
			continue
		}

		if nums[i] == nums[i+1] { // 出现对子直接返回 false
			return false
		}

		// 计算前一张和后一张牌的差值
		count += nums[i+1] - nums[i]
	}

	//　总数小于5 是顺子
	return count < 5
}


/*
   不排序
   连续５张牌，意味着不能出现重复的牌
   最大值和最小值的差值一定小于５
*/
//func isStraight(nums []int) bool {
//	m := make(map[int]struct{}, 5)
//
//	max := math.MinInt8
//	min := math.MaxInt8
//
//	for idx := 0; idx < len(nums); idx++ {
//		if _, ok := m[nums[idx]]; ok { // 存在相同的数字, 直接返回 false
//			return false
//		}
//
//		if 0 == nums[idx]{
//			continue
//		}
//
//		if max < nums[idx]{
//			max = nums[idx]
//		}
//
//		if min > nums[idx]{
//			min = nums[idx]
//		}
//
//		m[nums[idx]] = struct{}{}
//	}
//
//	return max - min < 5
//}

//func isStraight(nums []int) bool {
//	m := make(map[int]int, 14)
//
//	for idx := 0; idx < len(nums); idx++ {
//		m[nums[idx]] ++
//	}
//
//	for idx := 1; idx < 14; idx++ {
//		if m[idx] > 1 { //　存在对子
//			return false
//		}
//	}
//
//	min := 1
//	max := 13
//
//	for min < 14 && m[min] == 0 {
//		min++
//	}
//
//	for max >= 0 && m[max] == 0 {
//		max--
//	}
//
//	return max-min <= 4
//}