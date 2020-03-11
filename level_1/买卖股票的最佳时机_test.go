package level_1

import "testing"

func TestMaxProfit(t *testing.T) {
	n:= []int{7,1,5,3,6,4}
	max :=  MaxProfit(n)
	t.Log("max = ",max)
}
