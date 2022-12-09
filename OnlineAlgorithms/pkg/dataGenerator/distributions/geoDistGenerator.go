package distributions

import (
	stats "github.com/r0fls/gostats"
)

type GeoDistGenerator struct {
	gen   stats.GeometricType
	limit int
}

// GEO_Create takes bounds for distribution and returns Geometric distribution.
func GEO_Create(scale float64, limit int) *GeoDistGenerator {

	g := &GeoDistGenerator{gen: stats.Geometric(scale), limit: limit}
	return g
}

// GetRequest is implementation of GenericDataGenerator interface for Geometric distribution.
func (g *GeoDistGenerator) GetRequest() int {
	i := g.gen.Random()
	result := int(i) - 1
	if result > g.limit {
		return g.limit
	}
	return result
}
