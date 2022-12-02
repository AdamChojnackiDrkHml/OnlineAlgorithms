package datagenerator

import "errors"

const NUM_OF_DISTRIBUTIONS = 5

type GeneratorTypeEnum int

const (
	Uni GeneratorTypeEnum = iota
	Geo
	Pois
	Hrm
	Dhr
)

func (e GeneratorTypeEnum) String() string {
	switch e {
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

type GeneratorConfigS struct {
	DistributionType []GeneratorTypeEnum `yaml:"distributionType"`
	Minimum          int                 `yaml:"minimum"`
	FvalueGeo        float64             `yaml:"fvalueGeo"`
	FvaluePoiss      float64             `yaml:"fvaluePoiss"`
	Maximum          int                 `yaml:"maximum"`
	DoAll            bool                `default:"false" yaml:"doAll"`
}

func (generatorConfig *GeneratorConfigS) GetNumOfDistributions() int {
	return len(generatorConfig.DistributionType)
}

func (generatorConfig *GeneratorConfigS) Preprocess() {
	if generatorConfig.DoAll {
		generatorConfig.DistributionType = make([]GeneratorTypeEnum, 0)

		for i := 0; i < NUM_OF_DISTRIBUTIONS; i++ {
			generatorConfig.DistributionType = append(generatorConfig.DistributionType, GeneratorTypeEnum(i))
		}
	}
}

func (generatorConfig *GeneratorConfigS) Validate() error {
	for _, distribution := range generatorConfig.DistributionType {
		if distribution >= NUM_OF_DISTRIBUTIONS {
			return errors.New("wrong distribution identification number")
		}
	}

	return nil
}
