package datagenerator_test

import (
	dhrdistgenerator "OnlineAlgorithms/internal/dataGenerator/dhrDistGenerator"
	geodistgenerator "OnlineAlgorithms/internal/dataGenerator/geoDistGenerator"
	hrmdistgenerator "OnlineAlgorithms/internal/dataGenerator/hrmDistGenerator"
	poisdistgenerator "OnlineAlgorithms/internal/dataGenerator/poisDistGenerator"
	unidistgenerator "OnlineAlgorithms/internal/dataGenerator/uniDistGenerator"
	genUtils "OnlineAlgorithms/internal/utils/generalUtils"
)

type GenericDataGenerator interface {
	GetRequest() int
}

func CreateDataGenerator(generConf genUtils.GeneratorConfigS) []GenericDataGenerator {
	var generators []GenericDataGenerator

	for _, n := range generConf.DistributionType {
		generators = append(generators, initGenerator(generConf, n))
	}

	return generators
}

func initGenerator(generConf genUtils.GeneratorConfigS, generatorType genUtils.GeneratorTypeEnum) GenericDataGenerator {
	var gD GenericDataGenerator

	switch generatorType {
	case genUtils.Uni:
		gD = unidistgenerator.Create(generConf.Minimum, generConf.Maximum)

	case genUtils.Geo:
		gD = geodistgenerator.Create(generConf.FvalueGeo, generConf.Maximum)

	case genUtils.Pois:
		gD = poisdistgenerator.Create(generConf.FvaluePoiss, generConf.Maximum)

	case genUtils.Hrm:
		gD = hrmdistgenerator.Create(generConf.Minimum, generConf.Maximum)

	case genUtils.Dhr:
		gD = dhrdistgenerator.Create(generConf.Minimum, generConf.Maximum)
	}

	return gD
}
