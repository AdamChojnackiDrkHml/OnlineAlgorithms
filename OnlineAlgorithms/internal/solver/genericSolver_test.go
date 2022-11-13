package solver

import (
	"OnlineAlgorithms/internal/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {

	solverConfs := utils.SolverConfigS{ProblemType: 0, Size: 10, Alg: 1, Debug: true, DoAll: false}

	_, err := CreateSolver(solverConfs)

	assert.Equal(t, nil, err)

	solverConfs = utils.SolverConfigS{ProblemType: 0, Size: 10, Alg: 7, Debug: true, DoAll: false}

	_, err = CreateSolver(solverConfs)

	assert.NotEqual(t, nil, err)

}
