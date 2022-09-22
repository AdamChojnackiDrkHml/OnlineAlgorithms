package solver

import (
	pagingsolver "OnlineAlgorithms/internal/solver/pagingSolver"
	updatelistsolver "OnlineAlgorithms/internal/solver/updateListSolver"
)

type SolverTypeEnum int

const (
	Paging SolverTypeEnum = iota
	UpdateList
)

type GenericSolver interface {
	Serve(request int)
	Raport() (string, int)
}

func CreateSolver(conf [4]int) []GenericSolver {
	var gS []GenericSolver
	debug := conf[3] == 1
	control := conf[2] - 1
	all := false
	if conf[2] == 0 {
		control = 0
		all = true
	}
	for {
		gS = append(gS, initSolver(conf[1], control, debug, conf[0]))
		if !all || control == 4 {
			break
		}
		control++
	}
	return gS
}

func initSolver(size, alg int, debug bool, solver int) GenericSolver {
	switch SolverTypeEnum(solver) {
	case Paging:
		return pagingsolver.PagingSolver_Create(size, alg, debug)

	case UpdateList:
		return updatelistsolver.UpdateListSolver_Create(size, alg, debug)
	default:
		return nil

	}
}
