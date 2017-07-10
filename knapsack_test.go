package main

import (
	"testing"
)

var cap1 = 20.
var testItems1 = []item{
	{name: "clock", value: 175, weight: 10},
	{name: "painting", value: 90, weight: 9},
	{name: "radio", value: 20, weight: 4},
	{name: "vase", value: 50, weight: 2},
	{name: "book", value: 10, weight: 1},
	{name: "computer", value: 200, weight: 20},
}

var cap2 = 26.
var testItems2 = []item{
	{"item1", 24., 12.},
	{"item2", 13., 7.},
	{"item3", 23., 11.},
	{"item4", 15., 8.},
	{"item5", 16., 9.},
}

func TestGreedyDensity6(t *testing.T) {

	f := func(i, j int) bool {
		return testItems1[i].value/testItems1[i].weight > testItems1[j].value/testItems1[j].weight
	}

	v := 0.
	for _, i := range greedy(testItems1, cap1, f) {
		v += i.value
	}

	if v != 255.0 {
		t.Error("Expected 255, got ", v)
	}
}

func TestGreedyDensity5(t *testing.T) {

	f := func(i, j int) bool {
		return testItems2[i].value/testItems2[i].weight > testItems2[j].value/testItems2[j].weight
	}

	v := 0.
	for _, i := range greedy(testItems2, cap2, f) {
		v += i.value
	}

	if v != 47.0 {
		t.Error("Expected 47, got ", v)
	}
}

func TestOptimal(t *testing.T) {

	v, _ := bestSolution(testItems1, cap1)

	if v != 275.0 {
		t.Error("Expected 275, got ", v)
	}
}

func benchmarkDensity(b *testing.B, set []item, cap float64) {

	f := func(i, j int) bool {
		return testItems1[i].value/testItems1[i].weight > testItems1[j].value/testItems1[j].weight
	}

	for n := 0; n < b.N; n++ {
		greedy(set, cap, f)
	}
}

func BenchmarkDensity5(b *testing.B) { benchmarkDensity(b, testItems2, cap2) }
func BenchmarkDensity6(b *testing.B) { benchmarkDensity(b, testItems1, cap1) }
