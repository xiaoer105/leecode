package 动态规划

import (
	"math"
	"testing"
)

func cut(p []int, n int) int {
	dp := make([]int, n+1)

	for i := 1; i <= len(p); i++ { // 得到每一段最优的切割

		max := math.MinInt8

		for j := 1; j < i; j++ {
			max = p[j-1] + dp[i-j]
			if max > dp[n-i] {
				dp[i] = max
			} else {
				dp[i] = dp[i-j]
			}
		}
	}

	return dp[n]
}

func TestCoinChange(t *testing.T) {
	p := []int{1, 5, 8, 9}
	n := 4

	t.Log(cut(p, n))

	//n := fid(6)
	//t.Log(n)

	//var n = 10

	//memo := make([]int, n+1)

	//for idx := 1; idx <= n; idx++ {
	//	memo[idx] = -1
	//}
	//
	//t.Log(fidEx(n, memo))

	//memo[1] = 1
	//memo[2] = 1
	//
	//for idx := 3; idx <= n; idx++ {
	//	memo[idx] = memo[idx - 1] + memo[idx - 2]
	//}

	//t.Log(memo[n])

	//var before = 1
	//var later = 1
	//
	//for idx := 3; idx <= n; idx++ {
	//	tmp := later
	//	later = before + tmp
	//	before = tmp
	//}
	//
	//t.Log(later)
}

/*
	斐波拉契数列
	1,1,2,3,5,8,13,21,34,55...

	n == 1	// 1
	n == 1	// 2

	n = (n-1) +  (n-2) // else
*/
func fidEx(num int, memo []int) int {

	if memo[num] != -1 { // 存在，已经计算过了
		return memo[num]
	}

	if num <= 2 { // n == 1 , n == 2
		memo[num] = 1
	} else {
		memo[num] = fidEx(num-1, memo) + fidEx(num-2, memo)
	}

	return memo[num]
}

func fid(num int) int {
	if 1 == num || 2 == num {
		return 1
	}

	return fid(num-1) + fid(num-2)
}
