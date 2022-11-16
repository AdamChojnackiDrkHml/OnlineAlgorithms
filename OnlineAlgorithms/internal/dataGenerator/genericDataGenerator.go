package datagenerator

import (
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
		gD = UNI_Create(generConf.Minimum, generConf.Maximum)

	case genUtils.Geo:
		gD = GEO_Create(generConf.FvalueGeo, generConf.Maximum)

	case genUtils.Pois:
		gD = POIS_Create(generConf.FvaluePoiss, generConf.Maximum)

	case genUtils.Hrm:
		gD = HRM_Create(generConf.Minimum, generConf.Maximum)

	case genUtils.Dhr:
		gD = DHR_Create(generConf.Minimum, generConf.Maximum)
	}

	return gD
}
