package level_1

import "testing"

func TestDistributeCandies(t *testing.T) {
	array := DistributeCandies(10, 3)
	t.Log("最终数组为:", array)
}
