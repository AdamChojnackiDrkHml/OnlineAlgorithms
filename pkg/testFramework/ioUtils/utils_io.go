// Package testioutils defines functionalities to read test configuration
// and save results to file.
package testioutils

import (
	conf "OnlineAlgorithms/pkg/configuration"
	dgconf "OnlineAlgorithms/pkg/dataGenerator/dataGeneratorConfigs"
	dist "OnlineAlgorithms/pkg/dataGenerator/distributions"
	psalgs "OnlineAlgorithms/pkg/solver/pagingSolver/pagingSolverAlgs"
	svconf "OnlineAlgorithms/pkg/solver/solverConfigs"
	ulsalgs "OnlineAlgorithms/pkg/solver/updateListSolver/updateListSolverAlgs"
	"fmt"
	"os"
	"strconv"

	"github.com/creasty/defaults"
	"gopkg.in/yaml.v3"
)

// /////////////////////////////
// GENERAL IO EXPORTED FUNCTIONS
// /////////////////////////////

// SaveResToFile takes results slice and number of requests
// and writes that to passed file.
func SaveResToFile(f *os.File, ress []int, noOfReq int) {
	writeToFile(f, fmt.Sprint(noOfReq))
	writeToFile(f, " ")
	for _, res := range ress {
		writeToFile(f, fmt.Sprint(res))
		writeToFile(f, " ")
	}
	writeToFile(f, "\n")
}

// CreateAndOpenResFile takes path containing name of new results file
// and then creates it.
// Returns file descriptor of opened file.
func CreateAndOpenResFile(pathName string) *os.File {
	file, err2 := os.OpenFile(pathName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err2 != nil {
		exitWithError(err2.Error())
	}

	return file
}

func writeToFile(f *os.File, s string) {
	if f == nil {
		exitWithError("Nil file")
	}
	fmt.Fprint(f, s)

}

func exitWithError(err string) {
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

func parseYaml(configYaml []byte) (*conf.Config, error) {

	config := &conf.Config{}
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

// ReadYamlForConfig takes path to yaml configuration file
// and tries to parse it to Config struct.
// On fail exits program with error.
// On success returns parsed configuration.
func ReadYamlForConfig(path string) *conf.Config {
	yamlContent, err := readFile(path)

	if err != nil {
		exitWithError(err.Error())
	}

	config, err1 := parseYaml(yamlContent)

	if err1 != nil {
		exitWithError(err1.Error())
	}

	return config
}

// ParseCmdForConfig takes slice of configuration strings
// and tries to parse it to Config struct.
// On fail exits with error.
// On success returns parsed configuration.
// This function is able to parse only single testcase configuration!
func ParseCmdForConfig(confStrings []string) conf.Config {

	confInts := make([]int, 0)
	floatValueGeo := 0.0
	floatValuePoiss := 0.0
	for i, str := range confStrings {
		if i == 5 && confInts[5] != 0 {
			confF, errF := strconv.ParseFloat(str, 64)
			if errF != nil {
				exitWithError(fmt.Sprint("ERR 1 In config file argument", i, " = ", str, " is invalid"))
			}
			floatValueGeo = confF
			confInts = append(confInts, 0)
			continue
		}
		if i == 6 && confInts[6] != 0 {
			confF, errF := strconv.ParseFloat(str, 64)
			if errF != nil {
				exitWithError(fmt.Sprint("ERR 1 In config file argument", i, " = ", str, " is invalid"))
			}
			floatValuePoiss = confF
			confInts = append(confInts, 0)
			continue
		}
		conf, err := strconv.Atoi(str)
		if err != nil {
			exitWithError(fmt.Sprint("ERR 2 In config file argument", i, " = ", str, " is invalid"))
		}
		confInts = append(confInts, conf)
	}
	genConf := conf.GeneralConfigS{
		NoOfReq:    confInts[7],
		Iterations: confInts[8],
		Growth:     confInts[9],
		Repeats:    confInts[10]}

	solverConf := svconf.SolverConfigS{
		ProblemType: svconf.SolverTypeEnum(confInts[0]),
		Size:        confInts[1],
		AlgP:        []psalgs.PagingAlg{psalgs.PagingAlg(confInts[2])},
		AlgUL:       []ulsalgs.UpdateListAlg{ulsalgs.UpdateListAlg(confInts[2])},
		Debug:       confInts[3] == 1,
		DoAll:       confInts[4] == 1}

	generatorConf := dgconf.GeneratorConfigS{
		DistributionType: []dist.GeneratorTypeEnum{dist.GeneratorTypeEnum(confInts[5])},
		Minimum:          confInts[6],
		FvalueGeo:        floatValueGeo,
		FvaluePoiss:      floatValuePoiss,
		Maximum:          confInts[7],
		DoAll:            confInts[8] == 1}

	return conf.Config{
		TestConfigs: []conf.TestConfigS{{
			GeneralConfig:   genConf,
			SolverConfig:    solverConf,
			GeneratorConfig: generatorConf}}}

}

// /////////////////////////////////
// CREATING HEADER FOR RESULT FILES
// /////////////////////////////////

// CreateAndWriteHeader takes solver and generator configuration
// and write test header info to passed results file descriptor.
func CreateAndWriteHeader(f *os.File, solverConf *svconf.SolverConfigS, genConf *dgconf.GeneratorConfigS) {
	header := createHeader(solverConf, genConf)

	writeToFile(f, header)
}

func createHeader(solverConf *svconf.SolverConfigS, genConf *dgconf.GeneratorConfigS) string {
	header := ""

	header += solverConf.ProblemType.String()

	header += "\n"

	numOfAlgs := solverConf.GetNumOfAlgs()
	header += fmt.Sprint(numOfAlgs)
	header += "\n"

	switch solverConf.ProblemType {
	case svconf.Paging:
		for _, algP := range solverConf.AlgP {
			header += algP.String() + " "
		}
	case svconf.UpdateList:
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
