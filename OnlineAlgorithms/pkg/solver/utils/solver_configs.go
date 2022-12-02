package utils

import (
	pagingsolver "OnlineAlgorithms/pkg/solver/pagingSolver"
	updatelistsolver "OnlineAlgorithms/pkg/solver/updateListSolver"
	"errors"
	"strconv"
)

type SolverConfigS struct {
	ProblemType SolverTypeEnum                   `yaml:"problemType"`
	Size        int                              `yaml:"size"`
	AlgP        []pagingsolver.PagingAlg         `yaml:"algP"`
	AlgUL       []updatelistsolver.UpdateListAlg `yaml:"algUL"`
	Debug       bool                             `default:"false" yaml:"debug"`
	DoAll       bool                             `default:"false" yaml:"doAll"`
}

func (solverConfig *SolverConfigS) SolverConfig_Preprocess() {
	if solverConfig.DoAll {
		switch solverConfig.ProblemType {
		case Paging:
			solverConfig.AlgP = make([]pagingsolver.PagingAlg, 0)
			for i := 0; i < pagingsolver.NUM_OF_PAGING_ALGS; i++ {
				solverConfig.AlgP = append(solverConfig.AlgP, pagingsolver.PagingAlg(i))
			}
		case UpdateList:
			solverConfig.AlgUL = make([]updatelistsolver.UpdateListAlg, 0)
			for i := 0; i < updatelistsolver.NUM_OF_UPDATELIST_ALGS; i++ {
				solverConfig.AlgUL = append(solverConfig.AlgUL, updatelistsolver.UpdateListAlg(i))
			}

		}
	}
}

func (solverConfig *SolverConfigS) GetNumOfAlgs() int {
	switch solverConfig.ProblemType {
	case Paging:
		return len(solverConfig.AlgP)
	case UpdateList:
		return len(solverConfig.AlgUL)
	default:
		return solverConfig.GetMaxNumOfAlgs()
	}
}

func (solverConfig *SolverConfigS) GetMaxNumOfAlgs() int {
	switch solverConfig.ProblemType {
	case Paging:
		return pagingsolver.NUM_OF_PAGING_ALGS
	case UpdateList:
		return updatelistsolver.NUM_OF_UPDATELIST_ALGS
	default:
		return pagingsolver.NUM_OF_PAGING_ALGS
	}
}

func (solverConfig *SolverConfigS) Validate(size int) error {
	if solverConfig.ProblemType == Paging {
		for _, algP := range solverConfig.AlgP {
			if algP >= pagingsolver.NUM_OF_PAGING_ALGS {
				return errors.New("wrong paging algorithm identification number")
			}
		}
	}

	if solverConfig.ProblemType == UpdateList {
		for _, algUL := range solverConfig.AlgUL {
			if algUL >= updatelistsolver.NUM_OF_UPDATELIST_ALGS {
				return errors.New("wrong update list algorithm identification number")
			}
		}
	}

	if solverConfig.ProblemType == UpdateList && size >= solverConfig.Size {

		return errors.New("maximum request for" + strconv.Itoa(solverConfig.Size) + "sized update list, must be at most n-1, is" + strconv.Itoa(size))
	}

	return nil
}
