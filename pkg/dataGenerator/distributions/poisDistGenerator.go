package distributions

import (
	stats "github.com/r0fls/gostats"
)

// NOT USED
type PoisDistGenerator struct {
	gen   stats.PoissonType
	limit int
}

// POIS_Create takes bounds for distribution and returns Poisson distribution.
// NOT SUPPORTED ANYMORE
func POIS_Create(scale float64, limit int) *PoisDistGenerator {

	g := &PoisDistGenerator{gen: stats.Poisson(scale), limit: limit}
	return g
}

// GetRequest is implementation of GenericDataGenerator interface for Poisson distribution.
func (g *PoisDistGenerator) GetRequest() int {
	i := g.gen.Random()
	result := int(i)
	if result > g.limit {
		return g.limit
	}
	return result
}
