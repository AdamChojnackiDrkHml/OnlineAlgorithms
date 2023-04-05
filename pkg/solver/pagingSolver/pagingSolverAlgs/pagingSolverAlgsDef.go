// Package pagingsolveralgs contains implementation of supported algorithms for Paging.
// Defines enumeration for these algorithms.
package pagingsolveralgs

const NUM_OF_PAGING_ALGS = 6

// PagingAlg will hold enumerate for supported Update List algorithms.
type PagingAlg int

// Defined in PagingAlg algorithms.
const (
	LRU PagingAlg = iota
	FIFO
	LFU
	MARK_LRU
	MARK_FC
	RM
	FWF
	RAND
)

// String creates string from PagingAlg.
func (e PagingAlg) String() string {
	switch e {
	case LRU:
		return "LRU"
	case FIFO:
		return "FIFO"
	case LFU:
		return "LFU"
	case MARK_LRU:
		return "MARK_LRU"
	case MARK_FC:
		return "MARK_FC"
	case RM:
		return "RM"
	case FWF:
		return "FWF"
	case RAND:
		return "RAND"
	default:
		return "NULL"
	}
}
