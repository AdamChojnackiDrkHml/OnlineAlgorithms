package graphs

import (
	"fmt"
	"testing"
)

func TestConversion(t *testing.T) {
	test := uint8(7)
	var T Torus
	x := T.verticeIndexToCoordinates(test)

	fmt.Println(x)

	var H Hypercube

	y := H.Distance(uint8(1), uint8(12))
	fmt.Println(y)
}
