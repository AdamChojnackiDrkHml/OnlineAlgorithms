package dhrdistgenerator

import (
	"fmt"
	"testing"
)

func TestDiharmonicDist(t *testing.T) {
	g := Create(0, 10)
	counter := make([]int, 101)
	for i := 0; i < 10000; i++ {
		counter[g.GetRequest()]++

	}

	fmt.Println(counter)
	fmt.Println("aa")
}
