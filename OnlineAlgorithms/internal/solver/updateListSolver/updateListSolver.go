package updatelistsolver

import "fmt"

type UpdateListAlg int

const (
	LRU UpdateListAlg = iota
	FIFO
	LFU
	PD
)

type UpdateListSolver struct {
	size   int
	faults int
	alg    UpdateListSolvingAlg
}

func UpdateListSolver_Create(size int, alg UpdateListAlg, debug bool) *UpdateListSolver {
	pS := &UpdateListSolver{size: size, faults: 0}
	pS.createSolvingAlg(alg, debug)
	return pS
}

func (pS *UpdateListSolver) Serve(request int) {
	if !pS.alg.UpdateMemory(request) {
		pS.faults++
	}
}

func (ps *UpdateListSolver) createSolvingAlg(alg UpdateListAlg, debug bool) {
	switch alg {
	case LRU:
		{
			// ps.alg = LRUAlg_Create(ps.size, debug)
			break
		}
	case FIFO:
		{
			// ps.alg = FIFOAlg_Create(ps.size, debug)
			break
		}
	case LFU:
		{
			// ps.alg = LFUAlg_Create(ps.size, debug)
			break
		}
	case PD:
		break
	}
}

func (ps *UpdateListSolver) Raport() string {
	return fmt.Sprint(ps.faults)
}
