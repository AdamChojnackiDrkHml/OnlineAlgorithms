package main

import (
	dataGenerator "OnlineAlgorithms/internal/dataGenerator"
	"OnlineAlgorithms/internal/solver"
	"OnlineAlgorithms/internal/utils"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/exp/slices"
)

func main() {

	if ind := slices.Index(os.Args, "-f"); ind != -1 {
		fmt.Println("aa")
		config, err := utils.ParseYaml(os.Args[ind+1])
		if err != nil {
			utils.ExitWithError(err.Error())
		}
		runTestWithParametersFromFile(config)
	} else if ind := slices.Index(os.Args, "-p"); ind != -1 {
		runTestForCmdArguments(utils.ParseCmd(os.Args[ind+1:]))
	} else { //hand debug case
		genConf := utils.GeneralConfigS{NoOfReq: 50, Iterations: 1, Growth: 500, Repeats: 1}
		solverConf := utils.SolverConfigS{ProblemType: 0, Size: 10, AlgP: 1, Debug: true}
		generatorConf := utils.GeneratorConfigS{DistributionType: 0, Minimum: 0, Maximum: 10, DoAll: true, FvalueGeo: 0.2, FvaluePoiss: 0.3}
		runTestWithParametersFromFile(&utils.Config{TestConfigs: []utils.TestConfigS{{GeneralConfig: genConf, SolverConfig: solverConf, GeneratorConfig: generatorConf}}})
	}

}

func runTestForCmdArguments(testConf utils.TestConfigS) {

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

func runTestWithParametersFromFile(conf *utils.Config) {

	for i, testConf := range conf.TestConfigs {

		if err := utils.ValidateTestConfig(testConf); err != nil {
			fmt.Fprintln(os.Stderr, "Testcase ", fmt.Sprint(i), " error: ", err.Error())
			continue
		}

		fileName := filepath.Base(os.Args[2])
		resFilename := "data/res/" + strings.TrimSuffix(fileName, filepath.Ext(fileName)) + fmt.Sprint(i)

		f := utils.OpenFile(resFilename)

		solvConf := testConf.SolverConfig
		generConf := testConf.GeneratorConfig
		genConf := testConf.GeneralConfig

		utils.CreateHeader(f, &solvConf, &generConf)

		noOfAlgs := utils.GetNumOfAlgs(solvConf.ProblemType, solvConf.DoAll)
		noOfDistros := utils.GetNumOfDistributions(generConf)

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
				for request := 0; request < genConf.NoOfReq; request++ {
					for i, generator := range dGS {
						solversForGenerator := problemSolversForGenerators[i]
						request := generator.GetRequest()
						for _, problemSolver := range solversForGenerator {
							problemSolver.Serve(request)
						}
					}
				}
				for i := range dGS {
					solversForGenerator := problemSolversForGenerators[i]
					for j, problemSolver := range solversForGenerator {
						name, score = problemSolver.Raport()
						names[i+j*noOfAlgs] = name
						ress[i+j*noOfAlgs] += int(float64(score) / float64(genConf.Repeats))
					}
				}
			}
			utils.SaveResToFile(f, ress, genConf.NoOfReq)

			genConf.NoOfReq += genConf.Growth

		}

		f.Close()

	}
}
