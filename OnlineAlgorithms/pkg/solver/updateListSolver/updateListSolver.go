package updatelistsolver

import (
	"fmt"
)

type UpdateListSolvingAlg interface {
	UpdateList(request int) int
}

func CreateList(size int) []int {
	list := make([]int, size)

	for i := range list {
		list[i] = i
	}

	return list
}

type UpdateListSolver struct {
	size int
	cost int
	alg  UpdateListSolvingAlg
	algE UpdateListAlg
}

func UpdateListSolver_Create(size int, algUL UpdateListAlg, debug bool) *UpdateListSolver {
	uLS := &UpdateListSolver{size: size, cost: 0, algE: UpdateListAlg(algUL)}
	uLS.createSolvingAlg(algUL, debug)
	return uLS
}

func (uLS *UpdateListSolver) Serve(request int) {
	uLS.cost += uLS.alg.UpdateList(request)
}

func (uLS *UpdateListSolver) createSolvingAlg(algUL UpdateListAlg, debug bool) {
	switch algUL {
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
	case FC:
		{
			uLS.alg = FCAlg_Create(uLS.size, debug)
			break
		}
	case BIT:
		{
			uLS.alg = BITAlg_Create(uLS.size, debug)
			break
		}
	case TS:
		{
			uLS.alg = TSAlg_Create(uLS.size, debug)

			break
		}
	case Combination:
		{
			uLS.alg = CombinationAlg_Create(uLS.size, debug)
		}
	}
}

func (uLS *UpdateListSolver) Raport() (string, int) {
	return fmt.Sprint(uLS.algE), uLS.cost
}
