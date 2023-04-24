package pagingsolveralgs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const debug = false

func TestFifo(t *testing.T) {
	fifo := FIFOAlg_Create(10, debug)
	faults := 0
	initRequests := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, n := range initRequests {
		if !fifo.UpdateMemory(n) {
			faults++
		}
	}

	if !fifo.UpdateMemory(8) {
		faults++
	}

	assert.Equal(t, faults, 10, "Unexpected fault number")
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, fifo.memory)

	if !fifo.UpdateMemory(10) {
		faults++
	}

	assert.Equal(t, faults, 11, "Unexpected fault number")
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, fifo.memory)

}

func TestLru(t *testing.T) {
	lru := LRUAlg_Create(10, debug)

	faults := 0
	initRequests := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, n := range initRequests {
		if !lru.UpdateMemory(n) {
			faults++
		}
	}

	if !lru.UpdateMemory(8) {
		faults++
	}

	assert.Equal(t, faults, 10, "Unexpected fault number")
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 9, 8}, lru.memory, "NO dupa1")

	if !lru.UpdateMemory(0) {
		faults++
	}
	if !lru.UpdateMemory(10) {
		faults++
	}

	assert.Equal(t, faults, 11, "Unexpected fault number")
	assert.Equal(t, []int{2, 3, 4, 5, 6, 7, 9, 8, 0, 10}, lru.memory, "NO dupa2")

}

func TestLfu(t *testing.T) {
	lfu := LFUAlg_Create(10, debug)
	faults := 0

	initRequests := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, n := range initRequests {
		if !lfu.UpdateMemory(n) {
			faults++
		}
	}

	if !lfu.UpdateMemory(8) {
		faults++
	}

	assert.Equal(t, faults, 10, "Unexpected fault number")
	assert.Equal(t, [][2]int{
		{8, 2},
		{9, 1},
		{7, 1},
		{6, 1},
		{5, 1},
		{4, 1},
		{3, 1},
		{2, 1},
		{1, 1},
		{0, 1},
	}, lfu.unpackMemory(), "NO dupa1")

	for i := 0; i < 9; i++ {
		if !lfu.UpdateMemory(initRequests[i]) {
			faults++
		}
	}

	if !lfu.UpdateMemory(9) {
		faults++
	}
	if !lfu.UpdateMemory(9) {
		faults++
	}
	if !lfu.UpdateMemory(9) {
		faults++
	}

	if !lfu.UpdateMemory(10) {
		faults++
	}

	if !lfu.UpdateMemory(10) {
		faults++
	}
	assert.Equal(t, faults, 11, "Unexpected fault number")
	assert.Equal(t, [][2]int{
		{9, 4},
		{8, 3},
		{10, 2},
		{7, 2},
		{6, 2},
		{5, 2},
		{4, 2},
		{3, 2},
		{2, 2},
		{1, 2},
	}, lfu.unpackMemory(), "NO dupa2")

	if !lfu.UpdateMemory(0) {
		faults++
	}

	assert.Equal(t, faults, 12, "Unexpected fault number")
	assert.Equal(t, [][2]int{
		{9, 4},
		{0, 3},
		{8, 3},
		{10, 2},
		{7, 2},
		{6, 2},
		{5, 2},
		{4, 2},
		{3, 2},
		{2, 2},
	}, lfu.unpackMemory(), "NO dupa2")

}

func TestMark1(t *testing.T) {
	marklru := MARKLRUAlg_Create(10, debug)
	faults := 0

	initRequests := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, n := range initRequests {
		if !marklru.UpdateMemory(n) {
			faults++
		}
	}

	if !marklru.UpdateMemory(8) {
		faults++
	}

	assert.Equal(t, faults, 10, "Unexpected fault number")
	assert.Equal(t, []int{8, 9, 7, 6, 5, 4, 3, 2, 1, 0}, marklru.memory)

	if !marklru.UpdateMemory(10) {
		faults++
	}

	assert.Equal(t, faults, 11, "Unexpected fault number")
	assert.Equal(t, []int{10, 8, 9, 7, 6, 5, 4, 3, 2, 1}, marklru.memory)
}

func TestMark2(t *testing.T) {
	markfc := MARKFCAlg_Create(10, debug)

	faults := 0

	initRequests := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, n := range initRequests {
		if !markfc.UpdateMemory(n) {
			faults++
		}
	}

	if !markfc.UpdateMemory(7) {
		faults++
	}

	assert.Equal(t, faults, 10, "Unexpected fault number")
	assert.Equal(t, []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, markfc.memory)

	//INCREASE FQ OF ALL (EXCEPT 9 AND 8) BY 3
	for i := 0; i < 8; i++ {
		if !markfc.UpdateMemory(i) {
			faults++
		}
	}

	for i := 0; i < 8; i++ {
		if !markfc.UpdateMemory(i) {
			faults++
		}
	}

	for i := 0; i < 8; i++ {
		if !markfc.UpdateMemory(i) {
			faults++
		}
	}

	for i := 0; i < 8; i++ {
		if !markfc.UpdateMemory(i) {
			faults++
		}
	}

	//MARK 9
	if !markfc.UpdateMemory(9) {
		faults++
	}
	//MARK 8 (and set it to have fq > fq(9) and fq 2 less than others)
	if !markfc.UpdateMemory(8) {
		faults++
	}
	if !markfc.UpdateMemory(8) {
		faults++
	}

	//9 should be evicted and all unmarked
	if !markfc.UpdateMemory(10) {
		faults++
	}

	assert.Equal(t, faults, 11, "Unexpected fault number")
	assert.Equal(t, []int{10, 8, 7, 6, 5, 4, 3, 2, 1, 0}, markfc.memory)

	//MARK 8
	if !markfc.UpdateMemory(8) {
		faults++
	}

	if !markfc.UpdateMemory(11) {
		faults++
	}

	assert.Equal(t, faults, 12, "Unexpected fault number")

	//8 SHOULD NOT BE EVICTED
	assert.Equal(t, []int{11, 10, 8, 7, 5, 4, 3, 2, 1, 0}, markfc.memory)

}
