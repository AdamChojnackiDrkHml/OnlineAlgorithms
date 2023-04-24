package main

import (
	dg "OnlineAlgorithms/pkg/dataGenerator"
	datagenerator "OnlineAlgorithms/pkg/dataGenerator/dataGeneratorConfigs"
	"OnlineAlgorithms/pkg/dataGenerator/distributions"
	"OnlineAlgorithms/pkg/solver"
	palgs "OnlineAlgorithms/pkg/solver/pagingSolver/pagingSolverAlgs"
	solverconfigs "OnlineAlgorithms/pkg/solver/solverConfigs"
	"fmt"
)

func main() {
	generatorConfig := &datagenerator.GeneratorConfigS{
		FvalueGeo:        0.5,
		Minimum:          0,
		Maximum:          40,
		DistributionType: []distributions.GeneratorTypeEnum{distributions.Hrm},
	}

	solverconfigs := &solverconfigs.SolverConfigS{
		DoAll: false,
		AlgP: []palgs.PagingAlg{
			palgs.FIFO,
			palgs.FWF,
			palgs.LRU,
			palgs.LFU,
			palgs.RM,
			palgs.RAND,
		},
		ProblemType: solverconfigs.Paging,
		Debug:       false,
		Size:        8,
	}

	problemSolvers := solver.CreateSolversFromConfig(*solverconfigs)
	dGS := dg.CreateDataGenerator(*generatorConfig)

	fmt.Println(singleDistroTest(problemSolvers, dGS[0]))

}

func singleDistroTest(problemSolvers []solver.GenericSolver, gen dg.GenericDataGenerator) []float64 {
	million := 1_000_000
	req := dg.GetSliceOfRequests(&gen, million)

	for _, request := range req {
		for _, solv := range problemSolvers {
			solv.Serve(request)
		}
	}

	res := make([]float64, len(problemSolvers))

	for i, solv := range problemSolvers {
		_, in := solv.Raport()
		res[i] = float64(in) / float64(million)
	}

	return res
}
