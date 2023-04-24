package main

import (
	dg "OnlineAlgorithms/pkg/dataGenerator"
	datagenerator "OnlineAlgorithms/pkg/dataGenerator/dataGeneratorConfigs"
	"OnlineAlgorithms/pkg/dataGenerator/distributions"
	"OnlineAlgorithms/pkg/solver"
	palgs "OnlineAlgorithms/pkg/solver/pagingSolver/pagingSolverAlgs"
	solverconfigs "OnlineAlgorithms/pkg/solver/solverConfigs"
	ioutils "OnlineAlgorithms/pkg/testFramework/ioUtils"
	"fmt"
	"sync"

	"os"
)

func main() {
	fmt.Println(os.Getwd())
	Ns := []int{20, 30, 40, 50, 60, 70, 80, 90, 100}
	KsRatios2 := make([][]int, len(Ns))

	for i := range KsRatios2 {
		low := Ns[i] / 10
		high := Ns[i] / 5
		KsRatios2[i] = make([]int, 0)

		for j := low; j <= high; j++ {
			KsRatios2[i] = append(KsRatios2[i], j)
		}
		fmt.Println(KsRatios2[i])
	}
	generatorConfig := &datagenerator.GeneratorConfigS{
		FvalueGeo: 0.5,
		Minimum:   0,
		// Maximum:   40,
		DistributionType: []distributions.GeneratorTypeEnum{
			distributions.Uni,
			distributions.Geo,
			distributions.Hrm,
			distributions.Dhr,
		},
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
		// Size:        8,
	}

	templateFilename := "data/labs/"
	var wg sync.WaitGroup

	results := make([][][][]float64, len(Ns))

	for j, N := range Ns {

		results[j] = make([][][]float64, len(KsRatios2[j]))
		for i, cache := range KsRatios2[j] {
			generatorConfig.Maximum = N
			solverconfigs.Size = cache

			dGS := dg.CreateDataGenerator(*generatorConfig)
			results[j][i] = make([][]float64, len(dGS))

			for distIter, dist := range dGS {
				wg.Add(1)
				cacheSize := i
				setN := j
				distr := distIter
				problemSolvers := solver.CreateSolversFromConfig(*solverconfigs)

				generator := dist
				go func() {
					defer wg.Done()
					results[setN][cacheSize][distr] = singleDistroTest(problemSolvers, generator)
					fmt.Println(setN, cacheSize)

				}()
			}

		}

	}
	wg.Wait()

	for setN, N := range Ns {
		resFilename := templateFilename + fmt.Sprintf("%v_", N)

		for distro := range generatorConfig.DistributionType {
			f := ioutils.CreateAndOpenResFile(resFilename + distributions.FromInt(distro) + "2.txt")
			// ioutils.CreateAndWriteHeader(f, solverconfigs, generatorConfig)
			for ratio, cache := range KsRatios2[setN] {

				ioutils.SaveResToFile(f, results[setN][ratio][distro], cache)

			}
			f.Close()
		}
	}
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
