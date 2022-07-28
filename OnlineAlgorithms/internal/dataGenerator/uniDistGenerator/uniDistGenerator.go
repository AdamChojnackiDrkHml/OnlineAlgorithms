package unidistgenerator

import (
	"time"

	"golang.org/x/exp/rand"
	uniform "gonum.org/v1/gonum/stat/distuv"
)

type UniDistGenerator struct {
	gen uniform.Uniform
}

func Create(low, high int) *UniDistGenerator {

	g := &UniDistGenerator{gen: uniform.Uniform{Min: float64(low), Max: float64(high + 1), Src: rand.New(rand.NewSource(uint64(time.Now().UnixNano())))}}
	return g
}

func (g *UniDistGenerator) GetRequest() int {
	i := g.gen.Rand()

	return int(i)
}
