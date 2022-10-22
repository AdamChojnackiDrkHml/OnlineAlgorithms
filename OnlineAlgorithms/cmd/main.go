package main

import (
	dataGenerator "OnlineAlgorithms/internal/dataGenerator"
	"OnlineAlgorithms/internal/solver"
	pagingsolver "OnlineAlgorithms/internal/solver/pagingSolver"
	updatelistsolver "OnlineAlgorithms/internal/solver/updateListSolver"
	"OnlineAlgorithms/internal/utils"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func main() {
	conf, err := utils.ParseYaml("data/configs/generic_structure.yml")

	if err != nil {
		exitWithError(err.Error())
	}

	fmt.Println(conf)
	if ind := slices.Index(os.Args, "-f"); ind != -1 {
		file, err := os.OpenFile(os.Args[ind+1], os.O_RDONLY, 0)

		if err != nil {
			exitWithError(err.Error())
		}
		defer file.Close()

		runTestWithParametersFromFile(parseConfigFile(file))
	} else if ind := slices.Index(os.Args, "-p"); ind != -1 {
		runTestForCmdArguments(parseConfigCmd(os.Args[ind+1:]))
	} else { //hand debug case
		solverConfs := [4]int{0, 5, 4, 1}
		generatorConfs := [3]int{0, 0, 15}
		runTestForCmdArguments(solverConfs, generatorConfs, 3.0, 20, 1)
	}

}

func exitWithError(err string) {
	fmt.Println(err)
	os.Exit(1)
}

func parseConfigFile(config *os.File) ([4]int, [3]int, float64, int, int) {
	scanner := bufio.NewScanner(config)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	confStr := scanner.Text()
	return parseConfig(confStr)
}

func parseConfigCmd(args []string) ([4]int, [3]int, float64, int, int) {
	return parseConfig(strings.Join(args, " "))
}

func parseConfig(config string) ([4]int, [3]int, float64, int, int) {

	confStrings := strings.Split(config, " ")

	confInts := make([]int, 0)
	floatValue := 0.0
	for i, str := range confStrings {
		if i == 5 && confInts[4] != 0 {
			confF, errF := strconv.ParseFloat(str, 64)
			if errF != nil {
				exitWithError(fmt.Sprint("ERR 1 In config file argument", i, " = ", str, " is invalid"))
			}
			floatValue = confF
			confInts = append(confInts, 0)
			continue
		}
		conf, err := strconv.Atoi(str)
		if err != nil {
			exitWithError(fmt.Sprint("ERR 2 In config file argument", i, " = ", str, " is invalid"))
		}
		confInts = append(confInts, conf)
	}
	solverConfs := [4]int{confInts[0], confInts[1], confInts[2], confInts[3]}
	generatorConfs := [3]int{confInts[4], confInts[5], confInts[6]}

	return solverConfs, generatorConfs, floatValue, confInts[7], confInts[8]

}

func createHeader(solverConf [4]int, genConf [3]int) string {
	header := ""

	if solverConf[0] == 0 {
		header += "PAGING"
	} else {
		header += "UPDATE_LIST"
	}

	header += "\n"

	numOfAlgs := 0
	if solverConf[2] == 0 {
		if solverConf[0] == 0 {
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
		if solverConf[0] == 0 {
			header += fmt.Sprintf("%s", pagingsolver.PagingAlg(solverConf[2]))
		} else {
			header += fmt.Sprintf("%s", updatelistsolver.UpdateListAlg(solverConf[2]))
		}
	} else {
		if solverConf[0] == 0 {
			for i := 0; i < numOfAlgs; i++ {
				header += fmt.Sprintf("%s ", pagingsolver.PagingAlg(i))
			}
		} else {
			for i := 0; i < numOfAlgs; i++ {
				header += fmt.Sprintf("%s ", updatelistsolver.UpdateListAlg(i))
			}
		}
	}

	header += "\n"

	header += fmt.Sprintf("%d", genConf[0])
	header += "\n"

	if genConf[0] == 0 {
		for i := 0; i < 4; i++ {
			header += fmt.Sprintf("%s ", dataGenerator.GeneratorTypeEnum(i))
		}
	} else {
		header += fmt.Sprintf("%s", dataGenerator.GeneratorTypeEnum(genConf[0]))
	}

	header += "\n"

	header += fmt.Sprintf("%d ", solverConf[1])

	header += "\n"

	return header
}

func runTestWithParametersFromFile(solverConf [4]int, genConf [3]int, genFV float64, noOfReq int, testIterations int) {
	noOfAlgs := 1

	if solverConf[2] == 0 {
		if solverConf[0] == 0 {
			noOfAlgs = 3
		}
		if solverConf[0] == 1 {
			noOfAlgs = 5
		}
	}
	resName := "data/res/" + filepath.Base(os.Args[2])
	fmt.Println(resName)
	f, err2 := os.OpenFile(resName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err2 != nil {
		exitWithError(err2.Error())
	}
	defer f.Close()

	header := createHeader(solverConf, genConf)
	fmt.Fprint(f, header)

	var name string
	var score int
	for j := 0; j < 10; j++ {
		ress := make([]int, noOfAlgs)
		names := make([]string, noOfAlgs)
		for iteration := 0; iteration < testIterations; iteration++ {
			pSS := solver.CreateSolver(solverConf, noOfAlgs)
			dG := dataGenerator.CreateDataGenerator(genConf, genFV)

			for i := 0; i < noOfReq; i++ {
				for _, pS := range pSS {
					pS.Serve(dG.GetRequest())
				}
			}

			for i, pS := range pSS {
				name, score = pS.Raport()
				names[i] = name
				ress[i] += int(float64(score) / float64(testIterations))
			}

		}
		saveResToFile(f, ress, noOfReq)
		fmt.Println("aa")
		noOfReq += 500
	}
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

func runTestForCmdArguments(solverConf [4]int, genConf [3]int, genFV float64, noOfReq int, testIterations int) {
	for iteration := 0; iteration < testIterations; iteration++ {
		pSS := solver.CreateSolver(solverConf, 1)
		dG := dataGenerator.CreateDataGenerator(genConf, genFV)

		for i := 0; i < noOfReq; i++ {
			for _, pS := range pSS {
				pS.Serve(dG.GetRequest())
			}
		}
		ress := 0
		for _, pS := range pSS {
			name, score := pS.Raport()
			ress += int(float64(score) / float64(testIterations))
			fmt.Println(name, ress, noOfReq)
		}

	}
}
