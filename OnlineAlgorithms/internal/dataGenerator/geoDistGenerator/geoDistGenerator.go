package geodistgenerator

import (
	stats "github.com/r0fls/gostats"
)

type UniDistGenerator struct {
	gen stats.GeometricType
}

func Create(scale float64) *UniDistGenerator {

	g := &UniDistGenerator{gen: stats.Geometric(scale)}
	return g
}

func (g *UniDistGenerator) GetRequest() int {
	i := g.gen.Random()

	return int(i) - 1
}
