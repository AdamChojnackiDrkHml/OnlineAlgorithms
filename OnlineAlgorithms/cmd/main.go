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

	solverConf, genConf, floatValue, iterator := parseConfig(file)

	pSS := solver.CreateSolver(solverConf)
	dG := dataGenerator.CreateDataGenerator(genConf, floatValue)

	for i := 0; i < iterator; i++ {
		for _, pS := range pSS {
			pS.Serve(dG.GetRequest())
		}
	}

	for _, pS := range pSS {
		fmt.Println(pS.Raport())
	}

}

func exitWithError(err string) {
	fmt.Println(err)
	os.Exit(1)
}

func parseConfig(config *os.File) ([4]int, [3]int, float64, int) {
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

	return solverConfs, generatorConfs, floatValue, confInts[7]

}
