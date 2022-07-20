package main

import (
	dataGenerator "OnlineAlgorithms/internal/dataGenerator"
	uniDistGenerator "OnlineAlgorithms/internal/dataGenerator/uniDistGenerator"
	"OnlineAlgorithms/internal/solver"
	pagingsolver "OnlineAlgorithms/internal/solver/pagingSolver"
	"fmt"
)

func main() {
	var pS solver.GenericSolver
	pS = pagingsolver.PagingSolver_Create(5, 0)

	var dG dataGenerator.GenericDataGenerator
	dG = uniDistGenerator.Create(0, 7)

	for i := 0; i < 15; i++ {
		pS.Serve(dG.GetRequest())
	}

	fmt.Println(pS.Raport())
}
