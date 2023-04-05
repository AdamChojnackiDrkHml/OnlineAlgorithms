package distributions

import (
	"math/rand"
	"time"
)

type HrmDistGenerator struct {
	src  *rand.Rand
	dist []float64
	high int
}

// HRM_Create takes bounds for distribution and returns Harmonic distribution.
func HRM_Create(low, high int) *HrmDistGenerator {

	g := &HrmDistGenerator{src: rand.New(rand.NewSource(time.Now().UnixNano()))}
	g.dist = make([]float64, high+1)
	g.high = high

	S := 0.0

	for i := high; i >= 1; i-- {
		S += 1.0 / float64(i)
	}

	g.dist[0] = 0.0

	for i := 1; i <= high; i++ {
		g.dist[i] = g.dist[i-1] + 1.0/(float64(i)*S)
	}

	return g
}

// GetRequest is implementation of GenericDataGenerator interface for Harmonic distribution.
func (g *HrmDistGenerator) GetRequest() int {
	ran := g.src.Float64()

	for i := range g.dist {
		if g.dist[i] < ran && ran <= g.dist[i+1] {
			return i
		}
	}

	return g.high
}
