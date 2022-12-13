// Package datagenerator defines front for all used data generators.
package datagenerator

import (
	dgconf "OnlineAlgorithms/pkg/dataGenerator/dataGeneratorConfigs"
	dist "OnlineAlgorithms/pkg/dataGenerator/distributions"
)

// GenericDataGenerator provides fron for any data generator implementing it.
type GenericDataGenerator interface {
	GetRequest() int
}

// CreateDataGenerator function takes GeneratorConfigS struct
// and returns slice of GenericDataGenerator based on its contents.
func CreateDataGenerator(generConf dgconf.GeneratorConfigS) []GenericDataGenerator {
	var generators []GenericDataGenerator

	for _, n := range generConf.DistributionType {
		generators = append(generators, initGenerator(generConf, n))
	}

	return generators
}

func initGenerator(generConf dgconf.GeneratorConfigS, generatorType dist.GeneratorTypeEnum) GenericDataGenerator {
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
