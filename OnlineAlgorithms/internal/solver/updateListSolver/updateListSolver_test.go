package updatelistsolver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const debug = true

func TestFQ(t *testing.T) {

	fq := FQAlg_Create(10, debug)
	ulS := UpdateListSolver{alg: fq}

	initRequests := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, n := range initRequests {
		ulS.Serve(n)
	}

	assert.Equal(t, 45, ulS.cost, "Unexpected fault number")
	assert.Equal(t, [][]int{{9, 1}, {8, 1}, {7, 1}, {6, 1}, {5, 1}, {4, 1}, {3, 1}, {2, 1}, {1, 1}, {0, 1}}, fq.unpackMemory(), "NO dupa1")

	ulS.Serve(9)
	ulS.Serve(9)
	ulS.Serve(2)
	ulS.Serve(2)
	ulS.Serve(2)
	ulS.Serve(1)
	ulS.Serve(8)

	assert.Equal(t, 64, ulS.cost, "Unexpected fault number")
	assert.Equal(t, [][]int{{2, 4}, {9, 3}, {8, 2}, {1, 2}, {7, 1}, {6, 1}, {5, 1}, {4, 1}, {3, 1}, {0, 1}}, fq.unpackMemory(), "NO dupa1")

}

func TestBIT(t *testing.T) {

	bit := BITAlg_Create(10, debug)

	for i := 0; i < 10; i++ {
		bit.memory[i].bit = true
	}

	ulS := UpdateListSolver{alg: bit}

	initRequests := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, n := range initRequests {
		ulS.Serve(n)
	}

	assert.Equal(t, 45, ulS.cost, "Unexpected fault number")
	assert.Equal(t, [][]int{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {5, 0}, {6, 0}, {7, 0}, {8, 0}, {9, 0}}, bit.unpackMemory(), "NO dupa1")

	ulS.Serve(9)

	assert.Equal(t, 54, ulS.cost, "Unexpected fault number")
	assert.Equal(t, [][]int{{9, 1}, {0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {5, 0}, {6, 0}, {7, 0}, {8, 0}}, bit.unpackMemory(), "NO dupa1")

	ulS.Serve(8)
	ulS.Serve(9)

	assert.Equal(t, 64, ulS.cost, "Unexpected fault number")
	assert.Equal(t, [][]int{{8, 1}, {9, 0}, {0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {5, 0}, {6, 0}, {7, 0}}, bit.unpackMemory(), "NO dupa1")

}

func TestMTF(t *testing.T) {

	mtf := MTFAlg_Create(10, debug)

	ulS := UpdateListSolver{alg: mtf}

	initRequests := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, n := range initRequests {
		ulS.Serve(n)
	}

	assert.Equal(t, 45, ulS.cost, "Unexpected fault number")
	assert.Equal(t, []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, mtf.memory, "NO dupa1")

	ulS.Serve(1)

	assert.Equal(t, 53, ulS.cost, "Unexpected fault number")
	assert.Equal(t, []int{1, 9, 8, 7, 6, 5, 4, 3, 2, 0}, mtf.memory, "NO dupa1")
}

func TestTrans(t *testing.T) {

	trans := TransAlg_Create(10, debug)

	ulS := UpdateListSolver{alg: trans}

	initRequests := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, n := range initRequests {
		ulS.Serve(n)
	}

	assert.Equal(t, 45, ulS.cost, "Unexpected fault number")
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, trans.memory, "NO dupa1")

	ulS.Serve(1)

	assert.Equal(t, 45, ulS.cost, "Unexpected fault number")
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, trans.memory, "NO dupa1")

	ulS.Serve(9)

	assert.Equal(t, 53, ulS.cost, "Unexpected fault number")
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 9, 8, 0}, trans.memory, "NO dupa1")
}

func TestTS(t *testing.T) {

	ts := TSAlg_Create(10, debug)

	ulS := UpdateListSolver{alg: ts}

	var memoryHelper = func() []int {
		tsMem := make([]int, ts.size)

		for i, n := range ts.memory {
			tsMem[i] = n.mem
		}

		return tsMem
	}

	initRequests := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, n := range initRequests {
		ulS.Serve(n)
	}

	assert.Equal(t, 45, ulS.cost, "Unexpected fault number")
	assert.Equal(t, []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, memoryHelper(), "NO dupa1")

	ulS.Serve(0)

	assert.Equal(t, 54, ulS.cost, "Unexpected fault number")
	assert.Equal(t, []int{0, 9, 8, 7, 6, 5, 4, 3, 2, 1}, memoryHelper(), "NO dupa1")

	ulS.Serve(0)
	ulS.Serve(9)
	ulS.Serve(9)

	assert.Equal(t, 56, ulS.cost, "Unexpected fault number")
	assert.Equal(t, []int{9, 0, 8, 7, 6, 5, 4, 3, 2, 1}, memoryHelper(), "NO dupa1")

	ulS.Serve(1)

	assert.Equal(t, 65, ulS.cost, "Unexpected fault number")
	assert.Equal(t, []int{9, 0, 1, 8, 7, 6, 5, 4, 3, 2}, memoryHelper(), "NO dupa1")
}
