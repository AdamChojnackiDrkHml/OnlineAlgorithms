package main

import (
	datagenerator "OnlineAlgorithms/pkg/dataGenerator"
	datageneratorC "OnlineAlgorithms/pkg/dataGenerator/dataGeneratorConfigs"
	"OnlineAlgorithms/pkg/dataGenerator/distributions"
	"OnlineAlgorithms/pkg/graphs"
	pagemigrationsolver "OnlineAlgorithms/pkg/solver/pageMigrationSolver"
	pagemigrationalgs "OnlineAlgorithms/pkg/solver/pageMigrationSolver/pageMigrationAlgs"
	"fmt"
)

func main() {

	generatorConfig := &datageneratorC.GeneratorConfigS{
		Minimum: 0,
		Maximum: 63,
		DistributionType: []distributions.GeneratorTypeEnum{
			distributions.Uni,
			distributions.Geo,
			distributions.Hrm,
			distributions.Dhr,
		},
	}

	generators := datagenerator.CreateDataGenerator(*generatorConfig)

	algs := []pagemigrationalgs.PageMigrationAlgs{
		pagemigrationalgs.MTM,
		pagemigrationalgs.F,
	}

	solversHypercube := make([]pagemigrationsolver.PageMigrationSolver, len(algs))
	solversTorus := make([]pagemigrationsolver.PageMigrationSolver, len(algs))

	graphs := []graphs.Graph{graphs.Torus_newTorus(), graphs.Hypercube_newHypercube()}

	for i, alg := range algs {
		solversTorus[i] = *pagemigrationsolver.PageMigrationSolver_Create(alg, false, &graphs[0])
		solversHypercube[i] = *pagemigrationsolver.PageMigrationSolver_Create(alg, false, &graphs[1])
	}

	results := make([][][2]float64, len(generators))
	for i := range results {
		results[i] = make([][2]float64, len(solversTorus))
	}

	trials := 1_000_000

	for i := 0; i < trials; i++ {
		for j, generator := range generators {
			seq := datagenerator.GetSliceOfRequests(&generator, 1024)

			for i := range solversTorus {
				for _, elem := range seq {
					solversTorus[i].Serve(uint8(elem))
					solversHypercube[i].Serve(uint8(elem))
				}
			}

			for k := range solversTorus {
				_, resT := solversTorus[k].Raport()
				_, resH := solversHypercube[k].Raport()
				results[j][k][0] += float64(resT) / 1024.0
				results[j][k][1] += float64(resH) / 1024.0
			}

			for k := range solversTorus {
				solversTorus[k].Clear()
				solversHypercube[k].Clear()
			}

		}

		// fmt.Println("done iteration ", i)
	}

	for j := range generators {

		fmt.Println(distributions.FromInt(j))

		fmt.Println("\t", "Torus")
		fmt.Println("\t\t", algs[0], " ", float64(results[j][0][0])/float64(trials))
		fmt.Println("\t\t", algs[1], " ", float64(results[j][1][0])/float64(trials))

		fmt.Println("\t", "Hypercube")
		fmt.Println("\t\t", algs[0], " ", float64(results[j][0][1])/float64(trials))
		fmt.Println("\t\t", algs[1], " ", float64(results[j][1][1])/float64(trials))
	}
}
