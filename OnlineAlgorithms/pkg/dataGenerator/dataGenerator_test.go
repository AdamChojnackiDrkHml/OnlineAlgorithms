package datagenerator

import (
	"fmt"
	"testing"
)

func TestUniDist(t *testing.T) {
	g := UNI_Create(0, 100)
	counter := make([]int, 101)
	for i := 0; i < 10000; i++ {
		counter[g.GetRequest()]++

	}

	fmt.Println(counter)
	fmt.Println("aa")
}

func TestGeoDist(t *testing.T) {
	g := GEO_Create(0.2, 100)
	counter := make([]int, 101)
	for i := 0; i < 10000; i++ {
		counter[g.GetRequest()]++

	}

	fmt.Println(counter)
	fmt.Println("aa")
}

func TestPoisDist(t *testing.T) {
	g := POIS_Create(10, 100)
	counter := make([]int, 101)
	for i := 0; i < 10000; i++ {
		counter[g.GetRequest()]++

	}

	fmt.Println(counter)
	fmt.Println("aa")
}

func TestHarmonicDist(t *testing.T) {
	g := HRM_Create(0, 10)
	counter := make([]int, 101)
	for i := 0; i < 10000; i++ {
		counter[g.GetRequest()]++

	}

	fmt.Println(counter)
	fmt.Println("aa")
}

func TestDiharmonicDist(t *testing.T) {
	g := DHR_Create(0, 10)
	counter := make([]int, 101)
	for i := 0; i < 10000; i++ {
		counter[g.GetRequest()]++

	}

	fmt.Println(counter)
	fmt.Println("aa")
}
