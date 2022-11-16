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

	switch solverConf.ProblemType {
	case utils.Paging:
		for _, algP := range solverConf.AlgP {
			gS = append(gS, pagingsolver.PagingSolver_Create(solverConf.Size, algP, solverConf.Debug))
		}
	case utils.UpdateList:
		for _, algUL := range solverConf.AlgUL {
			gS = append(gS, updatelistsolver.UpdateListSolver_Create(solverConf.Size, algUL, solverConf.Debug))
		}
	}

	return gS
}
