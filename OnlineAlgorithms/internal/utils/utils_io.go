package utils

import (
	"fmt"
	"os"
	"strconv"

	"gopkg.in/yaml.v3"
)

func DebugPrint(s string, p bool) {
	if p {
		fmt.Print(s)
	}
}

func ExitWithError(err string) {
	fmt.Println(err)
	os.Exit(1)
}

func ParseYaml(configPath string) (*Config, error) {

	config := &Config{}

	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

func ParseCmd(confStrings []string) TestConfigS {

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
	genConf := GeneralConfigS{confInts[7], confInts[8], confInts[9], confInts[10]}
	solverConf := SolverConfigS{SolverTypeEnum(confInts[0]), confInts[1], PagingAlg(confInts[2]), UpdateListAlg(confInts[2]), confInts[3] == 1, confInts[4] == 1}
	generatorConf := GeneratorConfigS{GeneratorTypeEnum(confInts[5]), confInts[6], floatValueGeo, floatValuePoiss, confInts[7], confInts[8] == 1}

	return TestConfigS{genConf, solverConf, generatorConf}

}

func CreateHeader(f *os.File, solverConf *SolverConfigS, genConf *GeneratorConfigS) {
	header := ""

	header += solverConf.ProblemType.String()

	header += "\n"

	numOfAlgs := GetNumOfAlgs(solverConf.ProblemType, solverConf.DoAll)

	header += "\n"

	fmt.Println(numOfAlgs)

	if !solverConf.DoAll {
		switch solverConf.ProblemType {
		case Paging:
			header += solverConf.AlgP.String()
		case UpdateList:
			header += solverConf.AlgUL.String()
		}

	} else {

		switch solverConf.ProblemType {
		case Paging:
			for i := 0; i < numOfAlgs; i++ {
				header += PagingAlg(i).String()
			}
		case UpdateList:
			for i := 0; i < numOfAlgs; i++ {
				header += UpdateListAlg(i).String()
			}
		}

	}

	header += "\n"

	header += fmt.Sprintf("%d", genConf.DistributionType)
	header += "\n"

	if genConf.DoAll {
		for i := 0; i < NUM_OF_DISTRIBUTIONS; i++ {
			header += GeneratorTypeEnum(i).String()
		}
	} else {
		header += genConf.DistributionType.String()
	}

	header += "\n"

	header += fmt.Sprintf("%d ", solverConf.Size)

	header += "\n"

	WriteToFile(f, header)
}

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
