package solver

import (
	pagingsolver "OnlineAlgorithms/internal/solver/pagingSolver"
	updatelistsolver "OnlineAlgorithms/internal/solver/updateListSolver"
	"OnlineAlgorithms/internal/utils"
	"errors"
)

type GenericSolver interface {
	Serve(request int)
	Raport() (string, int)
}

func CreateSolver(solverConf utils.SolverConfigS) ([]GenericSolver, error) {

	if !validateConfigs(solverConf) {
		return nil, errors.New("invalid algorith number for given solver")
	}

	var gS []GenericSolver

	if solverConf.DoAll {
		noOfAlgs := utils.GetMaxNumOfAlgs(utils.SolverTypeEnum(solverConf.ProblemType))

		for i := 0; i < noOfAlgs; i++ {
			gS = append(gS, initSolver(solverConf.Size, i, solverConf.Debug, solverConf.ProblemType))
		}

	} else {
		gS = append(gS, initSolver(solverConf.Size, solverConf.Alg, solverConf.Debug, solverConf.ProblemType))

	}
	return gS, nil
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

func validateConfigs(solverConf utils.SolverConfigS) bool {
	if utils.SolverTypeEnum(solverConf.ProblemType) == utils.Paging && (solverConf.Alg >= utils.NUM_OF_PAGING_ALGS) {
		return false
	}

	if utils.SolverTypeEnum(solverConf.ProblemType) == utils.UpdateList && (solverConf.Alg >= utils.NUM_OF_UPDATELIST_ALGS) {
		return false
	}

	return true
}
