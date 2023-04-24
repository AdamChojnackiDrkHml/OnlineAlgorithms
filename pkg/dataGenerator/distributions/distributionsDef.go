// Package distributions contains implementation of supported distributions.
// Defines enumeration for these distributions.
package distributions

const NUM_OF_DISTRIBUTIONS = 4

// GeneratorTypeEnum will hold enumerate for supported distributions.
type GeneratorTypeEnum int

// Defined in GeneratorTypeEnum algorithms.
const (
	Uni GeneratorTypeEnum = iota
	Geo
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

	case Hrm:
		return "Hrm"
	case Dhr:
		return "Dhr"
	default:
		return "NULL"
	}
}

func FromInt(j int) string {
	switch j {
	case 0:
		return "Uni"
	case 1:
		return "Geo"
	case 2:
		return "Hrm"
	case 3:
		return "Dhr"
	default:
		return "Null"
	}
}
