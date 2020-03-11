package level_1

import (
	"log"
	"testing"
)

func TestGame(t *testing.T) {
	x := []int{0,2,3}
	y := []int{3,2,1}

	log.Println(Game(x,y))
}