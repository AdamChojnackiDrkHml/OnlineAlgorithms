package solver

import (
	pagingsolver "OnlineAlgorithms/internal/solver/pagingSolver"
	updatelistsolver "OnlineAlgorithms/internal/solver/updateListSolver"
	genUtils "OnlineAlgorithms/internal/utils/generalUtils"
)

type GenericSolver interface {
	Serve(request int)
	Raport() (string, int)
}

func CreateSolver(solverConf genUtils.SolverConfigS) []GenericSolver {

	var gS []GenericSolver

	switch solverConf.ProblemType {
	case genUtils.Paging:
		for _, algP := range solverConf.AlgP {
			gS = append(gS, pagingsolver.PagingSolver_Create(solverConf.Size, algP, solverConf.Debug))
		}
	case genUtils.UpdateList:
		for _, algUL := range solverConf.AlgUL {
			gS = append(gS, updatelistsolver.UpdateListSolver_Create(solverConf.Size, algUL, solverConf.Debug))
		}
	}

	return gS
}
