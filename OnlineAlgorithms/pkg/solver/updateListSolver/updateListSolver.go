package updatelistsolver

import (
	ulalgs "OnlineAlgorithms/pkg/solver/updateListSolver/updateListSolverAlgs"
	"fmt"
)

type UpdateListSolvingAlg interface {
	UpdateList(request int) int
}

type UpdateListSolver struct {
	size int
	cost int
	alg  UpdateListSolvingAlg
	algE ulalgs.UpdateListAlg
}

func UpdateListSolver_Create(size int, algUL ulalgs.UpdateListAlg, debug bool) *UpdateListSolver {
	uLS := &UpdateListSolver{size: size, cost: 0, algE: ulalgs.UpdateListAlg(algUL)}
	uLS.createSolvingAlg(algUL, debug)
	return uLS
}

func (uLS *UpdateListSolver) Serve(request int) {
	uLS.cost += uLS.alg.UpdateList(request)
}

func (uLS *UpdateListSolver) createSolvingAlg(algUL ulalgs.UpdateListAlg, debug bool) {
	switch algUL {
	case ulalgs.MTF:
		{
			uLS.alg = ulalgs.MTFAlg_Create(uLS.size, debug)
		}
	case ulalgs.TRANS:
		{
			uLS.alg = ulalgs.TransAlg_Create(uLS.size, debug)
		}
	case ulalgs.FC:
		{
			uLS.alg = ulalgs.FCAlg_Create(uLS.size, debug)
		}
	case ulalgs.BIT:
		{
			uLS.alg = ulalgs.BITAlg_Create(uLS.size, debug)
		}
	case ulalgs.TS:
		{
			uLS.alg = ulalgs.TSAlg_Create(uLS.size, debug)
		}
	case ulalgs.Combination:
		{
			uLS.alg = ulalgs.CombinationAlg_Create(uLS.size, debug)
		}
	}
}

func (uLS *UpdateListSolver) Raport() (string, int) {
	return fmt.Sprint(uLS.algE), uLS.cost
}
