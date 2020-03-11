package level_1

/*
输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。

序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。

 

示例 1：

输入：target = 9
输出：[[2,3,4],[4,5]]
示例 2：

输入：target = 15
输出：[[1,2,3,4,5],[4,5,6],[7,8]]
 

限制：

1 <= target <= 10^5
*/

func FindContinuousSequence(target int) [][]int {
	return findContinuousSequence(target)
}

func findContinuousSequence(target int) [][]int {
	var i = 1
	var j = 1
	var sum = 0
	var array [][]int

	for i < (target-1)/2 {
		if sum < target {
			sum += j
			j++
		} else if sum > target {
			sum -= i
			i++
		} else {
			var a []int
			for n := i; n <= j; n++ {
				a = append(a, n)
			}
			array = append(array, a)
			sum -= i

			i++
		}
	}

	return array
}
