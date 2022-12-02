package ioutils

import (
	dgutils "OnlineAlgorithms/pkg/dataGenerator/dataGeneratorUtils"
	genUtils "OnlineAlgorithms/pkg/generalUtils"
	pSolver "OnlineAlgorithms/pkg/solver/pagingSolver"
	uLSolver "OnlineAlgorithms/pkg/solver/updateListSolver"
	solverUtils "OnlineAlgorithms/pkg/solver/utils"
	"fmt"
	"os"
	"strconv"

	"github.com/creasty/defaults"
	"gopkg.in/yaml.v3"
)

// /////////////////////////////
// GENERAL IO EXPORTED FUNCTIONS
// /////////////////////////////
func SaveResToFile(f *os.File, ress []int, noOfReq int) {
	WriteToFile(f, fmt.Sprint(noOfReq))
	WriteToFile(f, " ")
	for _, res := range ress {
		WriteToFile(f, fmt.Sprint(res))
		WriteToFile(f, " ")
	}
	WriteToFile(f, "\n")
}

func OpenFile(name string) *os.File {
	file, err2 := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err2 != nil {
		ExitWithError(err2.Error())
	}

	return file
}

func WriteToFile(f *os.File, s string) {
	if f == nil {
		ExitWithError("Nil file")
	}
	fmt.Fprint(f, s)

}

func ExitWithError(err string) {
	fmt.Println(err)
	os.Exit(1)
}

// ////////////////////////////////////////////
// NOT EXPORTED YAML PARSING SUPPORT FUNCTIONS
// ////////////////////////////////////////////
func readFile(path string) ([]byte, error) {
	fileContent, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return fileContent, nil
}

func parseYaml(configYaml []byte) (*genUtils.Config, error) {

	config := &genUtils.Config{}
	defaults.Set(config)
	// Open config file

	// Init new YAML decode
	err := yaml.Unmarshal(configYaml, config)

	// Start YAML decoding from file
	if err != nil {
		return nil, err
	}

	return config, nil
}

// ////////////////////////////////////////////////////
// FUNCTIONS FOR CREATING CONFIG STRUCTURES FROM INPUT
// ////////////////////////////////////////////////////
func ReadYamlForConfig(path string) *genUtils.Config {
	yamlContent, err := readFile(path)

	if err != nil {
		ExitWithError(err.Error())
	}

	config, err1 := parseYaml(yamlContent)

	if err1 != nil {
		ExitWithError(err1.Error())
	}

	return config
}

func ParseCmd(confStrings []string) genUtils.Config {

	confInts := make([]int, 0)
	floatValueGeo := 0.0
	floatValuePoiss := 0.0
	for i, str := range confStrings {
		if i == 5 && confInts[5] != 0 {
			confF, errF := strconv.ParseFloat(str, 64)
			if errF != nil {
				ExitWithError(fmt.Sprint("ERR 1 In config file argument", i, " = ", str, " is invalid"))
			}
			floatValueGeo = confF
			confInts = append(confInts, 0)
			continue
		}
		if i == 6 && confInts[6] != 0 {
			confF, errF := strconv.ParseFloat(str, 64)
			if errF != nil {
				ExitWithError(fmt.Sprint("ERR 1 In config file argument", i, " = ", str, " is invalid"))
			}
			floatValuePoiss = confF
			confInts = append(confInts, 0)
			continue
		}
		conf, err := strconv.Atoi(str)
		if err != nil {
			ExitWithError(fmt.Sprint("ERR 2 In config file argument", i, " = ", str, " is invalid"))
		}
		confInts = append(confInts, conf)
	}
	genConf := genUtils.GeneralConfigS{NoOfReq: confInts[7], Iterations: confInts[8], Growth: confInts[9], Repeats: confInts[10]}
	solverConf := solverUtils.SolverConfigS{ProblemType: solverUtils.SolverTypeEnum(confInts[0]), Size: confInts[1], AlgP: []pSolver.PagingAlg{pSolver.PagingAlg(confInts[2])}, AlgUL: []uLSolver.UpdateListAlg{uLSolver.UpdateListAlg(confInts[2])}, Debug: confInts[3] == 1, DoAll: confInts[4] == 1}
	generatorConf := dgutils.GeneratorConfigS{DistributionType: []dgutils.GeneratorTypeEnum{dgutils.GeneratorTypeEnum(confInts[5])}, Minimum: confInts[6], FvalueGeo: floatValueGeo, FvaluePoiss: floatValuePoiss, Maximum: confInts[7], DoAll: confInts[8] == 1}

	return genUtils.Config{TestConfigs: []genUtils.TestConfigS{{GeneralConfig: genConf, SolverConfig: solverConf, GeneratorConfig: generatorConf}}}

}

// /////////////////////////////////
// CREATING HEADER FOR RESULT FILES
// /////////////////////////////////
func CreateAndWriteHeader(f *os.File, solverConf *solverUtils.SolverConfigS, genConf *dgutils.GeneratorConfigS) {
	header := createHeader(solverConf, genConf)

	WriteToFile(f, header)
}

func createHeader(solverConf *solverUtils.SolverConfigS, genConf *dgutils.GeneratorConfigS) string {
	header := ""

	header += solverConf.ProblemType.String()

	header += "\n"

	numOfAlgs := solverConf.GetNumOfAlgs()
	header += fmt.Sprint(numOfAlgs)
	header += "\n"

	switch solverConf.ProblemType {
	case solverUtils.Paging:
		for _, algP := range solverConf.AlgP {
			header += algP.String() + " "
		}
	case solverUtils.UpdateList:
		for _, algUL := range solverConf.AlgUL {
			header += algUL.String() + " "
		}
	}

	header += "\n"

	header += fmt.Sprint(genConf.GetNumOfDistributions())

	header += "\n"

	for _, distribution := range genConf.DistributionType {
		header += distribution.String() + " "
	}

	header += "\n"

	header += fmt.Sprintf("%d ", solverConf.Size)

	header += "\n"
	return header
}
