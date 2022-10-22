package utils

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func DebugPrint(s string, p bool) {
	if p {
		fmt.Print(s)
	}
}

type Config struct {
	Test struct {
		GeneralConfig struct {
			NoOfReq    int `yaml:"noOfReq"`
			Iterations int `yaml:"iterations"`
			Growth     int `yaml:"growth"`
		} `yaml:"generalConfig"`

		SolverConfig struct {
			ProblemType int  `yaml:"problemType"`
			Size        int  `yaml:"size"`
			Alg         int  `yaml:"alg"`
			Debug       bool `yaml:"debug"`
		} `yaml:"solverConfig"`

		GeneratorConfig struct {
			DistributionType int     `yaml:"distributionType"`
			Minimum          int     `yaml:"minimum"`
			Fvalue           float64 `yaml:"fvalue"`
			Maximum          int     `yaml:"maximum"`
		} `yaml:"generatorConfig"`
	} `yaml:"test"`
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
