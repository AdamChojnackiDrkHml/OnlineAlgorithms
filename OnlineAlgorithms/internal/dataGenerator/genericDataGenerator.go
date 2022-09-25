package datagenerator_test

import (
	geodistgenerator "OnlineAlgorithms/internal/dataGenerator/geoDistGenerator"
	poisdistgenerator "OnlineAlgorithms/internal/dataGenerator/poisDistGenerator"
	unidistgenerator "OnlineAlgorithms/internal/dataGenerator/uniDistGenerator"
)

type GeneratorTypeEnum int

const (
	All GeneratorTypeEnum = iota
	Uni
	Geo
	Pois
)

func (e GeneratorTypeEnum) String() string {
	switch e {
	case All:
		return "All"
	case Uni:
		return "Uni"
	case Geo:
		return "Geo"
	case Pois:
		return "Pois"
	default:
		return "NULL"
	}
}

type GenericDataGenerator interface {
	GetRequest() int
}

func CreateDataGenerator(conf [3]int, controlFloat float64) GenericDataGenerator {
	var gD GenericDataGenerator

	switch GeneratorTypeEnum(conf[0]) {
	case Uni:
		gD = unidistgenerator.Create(conf[1], conf[2])

	case Geo:
		gD = geodistgenerator.Create(controlFloat, conf[2])

	case Pois:
		gD = poisdistgenerator.Create(controlFloat, conf[2])
	}

	return gD
}
