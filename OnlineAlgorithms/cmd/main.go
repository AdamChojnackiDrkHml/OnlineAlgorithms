package main

import (
	dataGenerator "OnlineAlgorithms/internal/dataGenerator"
	"OnlineAlgorithms/internal/solver"
	"OnlineAlgorithms/internal/utils"
	"fmt"
	"os"
	"path/filepath"

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
		genConf := utils.GeneralConfigS{NoOfReq: 50, Iterations: 1, Growth: 0, Repeats: 1}
		solverConf := utils.SolverConfigS{ProblemType: 0, Size: 10, Alg: 1, Debug: true}
		generatorConf := utils.GeneratorConfigS{DistributionType: 0, Minimum: 0, Fvalue: 0.0, Maximum: 10}
		runTestForCmdArguments(&utils.Config{TestConfig: utils.TestConfigS{GeneralConfig: genConf, SolverConfig: solverConf, GeneratorConfig: generatorConf}})
	}

}

func createHeader(solverConf *utils.SolverConfigS, genConf *utils.GeneratorConfigS) string {
	header := ""

	if solverConf.ProblemType == 0 {
		header += "PAGING"
	} else {
		header += "UPDATE_LIST"
	}

	header += "\n"

	numOfAlgs := 0
	if solverConf.Alg == 0 {
		if solverConf.ProblemType == 0 {
			header += "3"
			numOfAlgs = 3
		} else {
			header += "5"
			numOfAlgs = 5
		}
	} else {
		header += "1"
		numOfAlgs = 1
	}

	header += "\n"
	fmt.Println(numOfAlgs)
	if numOfAlgs == 1 {
		if solverConf.ProblemType == 0 {
			header += fmt.Sprintf("%s", utils.PagingAlg(solverConf.Alg))
		} else {
			header += fmt.Sprintf("%s", utils.UpdateListAlg(solverConf.Alg))
		}
	} else {
		if solverConf.ProblemType == 0 {
			for i := 0; i < numOfAlgs; i++ {
				header += fmt.Sprintf("%s ", utils.PagingAlg(i))
			}
		} else {
			for i := 0; i < numOfAlgs; i++ {
				header += fmt.Sprintf("%s ", utils.UpdateListAlg(i))
			}
		}
	}

	header += "\n"

	header += fmt.Sprintf("%d", genConf.DistributionType)
	header += "\n"

	if genConf.DistributionType == 0 {
		for i := 0; i < 6; i++ {
			header += fmt.Sprintf("%s ", utils.GeneratorTypeEnum(i))
		}
	} else {
		header += fmt.Sprintf("%s", utils.GeneratorTypeEnum(genConf.DistributionType))
	}

	header += "\n"

	header += fmt.Sprintf("%d ", solverConf.Size)

	header += "\n"

	return header
}

func saveResToFile(f *os.File, ress []int, noOfReq int) {
	fmt.Fprint(f, noOfReq)
	fmt.Fprint(f, " ")
	for _, res := range ress {
		fmt.Fprint(f, res)
		fmt.Fprint(f, " ")
	}
	fmt.Fprintln(f)
}

func runTestForCmdArguments(conf *utils.Config) {

	solvConf := conf.TestConfig.SolverConfig
	generConf := conf.TestConfig.GeneratorConfig
	genConf := conf.TestConfig.GeneralConfig
	for iteration := 0; iteration < genConf.Iterations; iteration++ {
		for repeat := 0; repeat < genConf.Repeats; repeat++ {
			pSS := solver.CreateSolver(solvConf, 1)
			dG := dataGenerator.CreateDataGenerator(generConf)

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

	resName := "data/res/" + filepath.Base(os.Args[2])
	fmt.Println(resName)
	f, err2 := os.OpenFile(resName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err2 != nil {
		utils.ExitWithError(err2.Error())
	}
	defer f.Close()

	solvConf := conf.TestConfig.SolverConfig
	generConf := conf.TestConfig.GeneratorConfig
	genConf := conf.TestConfig.GeneralConfig

	header := createHeader(&solvConf, &generConf)

	fmt.Fprint(f, header)
	noOfAlgs := 1
	if solvConf.Alg == 0 {
		if solvConf.ProblemType == 0 {
			noOfAlgs = 4
		} else {
			noOfAlgs = 5
		}
	}

	var name string
	var score int
	for iteration := 0; iteration < genConf.Iterations; iteration++ {
		ress := make([]int, noOfAlgs)
		names := make([]string, noOfAlgs)
		for repeat := 0; repeat < genConf.Repeats; repeat++ {
			pSS := solver.CreateSolver(solvConf, noOfAlgs)
			dG := dataGenerator.CreateDataGenerator(generConf)

			for request := 0; request < genConf.NoOfReq; request++ {
				for _, pS := range pSS {
					pS.Serve(dG.GetRequest())
				}
			}

			for i, pS := range pSS {
				name, score = pS.Raport()
				names[i] = name
				ress[i] += int(float64(score) / float64(genConf.Repeats))
			}

		}
		saveResToFile(f, ress, genConf.NoOfReq)
		fmt.Println("aa")
		genConf.NoOfReq += genConf.Growth
	}
}
