package updatelistsolveralgs

const NUM_OF_UPDATELIST_ALGS = 6

type UpdateListAlg int

const (
	MTF UpdateListAlg = iota
	TRANS
	FC
	BIT
	TS
	Combination
)

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
