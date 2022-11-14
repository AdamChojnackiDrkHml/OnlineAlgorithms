package datagenerator_test

import (
	dhrdistgenerator "OnlineAlgorithms/internal/dataGenerator/dhrDistGenerator"
	geodistgenerator "OnlineAlgorithms/internal/dataGenerator/geoDistGenerator"
	hrmdistgenerator "OnlineAlgorithms/internal/dataGenerator/hrmDistGenerator"
	poisdistgenerator "OnlineAlgorithms/internal/dataGenerator/poisDistGenerator"
	unidistgenerator "OnlineAlgorithms/internal/dataGenerator/uniDistGenerator"
	"OnlineAlgorithms/internal/utils"
)

type GenericDataGenerator interface {
	GetRequest() int
}

func CreateDataGenerator(generConf utils.GeneratorConfigS) []GenericDataGenerator {
	var generators []GenericDataGenerator

	if generConf.DoAll {
		for i := 0; i < utils.NUM_OF_DISTRIBUTIONS; i++ {
			generators = append(generators, initGenerator(generConf, utils.GeneratorTypeEnum(i)))
		}
	} else {
		generators = append(generators, initGenerator(generConf, generConf.DistributionType))
	}

	return generators
}

func initGenerator(generConf utils.GeneratorConfigS, generatorType utils.GeneratorTypeEnum) GenericDataGenerator {
	var gD GenericDataGenerator

	switch generatorType {
	case utils.Uni:
		gD = unidistgenerator.Create(generConf.Minimum, generConf.Maximum)

	case utils.Geo:
		gD = geodistgenerator.Create(generConf.FvalueGeo, generConf.Maximum)

	case utils.Pois:
		gD = poisdistgenerator.Create(generConf.FvaluePoiss, generConf.Maximum)

	case utils.Hrm:
		gD = hrmdistgenerator.Create(generConf.Minimum, generConf.Maximum)

	case utils.Dhr:
		gD = dhrdistgenerator.Create(generConf.Minimum, generConf.Maximum)
	}

	return gD
}
