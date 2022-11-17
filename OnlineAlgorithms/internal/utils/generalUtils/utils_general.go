package generalutils

import (
	"errors"
)

// ///////////
// SOLVER ENUM
// ///////////
type SolverTypeEnum int

const (
	Paging SolverTypeEnum = iota
	UpdateList
)

func (e SolverTypeEnum) String() string {
	switch e {
	case Paging:
		return "Paging"
	case UpdateList:
		return "UpdateList"
	default:
		return "NULL"
	}

}

// ////////////////
// UPDATE LIST ENUM
// ////////////////
type UpdateListAlg int

const (
	MTF UpdateListAlg = iota
	TRANS
	FQ
	BIT
	TS
)

func (e UpdateListAlg) String() string {
	switch e {
	case MTF:
		return "MTF"
	case TRANS:
		return "TRANS"
	case FQ:
		return "FQ"
	case BIT:
		return "BIT"
	case TS:
		return "TS"
	default:
		return "NULL"
	}
}

// ///////////
// PAGING ENUM
// ///////////
type PagingAlg int

const (
	LRU PagingAlg = iota
	FIFO
	LFU
	MARK
	MARK2
	RM
)

func (e PagingAlg) String() string {
	switch e {
	case LRU:
		return "LRU"
	case FIFO:
		return "FIFO"
	case LFU:
		return "LFU"
	case MARK:
		return "MARK"
	case MARK2:
		return "MARK2"
	case RM:
		return "RM"
	default:
		return "NULL"
	}
}

// ///////////////
// GENERATOR ENUM
// ///////////////
type GeneratorTypeEnum int

const (
	Uni GeneratorTypeEnum = iota
	Geo
	Pois
	Hrm
	Dhr
)

func (e GeneratorTypeEnum) String() string {
	switch e {
	case Uni:
		return "Uni"
	case Geo:
		return "Geo"
	case Pois:
		return "Pois"
	case Hrm:
		return "Hrm"
	case Dhr:
		return "Dhr"
	default:
		return "NULL"
	}
}

// //////////////////////////////
// TEST CONFIG HOLDING STRUCTURES
// //////////////////////////////
type GeneralConfigS struct {
	NoOfReq    int `yaml:"noOfReq"`
	Iterations int `yaml:"iterations"`
	Growth     int `yaml:"growth"`
	Repeats    int `yaml:"repeats"`
}
type SolverConfigS struct {
	ProblemType SolverTypeEnum  `yaml:"problemType"`
	Size        int             `yaml:"size"`
	AlgP        []PagingAlg     `yaml:"algP"`
	AlgUL       []UpdateListAlg `yaml:"algUL"`
	Debug       bool            `default:"false" yaml:"debug"`
	DoAll       bool            `default:"false" yaml:"doAll"`
}
type GeneratorConfigS struct {
	DistributionType []GeneratorTypeEnum `yaml:"distributionType"`
	Minimum          int                 `yaml:"minimum"`
	FvalueGeo        float64             `yaml:"fvalueGeo"`
	FvaluePoiss      float64             `yaml:"fvaluePoiss"`
	Maximum          int                 `yaml:"maximum"`
	DoAll            bool                `default:"false" yaml:"doAll"`
}

type TestConfigS struct {
	GeneralConfig GeneralConfigS `yaml:"generalConfig"`

	SolverConfig SolverConfigS `yaml:"solverConfig"`

	GeneratorConfig GeneratorConfigS `yaml:"generatorConfig"`
}

type Config struct {
	TestConfigs []TestConfigS `yaml:"test"`
}

// /////////////////////////////////////
// GENERAL PROGRAM CONSTANTS AND GETTERS
// /////////////////////////////////////
const (
	NUM_OF_PAGING_ALGS     = 6
	NUM_OF_UPDATELIST_ALGS = 5
	NUM_OF_DISTRIBUTIONS   = 5
)

func GetNumOfAlgs(solverConfig SolverConfigS) int {
	switch solverConfig.ProblemType {
	case Paging:
		return len(solverConfig.AlgP)
	case UpdateList:
		return len(solverConfig.AlgUL)
	default:
		return GetMaxNumOfAlgs(solverConfig.ProblemType)
	}
}

func GetNumOfDistributions(generatorConfigs GeneratorConfigS) int {
	return len(generatorConfigs.DistributionType)
}

func GetMaxNumOfAlgs(solverType SolverTypeEnum) int {
	switch solverType {
	case Paging:
		return NUM_OF_PAGING_ALGS
	case UpdateList:
		return NUM_OF_UPDATELIST_ALGS
	default:
		return NUM_OF_PAGING_ALGS
	}
}

// ////////////////////////////////////////////
// UTILS FOR TEST CONFIG CHECKING AND PREPARING
// ////////////////////////////////////////////
func PreprocessTestConfig(testConf *TestConfigS) {
	solverConfig := &testConf.SolverConfig
	generatorConfig := &testConf.GeneratorConfig

	if solverConfig.DoAll {
		switch solverConfig.ProblemType {
		case Paging:
			solverConfig.AlgP = make([]PagingAlg, 0)
			for i := 0; i < NUM_OF_PAGING_ALGS; i++ {
				solverConfig.AlgP = append(solverConfig.AlgP, PagingAlg(i))
			}
		case UpdateList:
			solverConfig.AlgUL = make([]UpdateListAlg, 0)
			for i := 0; i < NUM_OF_UPDATELIST_ALGS; i++ {
				solverConfig.AlgUL = append(solverConfig.AlgUL, UpdateListAlg(i))
			}

		}
	}

	if generatorConfig.DoAll {
		generatorConfig.DistributionType = make([]GeneratorTypeEnum, 0)

		for i := 0; i < NUM_OF_DISTRIBUTIONS; i++ {
			generatorConfig.DistributionType = append(generatorConfig.DistributionType, GeneratorTypeEnum(i))
		}
	}

}

func ValidateTestConfig(testConf TestConfigS) error {
	solverConfig := testConf.SolverConfig
	generatorConfig := testConf.GeneratorConfig

	// if generatorConfig.DoAll && solverConfig.DoAll {
	// 	return errors.New("cannot do both do alls")
	// }

	for _, distribution := range generatorConfig.DistributionType {
		if distribution >= NUM_OF_DISTRIBUTIONS {
			return errors.New("wrong distribution identification number")
		}
	}

	if solverConfig.ProblemType == Paging {
		for _, algP := range solverConfig.AlgP {
			if algP >= NUM_OF_PAGING_ALGS {
				return errors.New("wrong paging algorithm identification number")
			}
		}
	}

	if solverConfig.ProblemType == UpdateList {
		for _, algUL := range solverConfig.AlgUL {
			if algUL >= NUM_OF_UPDATELIST_ALGS {
				return errors.New("wrong paging algorithm identification number")
			}
		}
	}

	if generatorConfig.Maximum >= solverConfig.Size && solverConfig.ProblemType == UpdateList {
		return errors.New("maximum request for n sized update list, must be at most n-1")
	}

	return nil
}
