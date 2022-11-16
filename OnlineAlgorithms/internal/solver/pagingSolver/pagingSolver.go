package pagingsolver

import (
	"OnlineAlgorithms/internal/utils"
	"fmt"
)

type PagingSolvingAlg interface {
	UpdateMemory(request int) bool
}

type PagingSolver struct {
	size   int
	faults int
	alg    PagingSolvingAlg
	algE   utils.PagingAlg
}

func PagingSolver_Create(size int, algP utils.PagingAlg, debug bool) *PagingSolver {
	pS := &PagingSolver{size: size, faults: 0, algE: utils.PagingAlg(algP)}
	pS.createSolvingAlg(algP, debug)
	return pS
}

func (pS *PagingSolver) Serve(request int) {
	if !pS.alg.UpdateMemory(request) {
		pS.faults++
	}
}

func (ps *PagingSolver) createSolvingAlg(algP utils.PagingAlg, debug bool) {
	switch utils.PagingAlg(algP) {
	case utils.LRU:
		{
			ps.alg = LRUAlg_Create(ps.size, debug)
			break
		}
	case utils.FIFO:
		{
			ps.alg = FIFOAlg_Create(ps.size, debug)
			break
		}
	case utils.LFU:
		{
			ps.alg = LFUAlg_Create(ps.size, debug)
			break
		}
	case utils.MARK:
		{
			ps.alg = MARKAlg_Create(ps.size, debug)
			break
		}
	case utils.MARK2:
		{
			ps.alg = MARK2Alg_Create(ps.size, debug)
			break
		}
	case utils.RM:
		{
			ps.alg = RMAlg_Create(ps.size, debug)
		}
	}
}

func (ps *PagingSolver) Raport() (string, int) {
	return fmt.Sprint(ps.algE), ps.faults
}
