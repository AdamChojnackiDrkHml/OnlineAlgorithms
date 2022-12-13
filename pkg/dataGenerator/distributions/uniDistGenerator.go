package distributions

import (
	"time"

	"golang.org/x/exp/rand"
	uniform "gonum.org/v1/gonum/stat/distuv"
)

type UniDistGenerator struct {
	gen uniform.Uniform
}

// UNI_Create takes bounds for distribution and returns Uniform distribution.
func UNI_Create(low, high int) *UniDistGenerator {

	g := &UniDistGenerator{gen: uniform.Uniform{Min: float64(low), Max: float64(high + 1), Src: rand.New(rand.NewSource(uint64(time.Now().UnixNano())))}}
	return g
}

// GetRequest is implementation of GenericDataGenerator interface for Uniform distribution.
func (g *UniDistGenerator) GetRequest() int {
	i := g.gen.Rand()

	return int(i)
}
