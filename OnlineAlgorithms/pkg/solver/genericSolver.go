package solver

import (
	generalUtils "OnlineAlgorithms/pkg/generalUtils"
	pagingsolver "OnlineAlgorithms/pkg/solver/pagingSolver"
	updatelistsolver "OnlineAlgorithms/pkg/solver/updateListSolver"
	solverutils "OnlineAlgorithms/pkg/solver/utils"
)

type GenericSolver interface {
	Serve(request int)
	Raport() (string, int)
}

func CreateSolversFromConfig(solverConf generalUtils.SolverConfigS) []GenericSolver {

	var gS []GenericSolver
	debug := solverConf.Debug
	size := solverConf.Size

	switch solverConf.ProblemType {
	case solverutils.Paging:
		for _, algP := range solverConf.AlgP {
			gS = append(gS, CreateSinglePagingSolver(size, algP, debug))
		}
	case solverutils.UpdateList:
		for _, algUL := range solverConf.AlgUL {
			gS = append(gS, CreateSingleUpdateListSolver(size, algUL, debug))
		}
	}

	return gS
}

func CreateSinglePagingSolver(size int, alg pagingsolver.PagingAlg, debug bool) GenericSolver {
	return pagingsolver.PagingSolver_Create(size, alg, debug)
}

func CreateSingleUpdateListSolver(size int, alg updatelistsolver.UpdateListAlg, debug bool) GenericSolver {
	return updatelistsolver.UpdateListSolver_Create(size, alg, debug)
}
