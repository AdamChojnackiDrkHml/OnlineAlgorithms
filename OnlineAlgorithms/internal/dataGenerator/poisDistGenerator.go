package datagenerator

import (
	stats "github.com/r0fls/gostats"
)

type PoisDistGenerator struct {
	gen   stats.PoissonType
	limit int
}

func POIS_Create(scale float64, limit int) *PoisDistGenerator {

	g := &PoisDistGenerator{gen: stats.Poisson(scale), limit: limit}
	return g
}

func (g *PoisDistGenerator) GetRequest() int {
	i := g.gen.Random()
	result := int(i)
	if result > g.limit {
		return g.limit
	}
	return result
}
