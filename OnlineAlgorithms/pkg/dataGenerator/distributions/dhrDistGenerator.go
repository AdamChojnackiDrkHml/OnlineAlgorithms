package distributions

import (
	"time"

	"golang.org/x/exp/rand"
	uniform "gonum.org/v1/gonum/stat/distuv"
)

type DhrDistGenerator struct {
	gen  uniform.Uniform
	dist []float64
	high int
}

// DHR_Create takes bounds for distribution and returns Diharmonic distribution.
func DHR_Create(low, high int) *DhrDistGenerator {

	g := &DhrDistGenerator{gen: uniform.Uniform{Min: float64(0), Max: float64(1), Src: rand.New(rand.NewSource(uint64(time.Now().UnixNano())))}}
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
	ran := g.gen.Rand()

	for i := range g.dist {
		if g.dist[i] < ran && ran <= g.dist[i+1] {
			return i
		}
	}

	return g.high
}
