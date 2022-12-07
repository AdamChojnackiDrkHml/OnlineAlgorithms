package generalutils

import (
	dgutils "OnlineAlgorithms/pkg/dataGenerator/dataGeneratorUtils"
	solverutils "OnlineAlgorithms/pkg/solver/solverUtils"
)

// //////////////////////////////
// TEST CONFIG HOLDING STRUCTURES
// //////////////////////////////
type GeneralConfigS struct {
	NoOfReq    int `yaml:"noOfReq"`
	Iterations int `yaml:"iterations"`
	Growth     int `yaml:"growth"`
	Repeats    int `yaml:"repeats"`
}

type TestConfigS struct {
	GeneralConfig GeneralConfigS `yaml:"generalConfig"`

	SolverConfig solverutils.SolverConfigS `yaml:"solverConfig"`

	GeneratorConfig dgutils.GeneratorConfigS `yaml:"generatorConfig"`
}

type Config struct {
	TestConfigs []TestConfigS `yaml:"test"`
}

// /////////////////////////////////////
// GENERAL PROGRAM CONSTANTS AND GETTERS
// /////////////////////////////////////

// ////////////////////////////////////////////
// UTILS FOR TEST CONFIG CHECKING AND PREPARING
// ////////////////////////////////////////////
func PreprocessTestConfig(testConf *TestConfigS) {
	solverConfig := &testConf.SolverConfig
	generatorConfig := &testConf.GeneratorConfig

	solverConfig.SolverConfig_Preprocess()
	generatorConfig.Preprocess()
}

func ValidateTestConfig(testConf TestConfigS) error {
	solverConfig := testConf.SolverConfig
	generatorConfig := testConf.GeneratorConfig

	errGenerator := generatorConfig.Validate()

	if errGenerator != nil {
		return errGenerator
	}

	errSolver := solverConfig.Validate(generatorConfig.Maximum)

	if errSolver != nil {
		return errSolver
	}

	return nil
}
