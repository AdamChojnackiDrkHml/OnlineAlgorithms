package pagingsolver

const NUM_OF_PAGING_ALGS = 6

type PagingAlg int

const (
	LRU PagingAlg = iota
	FIFO
	LFU
	MARK_LRU
	MARK_FC
	RM
)

func (e PagingAlg) String() string {
	switch e {
	case LRU:
		return "LRU"
	case FIFO:
		return "FIFO"
	case LFU:
		return "LFU"
	case MARK_LRU:
		return "MARK"
	case MARK_FC:
		return "MARK_FC"
	case RM:
		return "RM"
	default:
		return "NULL"
	}
}
