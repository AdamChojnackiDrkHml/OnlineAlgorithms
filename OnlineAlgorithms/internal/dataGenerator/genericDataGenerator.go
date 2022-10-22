package datagenerator_test

import (
	dhrdistgenerator "OnlineAlgorithms/internal/dataGenerator/dhrDistGenerator"
	geodistgenerator "OnlineAlgorithms/internal/dataGenerator/geoDistGenerator"
	hrmdistgenerator "OnlineAlgorithms/internal/dataGenerator/hrmDistGenerator"
	poisdistgenerator "OnlineAlgorithms/internal/dataGenerator/poisDistGenerator"
	unidistgenerator "OnlineAlgorithms/internal/dataGenerator/uniDistGenerator"
)

type GeneratorTypeEnum int

const (
	All GeneratorTypeEnum = iota
	Uni
	Geo
	Pois
	Hrm
	Dhr
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
	case Hrm:
		return "Hrm"
	case Dhr:
		return "Dhr"
	default:
		return "NULL"
	}
}

type GenericDataGenerator interface {
	GetRequest() int
}

func CreateDataGenerator(conf [3]int, controlFloat float64) GenericDataGenerator {
	var gD GenericDataGenerator
	switch GeneratorTypeEnum(conf[0] + 1) {
	case Uni:
		gD = unidistgenerator.Create(conf[1], conf[2])

	case Geo:
		gD = geodistgenerator.Create(controlFloat, conf[2])

	case Pois:
		gD = poisdistgenerator.Create(controlFloat, conf[2])

	case Hrm:
		gD = hrmdistgenerator.Create(conf[1], conf[2])

	case Dhr:
		gD = dhrdistgenerator.Create(conf[1], conf[2])
	}

	return gD
}
