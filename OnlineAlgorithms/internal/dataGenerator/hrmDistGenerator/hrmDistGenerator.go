package hrmdistgenerator

import (
	"time"

	"golang.org/x/exp/rand"
	uniform "gonum.org/v1/gonum/stat/distuv"
)

type HrmDistGenerator struct {
	gen  uniform.Uniform
	dist []float64
	high int
}

func Create(low, high int) *HrmDistGenerator {

	g := &HrmDistGenerator{gen: uniform.Uniform{Min: float64(0), Max: float64(1), Src: rand.New(rand.NewSource(uint64(time.Now().UnixNano())))}}
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

func (g *HrmDistGenerator) GetRequest() int {
	ran := g.gen.Rand()

	for i := range g.dist {
		if g.dist[i] < ran && ran <= g.dist[i+1] {
			return i
		}
	}

	return g.high
}
