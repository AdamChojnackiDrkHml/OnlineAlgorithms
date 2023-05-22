package main

import (
	datagenerator "OnlineAlgorithms/pkg/dataGenerator"
	datageneratorC "OnlineAlgorithms/pkg/dataGenerator/dataGeneratorConfigs"
	"OnlineAlgorithms/pkg/dataGenerator/distributions"
	bigpackingsolver "OnlineAlgorithms/pkg/solver/bigPackingSolver"
	bigpackingalgs "OnlineAlgorithms/pkg/solver/bigPackingSolver/bigPackingAlgs"
	"fmt"
	"math/rand"
	"time"
)

func main() {

	generatorConfig := &datageneratorC.GeneratorConfigS{
		FvalueGeo: 0.5,
		Minimum:   1,
		Maximum:   10,
		// Maximum:   40,
		DistributionType: []distributions.GeneratorTypeEnum{
			distributions.Uni,
			distributions.Geo,
			distributions.Hrm,
			distributions.Dhr,
		},
	}

	generators := datagenerator.CreateDataGenerator(*generatorConfig)

	algs := []bigpackingalgs.BigPackingAlg{
		bigpackingalgs.BF,
		bigpackingalgs.FF,
		bigpackingalgs.NF,
		bigpackingalgs.WF,
	}

	randomSolver := *bigpackingsolver.BigPackingSolver_Create(bigpackingalgs.RF, false)
	solvers := make([]bigpackingsolver.BigPackingSolver, len(algs))

	for i, alg := range algs {
		solvers[i] = *bigpackingsolver.BigPackingSolver_Create(alg, false)
	}

	results := make([][]float64, len(generators))
	for i := range results {
		results[i] = make([]float64, len(solvers)+1)
	}

	trials := 100_000

	for i := 0; i < trials; i++ {
		for j, generator := range generators {
			seq := getSequence(generator)
			opt := sumSeq(seq)
			for _, solver := range solvers {
				for _, elem := range seq {
					solver.Serve(elem)
				}
			}

			temporaryRandRes := 0
			randRepeats := 100
			for r := 0; r < randRepeats; r++ {
				randomSolver.Clear()
				for _, elem := range seq {
					randomSolver.Serve(elem)
				}
				_, res := randomSolver.Raport()
				temporaryRandRes += res
			}

			results[j][4] += float64(temporaryRandRes/randRepeats) / float64(opt)

			for k, solver := range solvers {
				_, res := solver.Raport()
				results[j][k] += float64(res) / float64(opt)
			}
			for _, solver := range solvers {
				solver.Clear()
			}
		}

		// fmt.Println("done iteration ", i)
	}

	for j := range generators {
		fmt.Println(distributions.FromInt(j))
		for k := range solvers {
			fmt.Println("\t", algs[k], " ", float64(results[j][k])/float64(trials))
		}
		fmt.Println("\t", bigpackingalgs.RF, " ", float64(results[j][4])/float64(trials))

	}
}

func getSequence(dG datagenerator.GenericDataGenerator) []float64 {
	valGen := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	sequence := make([]float64, 100)
	iterator := 0

	for iterator < 100 {
		value := valGen.Float64()
		quantity := dG.GetRequest()

		for i := 0; i < quantity; i++ {
			sequence[iterator] = value
			iterator++

			if iterator == 100 {
				break
			}
		}

	}

	return sequence
}

func sumSeq(sequence []float64) int {
	sum := 0.0

	for _, elem := range sequence {
		sum += elem
	}

	return int(sum)

}
