package solverutils

import (
	psAlgs "OnlineAlgorithms/pkg/solver/pagingSolver/pagingSolverAlgs"
	ulsAlgs "OnlineAlgorithms/pkg/solver/updateListSolver/updateListSolverAlgs"
	"errors"
	"strconv"
)

type SolverConfigS struct {
	ProblemType SolverTypeEnum          `yaml:"problemType"`
	Size        int                     `yaml:"size"`
	AlgP        []psAlgs.PagingAlg      `yaml:"algP"`
	AlgUL       []ulsAlgs.UpdateListAlg `yaml:"algUL"`
	Debug       bool                    `default:"false" yaml:"debug"`
	DoAll       bool                    `default:"false" yaml:"doAll"`
}

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
		return psAlgs.NUM_OF_PAGING_ALGS
	case UpdateList:
		return ulsAlgs.NUM_OF_UPDATELIST_ALGS
	default:
		return psAlgs.NUM_OF_PAGING_ALGS
	}
}

func (solverConfig *SolverConfigS) Validate(size int) error {
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

	if solverConfig.ProblemType == UpdateList && size >= solverConfig.Size {

		return errors.New("maximum request for" + strconv.Itoa(solverConfig.Size) + "sized update list, must be at most n-1, is" + strconv.Itoa(size))
	}

	return nil
}
