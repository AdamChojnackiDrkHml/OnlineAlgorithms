package pagingsolver

import "fmt"

type PagingSolvingAlg interface {
	UpdateMemory(request int) bool
}

type PagingSolver struct {
	size   int
	faults int
	alg    PagingSolvingAlg
	algE   PagingAlg
}

func PagingSolver_Create(size int, algP PagingAlg, debug bool) *PagingSolver {
	pS := &PagingSolver{size: size, faults: 0, algE: PagingAlg(algP)}
	pS.createSolvingAlg(algP, debug)
	return pS
}

func (pS *PagingSolver) Serve(request int) {
	if !pS.alg.UpdateMemory(request) {
		pS.faults++
	}
}

func (ps *PagingSolver) createSolvingAlg(algP PagingAlg, debug bool) {
	switch PagingAlg(algP) {
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
	case MARK_LRU:
		{
			ps.alg = MARKLRUAlg_Create(ps.size, debug)
			break
		}
	case MARK_FC:
		{
			ps.alg = MARKFCAlg_Create(ps.size, debug)
			break
		}
	case RM:
		{
			ps.alg = RMAlg_Create(ps.size, debug)
		}
	}
}

func (ps *PagingSolver) Raport() (string, int) {
	return fmt.Sprint(ps.algE), ps.faults
}
