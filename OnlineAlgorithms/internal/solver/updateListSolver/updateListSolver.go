package updatelistsolver

import "fmt"

type UpdateListAlg int

const (
	MTF UpdateListAlg = iota
	TRANS
	FQ
	PD
)

type UpdateListSolver struct {
	size int
	cost int
	alg  UpdateListSolvingAlg
}

func UpdateListSolver_Create(size int, alg UpdateListAlg, debug bool) *UpdateListSolver {
	uLS := &UpdateListSolver{size: size, cost: 0}
	uLS.createSolvingAlg(alg, debug)
	return uLS
}

func (uLS *UpdateListSolver) Serve(request int) {
	uLS.cost += uLS.alg.UpdateList(request)
}

func (uLS *UpdateListSolver) createSolvingAlg(alg UpdateListAlg, debug bool) {
	switch alg {
	case MTF:
		{
			uLS.alg = MTFAlg_Create(uLS.size, debug)
			break
		}
	case TRANS:
		{
			uLS.alg = TransAlg_Create(uLS.size, debug)
			break
		}
	case FQ:
		{
			uLS.alg = FQAlg_Create(uLS.size, debug)
			break
		}
	case PD:
		break
	}
}

func (uLS *UpdateListSolver) Raport() string {
	return fmt.Sprint(uLS.cost)
}
