package datagenerator

import (
	dGUtils "OnlineAlgorithms/pkg/dataGenerator/dataGeneratorUtils"
	distributions "OnlineAlgorithms/pkg/dataGenerator/distributions"
)

type GenericDataGenerator interface {
	GetRequest() int
}

func CreateDataGenerator(generConf dGUtils.GeneratorConfigS) []GenericDataGenerator {
	var generators []GenericDataGenerator

	for _, n := range generConf.DistributionType {
		generators = append(generators, initGenerator(generConf, n))
	}

	return generators
}

func initGenerator(generConf dGUtils.GeneratorConfigS, generatorType dGUtils.GeneratorTypeEnum) GenericDataGenerator {
	var gD GenericDataGenerator

	switch generatorType {
	case dGUtils.Uni:
		gD = distributions.UNI_Create(generConf.Minimum, generConf.Maximum)

	case dGUtils.Geo:
		gD = distributions.GEO_Create(generConf.FvalueGeo, generConf.Maximum)

	case dGUtils.Pois:
		gD = distributions.POIS_Create(generConf.FvaluePoiss, generConf.Maximum)

	case dGUtils.Hrm:
		gD = distributions.HRM_Create(generConf.Minimum, generConf.Maximum)

	case dGUtils.Dhr:
		gD = distributions.DHR_Create(generConf.Minimum, generConf.Maximum)
	}

	return gD
}
