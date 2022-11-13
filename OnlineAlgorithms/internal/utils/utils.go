package utils

import (
	"fmt"
	"os"
	"strconv"

	"gopkg.in/yaml.v3"
)

const (
	NUM_OF_PAGING_ALGS     = 6
	NUM_OF_UPDATELIST_ALGS = 5
)

func GetMaxNumOfAlgs(solver SolverTypeEnum) int {
	switch solver {
	case Paging:
		return NUM_OF_PAGING_ALGS
	case UpdateList:
		return NUM_OF_UPDATELIST_ALGS
	default:
		return NUM_OF_PAGING_ALGS
	}
}

func DebugPrint(s string, p bool) {
	if p {
		fmt.Print(s)
	}
}

func ExitWithError(err string) {
	fmt.Println(err)
	os.Exit(1)
}

type GeneralConfigS struct {
	NoOfReq    int `yaml:"noOfReq"`
	Iterations int `yaml:"iterations"`
	Growth     int `yaml:"growth"`
	Repeats    int `yaml:"repeats"`
}
type SolverConfigS struct {
	ProblemType int  `yaml:"problemType"`
	Size        int  `yaml:"size"`
	Alg         int  `yaml:"alg"`
	Debug       bool `yaml:"debug"`
	DoAll       bool `yaml:"doAll"`
}
type GeneratorConfigS struct {
	DistributionType int     `yaml:"distributionType"`
	Minimum          int     `yaml:"minimum"`
	Fvalue           float64 `yaml:"fvalue"`
	Maximum          int     `yaml:"maximum"`
}

type TestConfigS struct {
	GeneralConfig GeneralConfigS `yaml:"generalConfig"`

	SolverConfig SolverConfigS `yaml:"solverConfig"`

	GeneratorConfig GeneratorConfigS `yaml:"generatorConfig"`
}

type Config struct {
	TestConfig TestConfigS `yaml:"test"`
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

func ParseCmd(confStrings []string) *Config {

	confInts := make([]int, 0)
	floatValue := 0.0
	for i, str := range confStrings {
		if i == 5 && confInts[5] != 0 {
			confF, errF := strconv.ParseFloat(str, 64)
			if errF != nil {
				ExitWithError(fmt.Sprint("ERR 1 In config file argument", i, " = ", str, " is invalid"))
			}
			floatValue = confF
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
	solverConf := SolverConfigS{confInts[0], confInts[1], confInts[2], confInts[3] == 1, confInts[4] == 1}
	generatorConf := GeneratorConfigS{confInts[5], confInts[6], floatValue, confInts[7]}

	return &Config{TestConfig: TestConfigS{genConf, solverConf, generatorConf}}

}

type SolverTypeEnum int

const (
	Paging SolverTypeEnum = iota
	UpdateList
)

type UpdateListAlg int

const (
	MTF UpdateListAlg = iota
	TRANS
	FQ
	BIT
	TS
)

func (e UpdateListAlg) String() string {
	switch e {
	case MTF:
		return "MTF"
	case TRANS:
		return "TRANS"
	case FQ:
		return "FQ"
	case BIT:
		return "BIT"
	case TS:
		return "TS"
	default:
		return "NULL"
	}
}

type PagingAlg int

const (
	LRU PagingAlg = iota
	FIFO
	LFU
	MARK
	MARK2
	RM
)

func (e PagingAlg) String() string {
	switch e {
	case LRU:
		return "LRU"
	case FIFO:
		return "FIFO"
	case LFU:
		return "LFU"
	case MARK:
		return "MARK"
	case MARK2:
		return "MARK2"
	case RM:
		return "RM"
	default:
		return "NULL"
	}
}

type GeneratorTypeEnum int

const (
	All GeneratorTypeEnum = iota
	Uni
	Geo
	Pois
	Hrm
	Dhr
)

func (e GeneratorTypeEnum) String() string {
	switch e {
	case All:
		return "All"
	case Uni:
		return "Uni"
	case Geo:
		return "Geo"
	case Pois:
		return "Pois"
	case Hrm:
		return "Hrm"
	case Dhr:
		return "Dhr"
	default:
		return "NULL"
	}
}
