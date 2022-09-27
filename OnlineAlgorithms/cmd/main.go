package main

import (
	dataGenerator "OnlineAlgorithms/internal/dataGenerator"
	"OnlineAlgorithms/internal/solver"
	pagingsolver "OnlineAlgorithms/internal/solver/pagingSolver"
	updatelistsolver "OnlineAlgorithms/internal/solver/updateListSolver"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
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
	noOfAlgs := 1

	if solverConf[2] == 0 {
		if solverConf[0] == 0 {
			noOfAlgs = 3
		}
		if solverConf[0] == 1 {
			noOfAlgs = 5
		}
	}
	resName := "data/res/" + filepath.Base(os.Args[1])
	fmt.Println(resName)
	f, err2 := os.OpenFile(resName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err2 != nil {
		exitWithError(err.Error())
	}
	defer f.Close()

	header := createHeader(solverConf, genConf)
	fmt.Fprint(f, header)

	var name string
	var score int
	for i := 0; i < 10; i++ {
		ress := make([]int, noOfAlgs)
		names := make([]string, noOfAlgs)
		for iteration := 0; iteration < noOfIterations; iteration++ {
			pSS := solver.CreateSolver(solverConf, noOfAlgs)
			dG := dataGenerator.CreateDataGenerator(genConf, floatValue)

			for i := 0; i < noOfReq; i++ {
				for _, pS := range pSS {
					pS.Serve(dG.GetRequest())
				}
			}

			for i, pS := range pSS {
				name, score = pS.Raport()
				names[i] = name
				ress[i] += int(float64(score) / float64(noOfIterations))
			}

		}
		fmt.Fprint(f, noOfReq)
		fmt.Fprint(f, " ")
		for _, res := range ress {
			fmt.Fprint(f, res)
			fmt.Fprint(f, " ")
		}
		if i < 9 {
			fmt.Fprintln(f)
		}
		fmt.Println("aa")
		noOfReq += 500
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
