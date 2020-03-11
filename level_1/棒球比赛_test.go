package level_1

import (
	"testing"
)

func TestCalPoints(t *testing.T) {
	ops := []string{"C","D","9","+"}

	t.Log(CalPoints(ops))
}
