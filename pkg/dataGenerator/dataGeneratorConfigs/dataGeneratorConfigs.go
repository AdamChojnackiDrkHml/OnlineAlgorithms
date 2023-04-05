package datagenerator

import (
	dist "OnlineAlgorithms/pkg/dataGenerator/distributions"
	"errors"
)

// GeneratorConfigS holds data generator configuration.
// By default DoAll(distributions) flag is set to false.
type GeneratorConfigS struct {
	DistributionType []dist.GeneratorTypeEnum `yaml:"distributionType"`
	Minimum          int                      `yaml:"minimum"`
	FvalueGeo        float64                  `yaml:"fvalueGeo"`
	FvaluePoiss      float64                  `yaml:"fvaluePoiss"`
	Maximum          int                      `yaml:"maximum"`
	DoAll            bool                     `default:"false" yaml:"doAll"`
}

// GetNumOfDistributions returns number of data generators in configuration for set problem.
func (generatorConfig *GeneratorConfigS) GetNumOfDistributions() int {
	return len(generatorConfig.DistributionType)
}

// GetMaxNumOfAlgs returns number of defined distributions for set problem.
func GetMaxNumOfAlgs() int {
	return dist.NUM_OF_DISTRIBUTIONS
}

// Preprocess method should be called when using DoAll flag
// in order to fill proper distributions slice.
func (generatorConfig *GeneratorConfigS) Preprocess() {
	if generatorConfig.DoAll {
		generatorConfig.DistributionType = make([]dist.GeneratorTypeEnum, 0)

		for i := 0; i < dist.NUM_OF_DISTRIBUTIONS; i++ {
			generatorConfig.DistributionType = append(generatorConfig.DistributionType, dist.GeneratorTypeEnum(i))
		}
	}
}

// Validate checks if configuration is correct
// and will not cause errors in runtime.
// Returns error if finds incorrect config.
// Else returns nil.
func (generatorConfig *GeneratorConfigS) Validate() error {
	for _, distribution := range generatorConfig.DistributionType {
		if distribution >= dist.NUM_OF_DISTRIBUTIONS {
			return errors.New("wrong distribution identification number")
		}
	}

	return nil
}
