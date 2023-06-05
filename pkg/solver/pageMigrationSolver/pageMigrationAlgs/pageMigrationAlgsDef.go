package pagemigrationalgs

const NUM_OF_BIG_PACKING_ALGS = 2

// PageMigrationAlgs will hold enumerate for supported Update List algorithms.
type PageMigrationAlgs int

// Defined in PageMigrationAlgs algorithms.
const (
	MTM PageMigrationAlgs = iota
	F
)

const D = 32

// String creates string from PageMigrationAlgs.
func (e PageMigrationAlgs) String() string {
	switch e {
	case MTM:
		return "Move to Min"
	case F:
		return "Flip"
	default:
		return "NULL"
	}
}
