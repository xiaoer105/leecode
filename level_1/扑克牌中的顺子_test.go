package level_1

import "testing"

func TestIsStraight(t *testing.T) {
	nums := []int{0, 0, 1, 2, 5}
	t.Log(isStraight(nums))
}
