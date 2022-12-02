package datagenerator

import (
	dist "OnlineAlgorithms/pkg/dataGenerator/distributions"
	"errors"
)

type GeneratorConfigS struct {
	DistributionType []dist.GeneratorTypeEnum `yaml:"distributionType"`
	Minimum          int                      `yaml:"minimum"`
	FvalueGeo        float64                  `yaml:"fvalueGeo"`
	FvaluePoiss      float64                  `yaml:"fvaluePoiss"`
	Maximum          int                      `yaml:"maximum"`
	DoAll            bool                     `default:"false" yaml:"doAll"`
}

func (generatorConfig *GeneratorConfigS) GetNumOfDistributions() int {
	return len(generatorConfig.DistributionType)
}

func (generatorConfig *GeneratorConfigS) Preprocess() {
	if generatorConfig.DoAll {
		generatorConfig.DistributionType = make([]dist.GeneratorTypeEnum, 0)

		for i := 0; i < dist.NUM_OF_DISTRIBUTIONS; i++ {
			generatorConfig.DistributionType = append(generatorConfig.DistributionType, dist.GeneratorTypeEnum(i))
		}
	}
}

func (generatorConfig *GeneratorConfigS) Validate() error {
	for _, distribution := range generatorConfig.DistributionType {
		if distribution >= dist.NUM_OF_DISTRIBUTIONS {
			return errors.New("wrong distribution identification number")
		}
	}

	return nil
}
