// Package testcontroler provides functions to start and execute tests
package testcontroler

import (
	conf "OnlineAlgorithms/pkg/configuration"
	dg "OnlineAlgorithms/pkg/dataGenerator"
	"OnlineAlgorithms/pkg/solver"
	ioutils "OnlineAlgorithms/pkg/testFramework/ioUtils"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// RunTestForCmdArguments parses passed string slice containing configuration
// and then executes test for given config.
// It prints results to stdout.
// Created for running cmd configuration, but will work for any slice of strings.
// Is capable of running only single testcase!
// Legacy, use yaml configuration.
func RunTestForCmdArguments(cmdArgs []string) {
	conf := ioutils.ParseCmdForConfig(cmdArgs)

	testConf := conf.TestConfigs[0]

	solvConf := testConf.SolverConfig
	generConf := testConf.GeneratorConfig
	genConf := testConf.GeneralConfig
	for iteration := 0; iteration < genConf.Iterations; iteration++ {
		for repeat := 0; repeat < genConf.Repeats; repeat++ {
			pSS := solver.CreateSolversFromConfig(solvConf)

			dG := dg.CreateDataGenerator(generConf)[0]

			for request := 0; request < genConf.NoOfReq; request++ {
				for _, pS := range pSS {
					pS.Serve(dG.GetRequest())
				}
			}
			ress := 0

			for _, pS := range pSS {
				name, score := pS.Raport()
				ress += int(float64(score) / float64(genConf.Repeats))
				fmt.Println(name, ress, genConf.NoOfReq)
			}

		}
	}
}

// RunTestForConfig takes Config structure and executes tests defined in it.
// Second argument is results filename.
// Results are saved to files filename$(testcase number).
func RunTestForConfig(config *conf.Config, resFilename string) {

	for i, testConf := range config.TestConfigs {

		if err := conf.ValidateTestConfig(testConf); err != nil {
			fmt.Fprintln(os.Stderr, "Testcase ", fmt.Sprint(i), " error: ", err.Error())
			continue
		}

		conf.PreprocessTestConfig(&testConf)

		f := ioutils.CreateAndOpenResFile(resFilename + fmt.Sprint(i))

		solvConf := testConf.SolverConfig
		generConf := testConf.GeneratorConfig
		genConf := testConf.GeneralConfig

		ioutils.CreateAndWriteHeader(f, &solvConf, &generConf)

		noOfAlgs := solvConf.GetNumOfAlgs()
		noOfDistros := generConf.GetNumOfDistributions()

		var name string
		var score int

		dGS := dg.CreateDataGenerator(generConf)

		for iteration := 0; iteration < genConf.Iterations; iteration++ {

			ress := make([]int, noOfAlgs*noOfDistros)
			names := make([]string, noOfAlgs*noOfDistros)
			for repeat := 0; repeat < genConf.Repeats; repeat++ {
				problemSolversForGenerators := make([][]solver.GenericSolver, noOfDistros)
				for i := range dGS {
					problemSolversForGenerators[i] = solver.CreateSolversFromConfig(solvConf)
				}
				for requestIterator := 0; requestIterator < genConf.NoOfReq; requestIterator++ {
					for i, generator := range dGS {
						solversForGenerator := problemSolversForGenerators[i]
						request := generator.GetRequest()
						for _, problemSolver := range solversForGenerator {
							problemSolver.Serve(request)
						}
					}
				}

				resultsIterator := 0
				for i := range dGS {
					solversForGenerator := problemSolversForGenerators[i]
					for _, problemSolver := range solversForGenerator {
						name, score = problemSolver.Raport()
						names[resultsIterator] = name
						ress[resultsIterator] += int(float64(score) / float64(genConf.Repeats))

						resultsIterator++
					}
				}
			}
			ioutils.SaveResToFile(f, ress, genConf.NoOfReq)

			genConf.NoOfReq += genConf.Growth

		}

		f.Close()

	}
}

// RunTestForFileConfig takes as argument path to yaml file.
// This method is wrapper for yaml parsing and running RunTestForConfig().
// "results_"+filename is used as results filename.
func RunTestForFileConfig(path string) {
	config := ioutils.ReadYamlForConfig(path)
	fileName := filepath.Base(path)
	resFilename := "data/res/" + "results_" + strings.TrimSuffix(fileName, filepath.Ext(fileName))
	RunTestForConfig(config, resFilename)
}
