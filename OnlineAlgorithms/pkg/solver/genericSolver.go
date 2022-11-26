package solver

import (
	pagingsolver "OnlineAlgorithms/pkg/solver/pagingSolver"
	updatelistsolver "OnlineAlgorithms/pkg/solver/updateListSolver"
	genUtils "OnlineAlgorithms/pkg/utils/generalUtils"
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
