package distributions

import (
	"math/rand"
	"time"
)

type UniDistGenerator struct {
	min int
	max int
	src *rand.Rand
}

// UNI_Create takes bounds for distribution and returns Uniform distribution.
func UNI_Create(low, high int) *UniDistGenerator {
	g := &UniDistGenerator{min: low, max: high + 1, src: rand.New(rand.NewSource(time.Now().UnixNano()))}
	return g
}

// GetRequest is implementation of GenericDataGenerator interface for Uniform distribution.
func (g *UniDistGenerator) GetRequest() int {
	i := g.src.Intn(g.max-g.min) + g.min

	return i
}
