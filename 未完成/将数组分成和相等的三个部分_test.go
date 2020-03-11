package level_1

import (
	"testing"
)

func TestCanThreePartsEqualSum(t *testing.T) {
	nums := []int{0,2,1,-6,6,-7,9,1,2,0,1}

	b := canThreePartsEqualSum(nums)

	t.Log(b)

}
