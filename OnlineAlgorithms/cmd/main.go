package main

import (
	dataGenerator "OnlineAlgorithms/internal/dataGenerator"
	"OnlineAlgorithms/internal/solver"
	genUtils "OnlineAlgorithms/internal/utils/generalUtils"
	ioutils "OnlineAlgorithms/internal/utils/ioUtils"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/exp/slices"
)

func main() {

	if ind := slices.Index(os.Args, "-f"); ind != -1 {
		config := ioutils.ReadYamlForConfig(os.Args[ind+1])
		runTestWithParametersFromFile(config)
	} else if ind := slices.Index(os.Args, "-p"); ind != -1 {
		runTestForCmdArguments(ioutils.ParseCmd(os.Args[ind+1:]))
	} else { //hand debug case
		genConf := genUtils.GeneralConfigS{NoOfReq: 500, Iterations: 20, Growth: 500, Repeats: 20}
		solverConf := genUtils.SolverConfigS{ProblemType: 0, Size: 30, AlgP: []genUtils.PagingAlg{1}, Debug: false, DoAll: true}
		generatorConf := genUtils.GeneratorConfigS{DistributionType: []genUtils.GeneratorTypeEnum{0}, Minimum: 0, Maximum: 150, DoAll: false, FvalueGeo: 0.2, FvaluePoiss: 0.3}
		runTestWithParametersFromFile(&genUtils.Config{TestConfigs: []genUtils.TestConfigS{{GeneralConfig: genConf, SolverConfig: solverConf, GeneratorConfig: generatorConf}}})
	}

}

func runTestForCmdArguments(conf genUtils.Config) {

	testConf := conf.TestConfigs[0]

	solvConf := testConf.SolverConfig
	generConf := testConf.GeneratorConfig
	genConf := testConf.GeneralConfig
	for iteration := 0; iteration < genConf.Iterations; iteration++ {
		for repeat := 0; repeat < genConf.Repeats; repeat++ {
			pSS := solver.CreateSolver(solvConf)

			dG := dataGenerator.CreateDataGenerator(generConf)[0]

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

func runTestWithParametersFromFile(conf *genUtils.Config) {

	for i, testConf := range conf.TestConfigs {

		if err := genUtils.ValidateTestConfig(testConf); err != nil {
			fmt.Fprintln(os.Stderr, "Testcase ", fmt.Sprint(i), " error: ", err.Error())
			continue
		}

		genUtils.PreprocessTestConfig(&testConf)

		fileName := filepath.Base(os.Args[2])
		resFilename := "data/res/" + "results_" + strings.TrimSuffix(fileName, filepath.Ext(fileName)) + fmt.Sprint(i)

		f := ioutils.OpenFile(resFilename)

		solvConf := testConf.SolverConfig
		generConf := testConf.GeneratorConfig
		genConf := testConf.GeneralConfig

		ioutils.CreateAndWriteHeader(f, &solvConf, &generConf)

		noOfAlgs := genUtils.GetNumOfAlgs(solvConf)
		noOfDistros := genUtils.GetNumOfDistributions(generConf)

		var name string
		var score int

		dGS := dataGenerator.CreateDataGenerator(generConf)

		for iteration := 0; iteration < genConf.Iterations; iteration++ {

			ress := make([]int, noOfAlgs*noOfDistros)
			names := make([]string, noOfAlgs*noOfDistros)
			for repeat := 0; repeat < genConf.Repeats; repeat++ {
				problemSolversForGenerators := make([][]solver.GenericSolver, noOfDistros)
				for i := range dGS {
					problemSolversForGenerators[i] = solver.CreateSolver(solvConf)
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
