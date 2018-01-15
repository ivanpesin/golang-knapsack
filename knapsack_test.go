package knapsack

import (
	"testing"
)

var cap1 = 20.
var testItems1 = []Item{
	{Name: "clock", Value: 175, Weight: 10},
	{Name: "painting", Value: 90, Weight: 9},
	{Name: "radio", Value: 20, Weight: 4},
	{Name: "vase", Value: 50, Weight: 2},
	{Name: "book", Value: 10, Weight: 1},
	{Name: "computer", Value: 200, Weight: 20},
}

var cap2 = 26.
var testItems2 = []Item{
	{"item1", 24., 12.},
	{"item2", 13., 7.},
	{"item3", 23., 11.},
	{"item4", 15., 8.},
	{"item5", 16., 9.},
}

func TestGreedyDensity6(t *testing.T) {

	f := func(i, j int) bool {
		return testItems1[i].Value/testItems1[i].Weight > testItems1[j].Value/testItems1[j].Weight
	}

	v := 0.
	for _, i := range Greedy(testItems1, cap1, f) {
		v += i.Value
	}

	if v != 255.0 {
		t.Error("Expected 255, got ", v)
	}
}

func TestGreedyDensity5(t *testing.T) {

	f := func(i, j int) bool {
		return testItems2[i].Value/testItems2[i].Weight > testItems2[j].Value/testItems2[j].Weight
	}

	v := 0.
	for _, i := range Greedy(testItems2, cap2, f) {
		v += i.Value
	}

	if v != 47.0 {
		t.Error("Expected 47, got ", v)
	}
}

func TestOptimal(t *testing.T) {

	v, _ := BestSolution(testItems1, cap1)

	if v != 275.0 {
		t.Error("Expected 275, got ", v)
	}
}

func benchmarkDensity(b *testing.B, set []Item, cap float64) {

	f := func(i, j int) bool {
		return testItems1[i].Value/testItems1[i].Weight > testItems1[j].Value/testItems1[j].Weight
	}

	for n := 0; n < b.N; n++ {
		Greedy(set, cap, f)
	}
}

func BenchmarkDensity5(b *testing.B) { benchmarkDensity(b, testItems2, cap2) }
func BenchmarkDensity6(b *testing.B) { benchmarkDensity(b, testItems1, cap1) }
