package updatelistsolveralgs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const debug = true

func TestFQ(t *testing.T) {

	alg := FCAlg_Create(10, debug)
	cost := 0

	initRequests := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, n := range initRequests {
		cost += alg.UpdateList(n)
	}

	assert.Equal(t, 45, cost, "Unexpected fault number")
	assert.Equal(t, [][]int{{9, 1}, {8, 1}, {7, 1}, {6, 1}, {5, 1}, {4, 1}, {3, 1}, {2, 1}, {1, 1}, {0, 1}}, alg.unpackMemory(), "NO dupa1")

	cost += alg.UpdateList(9)
	cost += alg.UpdateList(9)
	cost += alg.UpdateList(2)
	cost += alg.UpdateList(2)
	cost += alg.UpdateList(2)
	cost += alg.UpdateList(1)
	cost += alg.UpdateList(8)

	assert.Equal(t, 64, cost, "Unexpected fault number")
	assert.Equal(t, [][]int{{2, 4}, {9, 3}, {8, 2}, {1, 2}, {7, 1}, {6, 1}, {5, 1}, {4, 1}, {3, 1}, {0, 1}}, alg.unpackMemory(), "NO dupa1")

}

func TestBIT(t *testing.T) {

	alg := BITAlg_Create(10, debug)

	for i := 0; i < 10; i++ {
		alg.memory[i].bit = true
	}

	cost := 0

	initRequests := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, n := range initRequests {
		cost += alg.UpdateList(n)
	}

	assert.Equal(t, 45, cost, "Unexpected fault number")
	assert.Equal(t, [][]int{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {5, 0}, {6, 0}, {7, 0}, {8, 0}, {9, 0}}, alg.unpackMemory(), "NO dupa1")

	cost += alg.UpdateList(9)

	assert.Equal(t, 54, cost, "Unexpected fault number")
	assert.Equal(t, [][]int{{9, 1}, {0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {5, 0}, {6, 0}, {7, 0}, {8, 0}}, alg.unpackMemory(), "NO dupa1")

	cost += alg.UpdateList(8)
	cost += alg.UpdateList(9)

	assert.Equal(t, 64, cost, "Unexpected fault number")
	assert.Equal(t, [][]int{{8, 1}, {9, 0}, {0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {5, 0}, {6, 0}, {7, 0}}, alg.unpackMemory(), "NO dupa1")

}

func TestMTF(t *testing.T) {

	alg := MTFAlg_Create(10, debug)

	cost := 0

	initRequests := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, n := range initRequests {
		cost += alg.UpdateList(n)
	}

	assert.Equal(t, 45, cost, "Unexpected fault number")
	assert.Equal(t, []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, alg.memory, "NO dupa1")

	cost += alg.UpdateList(1)

	assert.Equal(t, 53, cost, "Unexpected fault number")
	assert.Equal(t, []int{1, 9, 8, 7, 6, 5, 4, 3, 2, 0}, alg.memory, "NO dupa1")
}

func TestTrans(t *testing.T) {

	alg := TransAlg_Create(10, debug)

	cost := 0

	initRequests := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, n := range initRequests {
		cost += alg.UpdateList(n)
	}

	assert.Equal(t, 45, cost, "Unexpected fault number")
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, alg.memory, "NO dupa1")

	cost += alg.UpdateList(1)

	assert.Equal(t, 45, cost, "Unexpected fault number")
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, alg.memory, "NO dupa1")

	cost += alg.UpdateList(9)

	assert.Equal(t, 53, cost, "Unexpected fault number")
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 9, 8, 0}, alg.memory, "NO dupa1")
}

func TestTS(t *testing.T) {

	alg := TSAlg_Create(10, debug)

	cost := 0

	var memoryHelper = func() []int {
		tsMem := make([]int, alg.size)

		for i, n := range alg.memory {
			tsMem[i] = n.mem
		}

		return tsMem
	}

	initRequests := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, n := range initRequests {
		cost += alg.UpdateList(n)
	}

	assert.Equal(t, 45, cost, "Unexpected fault number")
	assert.Equal(t, []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, memoryHelper(), "NO dupa1")

	cost += alg.UpdateList(0)

	assert.Equal(t, 54, cost, "Unexpected fault number")
	assert.Equal(t, []int{0, 9, 8, 7, 6, 5, 4, 3, 2, 1}, memoryHelper(), "NO dupa1")

	cost += alg.UpdateList(0)
	cost += alg.UpdateList(9)
	cost += alg.UpdateList(9)

	assert.Equal(t, 56, cost, "Unexpected fault number")
	assert.Equal(t, []int{9, 0, 8, 7, 6, 5, 4, 3, 2, 1}, memoryHelper(), "NO dupa1")

	cost += alg.UpdateList(1)

	assert.Equal(t, 65, cost, "Unexpected fault number")
	assert.Equal(t, []int{9, 0, 1, 8, 7, 6, 5, 4, 3, 2}, memoryHelper(), "NO dupa1")
}
