package pagingsolver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const debug = false

func TestFifo(t *testing.T) {
	fifo := FIFOAlg_Create(10, debug)
	pS := PagingSolver{alg: fifo}

	initRequests := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, n := range initRequests {
		pS.Serve(n)
	}

	pS.Serve(8)

	assert.Equal(t, pS.faults, 10, "Unexpected fault number")
	assert.Equal(t, []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, fifo.memory)

	pS.Serve(10)

	assert.Equal(t, pS.faults, 11, "Unexpected fault number")
	assert.Equal(t, []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, fifo.memory)

}

func TestLru(t *testing.T) {
	lru := LRUAlg_Create(10, debug)
	pS := PagingSolver{alg: lru}

	initRequests := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, n := range initRequests {
		pS.Serve(n)
	}

	pS.Serve(8)

	assert.Equal(t, pS.faults, 10, "Unexpected fault number")
	assert.Equal(t, [][]int{{0, 10}, {1, 9}, {2, 8}, {3, 7}, {4, 6}, {5, 5}, {6, 4}, {7, 3}, {8, 0}, {9, 1}}, lru.unpackMemory(), "NO dupa1")

	pS.Serve(0)
	pS.Serve(10)

	assert.Equal(t, pS.faults, 11, "Unexpected fault number")
	assert.Equal(t, [][]int{{2, 10}, {3, 9}, {5, 7}, {7, 5}, {4, 8}, {9, 3}, {6, 6}, {0, 1}, {8, 2}, {10, 0}}, lru.unpackMemory(), "NO dupa2")

}

func TestLfu(t *testing.T) {
	lfu := LFUAlg_Create(10, debug)
	pS := PagingSolver{alg: lfu}

	initRequests := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, n := range initRequests {
		pS.Serve(n)
	}

	pS.Serve(8)

	assert.Equal(t, pS.faults, 10, "Unexpected fault number")
	assert.Equal(t, [][]int{{0, 1}, {1, 1}, {2, 1}, {3, 1}, {4, 1}, {5, 1}, {6, 1}, {7, 1}, {8, 2}, {9, 1}}, lfu.unpackMemory(), "NO dupa1")

	for i := 0; i < 9; i++ {
		pS.Serve(initRequests[i])
	}

	pS.Serve(10)

	assert.Equal(t, pS.faults, 11, "Unexpected fault number")
	assert.Equal(t, [][]int{{10, 1}, {3, 2}, {6, 2}, {1, 2}, {7, 2}, {2, 2}, {5, 2}, {0, 2}, {8, 3}, {4, 2}}, lfu.unpackMemory(), "NO dupa2")

}

func TestMark1(t *testing.T) {
	mark := MARKLRUAlg_Create(10, debug)
	pS := PagingSolver{alg: mark}

	initRequests := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, n := range initRequests {
		pS.Serve(n)
	}

	pS.Serve(8)

	assert.Equal(t, pS.faults, 10, "Unexpected fault number")
	assert.Equal(t, []int{8, 9, 7, 6, 5, 4, 3, 2, 1, 0}, mark.memory)

	pS.Serve(10)

	assert.Equal(t, pS.faults, 11, "Unexpected fault number")
	assert.Equal(t, []int{10, 8, 9, 7, 6, 5, 4, 3, 2, 1}, mark.memory)
}

func TestMark2(t *testing.T) {
	mark := MARKFCAlg_Create(10, debug)
	pS := PagingSolver{alg: mark}

	initRequests := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, n := range initRequests {
		pS.Serve(n)
	}

	pS.Serve(7)

	assert.Equal(t, pS.faults, 10, "Unexpected fault number")
	assert.Equal(t, []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, mark.memory)

	//INCREASE FQ OF ALL (EXCEPT 9 AND 8) BY 3
	for i := 0; i < 8; i++ {
		pS.Serve(i)
	}

	for i := 0; i < 8; i++ {
		pS.Serve(i)
	}

	for i := 0; i < 8; i++ {
		pS.Serve(i)
	}

	for i := 0; i < 8; i++ {
		pS.Serve(i)
	}

	//MARK 9
	pS.Serve(9)
	//MARK 8 (and set it to have fq > fq(9) and fq 2 less than others)
	pS.Serve(8)
	pS.Serve(8)

	//9 should be evicted and all unmarked
	pS.Serve(10)

	assert.Equal(t, pS.faults, 11, "Unexpected fault number")
	assert.Equal(t, []int{10, 8, 7, 6, 5, 4, 3, 2, 1, 0}, mark.memory)

	//MARK 8
	pS.Serve(8)

	pS.Serve(11)

	assert.Equal(t, pS.faults, 12, "Unexpected fault number")

	//8 SHOULD NOT BE EVICTED
	assert.Equal(t, []int{11, 10, 8, 7, 5, 4, 3, 2, 1, 0}, mark.memory)

}
