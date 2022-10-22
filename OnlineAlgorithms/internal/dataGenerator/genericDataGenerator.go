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

func CreateDataGenerator(generConf utils.GeneratorConfigS) GenericDataGenerator {
	var gD GenericDataGenerator
	switch utils.GeneratorTypeEnum(generConf.DistributionType + 1) {
	case utils.Uni:
		gD = unidistgenerator.Create(generConf.Minimum, generConf.Maximum)

	case utils.Geo:
		gD = geodistgenerator.Create(generConf.Fvalue, generConf.Maximum)

	case utils.Pois:
		gD = poisdistgenerator.Create(generConf.Fvalue, generConf.Maximum)

	case utils.Hrm:
		gD = hrmdistgenerator.Create(generConf.Minimum, generConf.Maximum)

	case utils.Dhr:
		gD = dhrdistgenerator.Create(generConf.Minimum, generConf.Maximum)
	}

	return gD
}
