package pagingsolver

import "fmt"

type PagingAlg int

const (
	LRU PagingAlg = iota
	FIFO
	LFU
	PD
)

func (e PagingAlg) String() string {
	switch e {
	case LRU:
		return "LRU"
	case FIFO:
		return "FIFO"
	case LFU:
		return "LFU"
	default:
		return "NULL"
	}
}

type PagingSolver struct {
	size   int
	faults int
	alg    PagingSolvingAlg
	algE   PagingAlg
}

func PagingSolver_Create(size int, alg int, debug bool) *PagingSolver {
	pS := &PagingSolver{size: size, faults: 0, algE: PagingAlg(alg)}
	pS.createSolvingAlg(alg, debug)
	return pS
}

func (pS *PagingSolver) Serve(request int) {
	if !pS.alg.UpdateMemory(request) {
		pS.faults++
	}
}

func (ps *PagingSolver) createSolvingAlg(alg int, debug bool) {
	switch PagingAlg(alg) {
	case LRU:
		{
			ps.alg = LRUAlg_Create(ps.size, debug)
			break
		}
	case FIFO:
		{
			ps.alg = FIFOAlg_Create(ps.size, debug)
			break
		}
	case LFU:
		{
			ps.alg = LFUAlg_Create(ps.size, debug)
			break
		}
	case PD:
		break
	}
}

func (ps *PagingSolver) Raport() (string, int) {
	return fmt.Sprint(ps.algE), ps.faults
}
