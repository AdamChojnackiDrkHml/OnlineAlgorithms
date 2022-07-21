package pagingsolver

import "fmt"

type PagingAlg int

const (
	PA PagingAlg = iota
	PB
	PC
	PD
)

type PagingSolver struct {
	size   int
	faults int
	alg    PagingSolvingAlg
}

func PagingSolver_Create(size int, alg PagingAlg, debug bool) *PagingSolver {
	pS := &PagingSolver{size: size, faults: 0}
	pS.createSolvingAlg(alg, debug)
	return pS
}

func (pS *PagingSolver) Serve(request int) {
	if !pS.alg.UpdateMemory(request) {
		pS.faults++
	}
}

func (ps *PagingSolver) createSolvingAlg(alg PagingAlg, debug bool) {
	switch alg {
	case PA:
		{
			ps.alg = LRUAlg_Create(ps.size, debug)
			break
		}
	case PB:
		{
			ps.alg = FIFOAlg_Create(ps.size, debug)
			break
		}
	case PC:
		{
			ps.alg = LFUAlg_Create(ps.size, debug)
			break
		}
	case PD:
		break
	}
}

func (ps *PagingSolver) Raport() string {
	return fmt.Sprint(ps.faults)
}
