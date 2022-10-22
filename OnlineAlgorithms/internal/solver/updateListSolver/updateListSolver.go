package updatelistsolver

import (
	"OnlineAlgorithms/internal/utils"
	"fmt"
)

type UpdateListSolver struct {
	size int
	cost int
	alg  UpdateListSolvingAlg
	algE utils.UpdateListAlg
}

func UpdateListSolver_Create(size int, alg int, debug bool) *UpdateListSolver {
	uLS := &UpdateListSolver{size: size, cost: 0, algE: utils.UpdateListAlg(alg)}
	uLS.createSolvingAlg(alg, debug)
	return uLS
}

func (uLS *UpdateListSolver) Serve(request int) {
	uLS.cost += uLS.alg.UpdateList(request)
}

func (uLS *UpdateListSolver) createSolvingAlg(alg int, debug bool) {
	switch utils.UpdateListAlg(alg) {
	case utils.MTF:
		{
			uLS.alg = MTFAlg_Create(uLS.size, debug)
			break
		}
	case utils.TRANS:
		{
			uLS.alg = TransAlg_Create(uLS.size, debug)
			break
		}
	case utils.FQ:
		{
			uLS.alg = FQAlg_Create(uLS.size, debug)
			break
		}
	case utils.BIT:
		{
			uLS.alg = BITAlg_Create(uLS.size, debug)
			utils.DebugPrint("DUPA BIT\n", debug)
			break
		}
	case utils.TS:
		{
			uLS.alg = TSAlg_Create(uLS.size, debug)
			utils.DebugPrint("DUPA TS\n", debug)

			break
		}
	}
}

func (uLS *UpdateListSolver) Raport() (string, int) {
	return fmt.Sprint(uLS.algE), uLS.cost
}
