package updatelistsolver

import (
	genUtils "OnlineAlgorithms/pkg/utils/generalUtils"
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
	algE genUtils.UpdateListAlg
}

func UpdateListSolver_Create(size int, algUL genUtils.UpdateListAlg, debug bool) *UpdateListSolver {
	uLS := &UpdateListSolver{size: size, cost: 0, algE: genUtils.UpdateListAlg(algUL)}
	uLS.createSolvingAlg(algUL, debug)
	return uLS
}

func (uLS *UpdateListSolver) Serve(request int) {
	uLS.cost += uLS.alg.UpdateList(request)
}

func (uLS *UpdateListSolver) createSolvingAlg(algUL genUtils.UpdateListAlg, debug bool) {
	switch algUL {
	case genUtils.MTF:
		{
			uLS.alg = MTFAlg_Create(uLS.size, debug)
			break
		}
	case genUtils.TRANS:
		{
			uLS.alg = TransAlg_Create(uLS.size, debug)
			break
		}
	case genUtils.FC:
		{
			uLS.alg = FCAlg_Create(uLS.size, debug)
			break
		}
	case genUtils.BIT:
		{
			uLS.alg = BITAlg_Create(uLS.size, debug)
			break
		}
	case genUtils.TS:
		{
			uLS.alg = TSAlg_Create(uLS.size, debug)

			break
		}
	case genUtils.Combination:
		{
			uLS.alg = CombinationAlg_Create(uLS.size, debug)
		}
	}
}

func (uLS *UpdateListSolver) Raport() (string, int) {
	return fmt.Sprint(uLS.algE), uLS.cost
}
