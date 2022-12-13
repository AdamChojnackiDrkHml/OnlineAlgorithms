// Package solver defines general functionalities for online algorithms solvers.
// Provides interface for new solvers and general constructors.
package solver

import (
	pagingsolver "OnlineAlgorithms/pkg/solver/pagingSolver"
	psalgs "OnlineAlgorithms/pkg/solver/pagingSolver/pagingSolverAlgs"
	svconf "OnlineAlgorithms/pkg/solver/solverConfigs"
	updatelistsolver "OnlineAlgorithms/pkg/solver/updateListSolver"
	ulsalgs "OnlineAlgorithms/pkg/solver/updateListSolver/updateListSolverAlgs"
)

// GenericSolver provides front for any solver implementing it.
type GenericSolver interface {
	Serve(request int)
	Raport() (string, int)
}

// CreateSolversFromConfig function takes SolverConfigS struct
// and returns slice of GenericSolver based on its contents.
// Remember this function supports only one problem at a time.
func CreateSolversFromConfig(solverConf svconf.SolverConfigS) []GenericSolver {

	var gS []GenericSolver
	debug := solverConf.Debug
	size := solverConf.Size

	switch solverConf.ProblemType {
	case svconf.Paging:
		for _, algP := range solverConf.AlgP {
			gS = append(gS, CreateSinglePagingSolver(size, algP, debug))
		}
	case svconf.UpdateList:
		for _, algUL := range solverConf.AlgUL {
			gS = append(gS, CreateSingleUpdateListSolver(size, algUL, debug))
		}
	}

	return gS
}

// CreateSinglePagingSolver takes configuration for single paging solver
// and returns it.
func CreateSinglePagingSolver(size int, alg psalgs.PagingAlg, debug bool) GenericSolver {
	return pagingsolver.PagingSolver_Create(size, alg, debug)
}

// CreateSingleUpdateListSolver takes configuration for single paging solver
// and returns it.
func CreateSingleUpdateListSolver(size int, alg ulsalgs.UpdateListAlg, debug bool) GenericSolver {
	return updatelistsolver.UpdateListSolver_Create(size, alg, debug)
}
