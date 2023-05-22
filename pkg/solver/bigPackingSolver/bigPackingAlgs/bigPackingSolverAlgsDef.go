package bigpackingalgs

const NUM_OF_BIG_PACKING_ALGS = 5

// bigPackingAlg will hold enumerate for supported Update List algorithms.
type BigPackingAlg int

// Defined in BigPackingAlg algorithms.
const (
	NF BigPackingAlg = iota
	RF
	FF
	BF
	WF
)

// String creates string from BigPackingAlg.
func (e BigPackingAlg) String() string {
	switch e {
	case NF:
		return "Next Fit"
	case RF:
		return "Random Fit"
	case FF:
		return "First Fit"
	case BF:
		return "Best Fit"
	case WF:
		return "Worst Fit"
	default:
		return "NULL"
	}
}
