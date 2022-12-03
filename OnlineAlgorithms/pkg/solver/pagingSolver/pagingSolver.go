package pagingsolver

import (
	palgs "OnlineAlgorithms/pkg/solver/pagingSolver/pagingSolverAlgs"
	"fmt"
)

type PagingSolvingAlg interface {
	UpdateMemory(request int) bool
}

type PagingSolver struct {
	size   int
	faults int
	alg    PagingSolvingAlg
	algE   palgs.PagingAlg
}

func PagingSolver_Create(size int, algP palgs.PagingAlg, debug bool) *PagingSolver {
	pS := &PagingSolver{size: size, faults: 0, algE: palgs.PagingAlg(algP)}
	pS.createSolvingAlg(algP, debug)
	return pS
}

func (pS *PagingSolver) Serve(request int) {
	if !pS.alg.UpdateMemory(request) {
		pS.faults++
	}
}

func (ps *PagingSolver) createSolvingAlg(algP palgs.PagingAlg, debug bool) {
	switch palgs.PagingAlg(algP) {
	case palgs.LRU:
		{
			ps.alg = palgs.LRUAlg_Create(ps.size, debug)
			break
		}
	case palgs.FIFO:
		{
			ps.alg = palgs.FIFOAlg_Create(ps.size, debug)
			break
		}
	case palgs.LFU:
		{
			ps.alg = palgs.LFUAlg_Create(ps.size, debug)
			break
		}
	case palgs.MARK_LRU:
		{
			ps.alg = palgs.MARKLRUAlg_Create(ps.size, debug)
			break
		}
	case palgs.MARK_FC:
		{
			ps.alg = palgs.MARKFCAlg_Create(ps.size, debug)
			break
		}
	case palgs.RM:
		{
			ps.alg = palgs.RMAlg_Create(ps.size, debug)
		}
	}
}

func (ps *PagingSolver) Raport() (string, int) {
	return fmt.Sprint(ps.algE), ps.faults
}
