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

func CreateSolver(solverConf utils.SolverConfigS) []GenericSolver {

	var gS []GenericSolver

	if solverConf.DoAll {
		noOfAlgs := utils.GetMaxNumOfAlgs(utils.SolverTypeEnum(solverConf.ProblemType))

		for i := 0; i < noOfAlgs; i++ {
			gS = append(gS, initSolver(solverConf.Size, i, solverConf.Debug, solverConf.ProblemType))
		}

	} else {
		gS = append(gS, initSolver(solverConf.Size, int(solverConf.AlgUL), solverConf.Debug, solverConf.ProblemType))

	}
	return gS
}

func initSolver(size, alg int, debug bool, solver utils.SolverTypeEnum) GenericSolver {
	switch solver {
	case utils.Paging:
		return pagingsolver.PagingSolver_Create(size, alg, debug)

	case utils.UpdateList:
		return updatelistsolver.UpdateListSolver_Create(size, alg, debug)
	default:
		return nil

	}
}
