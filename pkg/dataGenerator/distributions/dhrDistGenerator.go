package distributions

import (
	"math/rand"
	"time"
)

type DhrDistGenerator struct {
	src  *rand.Rand
	dist []float64
	high int
}

// DHR_Create takes bounds for distribution and returns Diharmonic distribution.
func DHR_Create(low, high int) *DhrDistGenerator {

	g := &DhrDistGenerator{src: rand.New(rand.NewSource(time.Now().UnixNano()))}
	g.dist = make([]float64, high+1)
	g.high = high

	S := 0.0

	for i := high; i >= 1; i-- {
		S += 1.0 / (float64(i) * float64(i))
	}

	g.dist[0] = 0.0

	for i := 1; i <= high; i++ {
		g.dist[i] = g.dist[i-1] + 1.0/(float64(i)*float64(i)*S)
	}

	return g
}

// GetRequest is implementation of GenericDataGenerator interface for Diharmonic distribution.
func (g *DhrDistGenerator) GetRequest() int {
	ran := g.src.Float64()

	for i := range g.dist {
		if g.dist[i] < ran && ran <= g.dist[i+1] {
			return i
		}
	}

	return g.high
}
