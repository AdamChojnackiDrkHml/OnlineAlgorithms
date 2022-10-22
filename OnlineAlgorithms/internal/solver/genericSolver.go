package solver

import (
	pagingsolver "OnlineAlgorithms/internal/solver/pagingSolver"
	updatelistsolver "OnlineAlgorithms/internal/solver/updateListSolver"
	"OnlineAlgorithms/internal/utils"
)

type GenericSolver interface {
	Serve(request int)
	Raport() (string, int)
}

func CreateSolver(solverConf utils.SolverConfigS, noOfAlgs int) []GenericSolver {
	var gS []GenericSolver

	control := solverConf.Alg - 1
	all := false
	if solverConf.Alg == 0 {
		control = 0
		all = true
	}
	for {
		gS = append(gS, initSolver(solverConf.Size, control, solverConf.Debug, solverConf.ProblemType))
		if !all || control == noOfAlgs-1 {
			break
		}
		control++
	}
	return gS
}

func initSolver(size, alg int, debug bool, solver int) GenericSolver {
	switch utils.SolverTypeEnum(solver) {
	case utils.Paging:
		return pagingsolver.PagingSolver_Create(size, alg, debug)

	case utils.UpdateList:
		return updatelistsolver.UpdateListSolver_Create(size, alg, debug)
	default:
		return nil

	}
}
