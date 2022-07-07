package unidistgenerator

import (
	"fmt"
	"testing"
)

func TestUniformDist(t *testing.T) {
	g := Craete(0, 100)
	counter := make([]int, 101)
	for i := 0; i < 10000; i++ {
		counter[g.GetRequest()]++

	}

	fmt.Println(counter)
	fmt.Println("aa")
}
