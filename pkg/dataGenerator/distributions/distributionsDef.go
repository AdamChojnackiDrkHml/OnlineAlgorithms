// Package distributions contains implementation of supported distributions.
// Defines enumeration for these distributions.
package distributions

const NUM_OF_DISTRIBUTIONS = 5

// GeneratorTypeEnum will hold enumerate for supported distributions.
type GeneratorTypeEnum int

// Defined in GeneratorTypeEnum algorithms.
const (
	Uni GeneratorTypeEnum = iota
	Geo
	Pois
	Hrm
	Dhr
)

// String creates string from GeneratorTypeEnum.
func (e GeneratorTypeEnum) String() string {
	switch e {
	case Uni:
		return "Uni"
	case Geo:
		return "Geo"
	case Pois:
		return "Pois"
	case Hrm:
		return "Hrm"
	case Dhr:
		return "Dhr"
	default:
		return "NULL"
	}
}
