package pagingsolver

import (
	"OnlineAlgorithms/internal/utils"
	"fmt"
)

type PagingSolver struct {
	size   int
	faults int
	alg    PagingSolvingAlg
	algE   utils.PagingAlg
}

func PagingSolver_Create(size int, alg int, debug bool) *PagingSolver {
	pS := &PagingSolver{size: size, faults: 0, algE: utils.PagingAlg(alg)}
	pS.createSolvingAlg(alg, debug)
	return pS
}

func (pS *PagingSolver) Serve(request int) {
	if !pS.alg.UpdateMemory(request) {
		pS.faults++
	}
}

func (ps *PagingSolver) createSolvingAlg(alg int, debug bool) {
	switch utils.PagingAlg(alg) {
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
	}
}

func (ps *PagingSolver) Raport() (string, int) {
	return fmt.Sprint(ps.algE), ps.faults
}
