package main

import (
	dataGenerator "OnlineAlgorithms/internal/dataGenerator"
	"OnlineAlgorithms/internal/solver"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		exitWithError("Pass only file name as program argument")
	}

	file, err := os.OpenFile(os.Args[1], os.O_RDONLY, 0)

	if err != nil {
		exitWithError(err.Error())
	}
	defer file.Close()

	solverConf, genConf, floatValue, noOfReq, noOfIterations := parseConfig(file)
	noOfRes := 0

	if solverConf[2] == 0 {
		noOfRes = 5
	} else {
		noOfRes = 1
	}

	res := make([]float64, noOfRes)
	names := make([]string, noOfRes)
	var name string
	var score int
	for iteration := 0; iteration < noOfIterations; iteration++ {
		pSS := solver.CreateSolver(solverConf)
		dG := dataGenerator.CreateDataGenerator(genConf, floatValue)

		for i := 0; i < noOfReq; i++ {
			for _, pS := range pSS {
				pS.Serve(dG.GetRequest())
			}
		}

		for i, pS := range pSS {
			name, score = pS.Raport()
			names[i] = name
			res[i] += float64(score) / float64(noOfIterations)
		}

	}

	for i, n := range names {
		fmt.Println(n, " - ", res[i])
	}

}

func exitWithError(err string) {
	fmt.Println(err)
	os.Exit(1)
}

func parseConfig(config *os.File) ([4]int, [3]int, float64, int, int) {
	scanner := bufio.NewScanner(config)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	confStr := scanner.Text()
	confStrings := strings.Split(confStr, " ")

	confInts := make([]int, 0)
	floatValue := 0.0
	for i, str := range confStrings {
		if i == 5 && confInts[4] != 0 {
			confF, errF := strconv.ParseFloat(str, 64)
			if errF != nil {
				exitWithError(fmt.Sprint("In config file argument", i, " = ", str, " is invalid"))
			}
			floatValue = confF
			confInts = append(confInts, 0)
			continue
		}
		conf, err := strconv.Atoi(str)
		if err != nil {
			exitWithError(fmt.Sprint("In config file argument", i, " = ", str, " is invalid"))
		}
		confInts = append(confInts, conf)
	}
	solverConfs := [4]int{confInts[0], confInts[1], confInts[2], confInts[3]}
	generatorConfs := [3]int{confInts[4], confInts[5], confInts[6]}

	return solverConfs, generatorConfs, floatValue, confInts[7], confInts[8]

}
