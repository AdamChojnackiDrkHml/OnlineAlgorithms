package pagingsolver

import (
	genUtils "OnlineAlgorithms/internal/utils/generalUtils"
	"fmt"
)

type PagingSolvingAlg interface {
	UpdateMemory(request int) bool
}

type PagingSolver struct {
	size   int
	faults int
	alg    PagingSolvingAlg
	algE   genUtils.PagingAlg
}

func PagingSolver_Create(size int, algP genUtils.PagingAlg, debug bool) *PagingSolver {
	pS := &PagingSolver{size: size, faults: 0, algE: genUtils.PagingAlg(algP)}
	pS.createSolvingAlg(algP, debug)
	return pS
}

func (pS *PagingSolver) Serve(request int) {
	if !pS.alg.UpdateMemory(request) {
		pS.faults++
	}
}

func (ps *PagingSolver) createSolvingAlg(algP genUtils.PagingAlg, debug bool) {
	switch genUtils.PagingAlg(algP) {
	case genUtils.LRU:
		{
			ps.alg = LRUAlg_Create(ps.size, debug)
			break
		}
	case genUtils.FIFO:
		{
			ps.alg = FIFOAlg_Create(ps.size, debug)
			break
		}
	case genUtils.LFU:
		{
			ps.alg = LFUAlg_Create(ps.size, debug)
			break
		}
	case genUtils.MARK:
		{
			ps.alg = MARKAlg_Create(ps.size, debug)
			break
		}
	case genUtils.MARK2:
		{
			ps.alg = MARK2Alg_Create(ps.size, debug)
			break
		}
	case genUtils.RM:
		{
			ps.alg = RMAlg_Create(ps.size, debug)
		}
	}
}

func (ps *PagingSolver) Raport() (string, int) {
	return fmt.Sprint(ps.algE), ps.faults
}
