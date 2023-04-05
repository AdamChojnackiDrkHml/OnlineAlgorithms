package main

import (
	generalutils "OnlineAlgorithms/pkg/configuration"
	datagenerator "OnlineAlgorithms/pkg/dataGenerator/dataGeneratorConfigs"
	"OnlineAlgorithms/pkg/dataGenerator/distributions"
	solverconfigs "OnlineAlgorithms/pkg/solver/solverConfigs"
	testcontroler "OnlineAlgorithms/pkg/testFramework/testControler"
	"fmt"
	"math"
	"os"
)

func main() {
	fmt.Println(os.Getwd())
	Ns := []int{20, 30, 40, 50, 60, 70, 80, 90}
	KsRatios := []float64{1.0 / 10.0, 1.0 / 9.0, 1.0 / 8.0, 1.0 / 7.0, 1.0 / 6.0, 1.0 / 5.0}
	Distrios := []distributions.GeneratorTypeEnum{distributions.Uni, distributions.Geo, distributions.Hrm, distributions.Dhr}

	generatorConfig := &datagenerator.GeneratorConfigS{
		FvalueGeo: 0.5,
		Minimum:   0,
	}

	solverconfigs := &solverconfigs.SolverConfigS{
		DoAll:       true,
		ProblemType: solverconfigs.Paging,
		Debug:       false,
	}

	generalConfig := &generalutils.GeneralConfigS{
		NoOfReq:    10000,
		Iterations: 10,
		Growth:     10000,
		Repeats:    5,
	}

	for _, N := range Ns {
		for _, KRatio := range KsRatios {
			config := generalutils.Config{
				TestConfigs: make([]generalutils.TestConfigS, len(Distrios)),
			}
			for i, dist := range Distrios {
				testConfig := &generalutils.TestConfigS{}
				generatorConfig.DistributionType = []distributions.GeneratorTypeEnum{dist}
				generatorConfig.Maximum = N
				solverconfigs.Size = int(math.Ceil(float64(N) * KRatio))
				testConfig.GeneralConfig = *generalConfig
				testConfig.GeneratorConfig = *generatorConfig
				testConfig.SolverConfig = *solverconfigs
				config.TestConfigs[i] = *testConfig
			}

			resFilename := "data/res/" + "results_" + fmt.Sprintf("N=%v_K=%v_", N, int(float64(N)*KRatio))
			testcontroler.RunTestForConfig(&config, resFilename)

		}
	}
}
