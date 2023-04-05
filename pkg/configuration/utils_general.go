// Package generalutils provides structs and utilities for managing
// solver, generator and environament parameters for testing and use.
package generalutils

import (
	dgconfig "OnlineAlgorithms/pkg/dataGenerator/dataGeneratorConfigs"
	svconf "OnlineAlgorithms/pkg/solver/solverConfigs"
)

// GeneralConfigS holds general testing environament configuration.
type GeneralConfigS struct {
	NoOfReq    int `yaml:"noOfReq"`
	Iterations int `yaml:"iterations"`
	Growth     int `yaml:"growth"`
	Repeats    int `yaml:"repeats"`
}

// TestConfigS holds all structures defined for single testcase.
type TestConfigS struct {
	GeneralConfig GeneralConfigS `yaml:"generalConfig"`

	SolverConfig svconf.SolverConfigS `yaml:"solverConfig"`

	GeneratorConfig dgconfig.GeneratorConfigS `yaml:"generatorConfig"`
}

// Config structure purpose is holding all defined testcases for testing purposes.
type Config struct {
	TestConfigs []TestConfigS `yaml:"test"`
}

// PreporcessTestConfig method should be called before starting test case
// in order to prepare structures and resolve doAll flags.
func PreprocessTestConfig(testConf *TestConfigS) {
	solverConfig := &testConf.SolverConfig
	generatorConfig := &testConf.GeneratorConfig

	solverConfig.Preprocess()
	generatorConfig.Preprocess()
}

// ValidateTestConfig function should be called before test start
// with a goal to find all possible flaws in configuration that could ruin test.
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
