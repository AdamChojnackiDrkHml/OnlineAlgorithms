// Package solverconfigs defines structures and functionalities
// for usage of solver configuration.
// It also defines existing solvers.
package solverconfigs

import (
	psAlgs "OnlineAlgorithms/pkg/solver/pagingSolver/pagingSolverAlgs"
	ulsAlgs "OnlineAlgorithms/pkg/solver/updateListSolver/updateListSolverAlgs"
	"errors"
	"strconv"
)

// SolverConfigS holds solver configuration.
// By default Debug and DoAll(algorithms) flags are set to false.
// Note that while struct can hold algorithms for multiple problems
// it can only be used for one problem at a time.
type SolverConfigS struct {
	ProblemType SolverTypeEnum          `yaml:"problemType"`
	Size        int                     `yaml:"size"`
	AlgP        []psAlgs.PagingAlg      `yaml:"algP"`
	AlgUL       []ulsAlgs.UpdateListAlg `yaml:"algUL"`
	Debug       bool                    `default:"false" yaml:"debug"`
	DoAll       bool                    `default:"false" yaml:"doAll"`
}

// Preprocess method should be called when using DoAll flag
// in order to fill proper algorithm slice.
func (solverConfig *SolverConfigS) Preprocess() {
	if solverConfig.DoAll {
		switch solverConfig.ProblemType {
		case Paging:
			solverConfig.AlgP = make([]psAlgs.PagingAlg, 0)
			for i := 0; i < psAlgs.NUM_OF_PAGING_ALGS; i++ {
				solverConfig.AlgP = append(solverConfig.AlgP, psAlgs.PagingAlg(i))
			}
		case UpdateList:
			solverConfig.AlgUL = make([]ulsAlgs.UpdateListAlg, 0)
			for i := 0; i < ulsAlgs.NUM_OF_UPDATELIST_ALGS; i++ {
				solverConfig.AlgUL = append(solverConfig.AlgUL, ulsAlgs.UpdateListAlg(i))
			}

		}
	}
}

// GetNumOfAlgs returns number of algorithms in configuration for set problem.
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

// GetMaxNumOfAlgs returns number of defined algorithms for set problem.
func (solverConfig *SolverConfigS) GetMaxNumOfAlgs() int {
	switch solverConfig.ProblemType {
	case Paging:
		return psAlgs.NUM_OF_PAGING_ALGS
	case UpdateList:
		return ulsAlgs.NUM_OF_UPDATELIST_ALGS
	default:
		return psAlgs.NUM_OF_PAGING_ALGS
	}
}

// Validate takes highest possible request.
// This function checks if configuration is correct
// and will not cause errors in runtime.
// Returns error if finds incorrect config.
// Else returns nil.
func (solverConfig *SolverConfigS) Validate(maxRequestGen int) error {
	if solverConfig.ProblemType == Paging {
		for _, algP := range solverConfig.AlgP {
			if algP >= psAlgs.NUM_OF_PAGING_ALGS {
				return errors.New("wrong paging algorithm identification number")
			}
		}
	}

	if solverConfig.ProblemType == UpdateList {
		for _, algUL := range solverConfig.AlgUL {
			if algUL >= ulsAlgs.NUM_OF_UPDATELIST_ALGS {
				return errors.New("wrong update list algorithm identification number")
			}
		}
	}

	if solverConfig.ProblemType == UpdateList && maxRequestGen >= solverConfig.Size {

		return errors.New("maximum request for" + strconv.Itoa(solverConfig.Size) + "sized update list, must be at most n-1, is" + strconv.Itoa(maxRequestGen))
	}

	return nil
}
