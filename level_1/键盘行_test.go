package level_1

import "testing"

func TestFindWords(t *testing.T) {
	strs := []string{"Hello", "Alaska", "Dad", "Peace"}
	t.Log(findWords(strs))
}