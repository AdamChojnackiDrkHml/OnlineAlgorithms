// Package updatelistsolveralgs contains implementation of supported algorithms for Update List.
// Defines enumeration for these algorithms.
package updatelistsolveralgs

const NUM_OF_UPDATELIST_ALGS = 6

// UpdateListAlg will hold enumerate for supported Update List algorithms.
type UpdateListAlg int

// Defined in UpdateListAlg algorithms.
const (
	MTF UpdateListAlg = iota
	TRANS
	FC
	BIT
	TS
	Combination
)

// String creates string from UpdateListAlg.
func (e UpdateListAlg) String() string {
	switch e {
	case MTF:
		return "MTF"
	case TRANS:
		return "TRANS"
	case FC:
		return "FC"
	case BIT:
		return "BIT"
	case TS:
		return "TS"
	case Combination:
		return "Combination"
	default:
		return "NULL"
	}
}
