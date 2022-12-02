package datagenerator

import (
	dGUtils "OnlineAlgorithms/pkg/dataGenerator/dataGeneratorUtils"
	dist "OnlineAlgorithms/pkg/dataGenerator/distributions"
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

func initGenerator(generConf dGUtils.GeneratorConfigS, generatorType dist.GeneratorTypeEnum) GenericDataGenerator {
	var gD GenericDataGenerator

	switch generatorType {
	case dist.Uni:
		gD = dist.UNI_Create(generConf.Minimum, generConf.Maximum)

	case dist.Geo:
		gD = dist.GEO_Create(generConf.FvalueGeo, generConf.Maximum)

	case dist.Pois:
		gD = dist.POIS_Create(generConf.FvaluePoiss, generConf.Maximum)

	case dist.Hrm:
		gD = dist.HRM_Create(generConf.Minimum, generConf.Maximum)

	case dist.Dhr:
		gD = dist.DHR_Create(generConf.Minimum, generConf.Maximum)
	}

	return gD
}
