package main

import (
	"fmt"
	"sort"
)

type item struct {
	name   string
	value  float64
	weight float64
}

const knapsackSize = 20

var store = []item{
	{name: "clock", value: 175, weight: 10},
	{name: "painting", value: 90, weight: 9},
	{name: "radio", value: 20, weight: 4},
	{name: "vase", value: 50, weight: 2},
	{name: "book", value: 10, weight: 1},
	{name: "computer", value: 200, weight: 20},
}

func (i item) String() string {
	return fmt.Sprintf("<%10s| $%6.2f, %6.2f kg>", i.name, i.value, i.weight)
}

func greedy(items []item, maxWeight float64, metric func(i, j int) bool) (r []item) {

	sort.Slice(items, metric)
	//fmt.Println("D: ", items)
	w := 0.

	for _, i := range items {
		if w+i.weight <= maxWeight {
			r = append(r, i)
			w += i.weight
		}
	}

	return

}

func combinations(items []item) {
	sort.Slice(items, func(i, j int) bool { return items[i].name < items[i].name })
}

func main() {

	fmt.Println("List of goods in store: ")
	for _, v := range store {
		fmt.Println("   ", v)
	}

	items := make([]item, len(store))
	copy(items, store)

	funcs := make([]func(i, j int) bool, 3)
	funcs[0] = func(i, j int) bool { return items[i].value > items[j].value }
	funcs[1] = func(i, j int) bool { return store[i].weight < store[j].weight }
	funcs[2] = func(i, j int) bool {
		return store[i].value/store[i].weight > store[j].value/store[j].weight
	}

	funcNames := make([]string, 3)
	funcNames[0] = "value"
	funcNames[1] = "weight"
	funcNames[2] = "density"

	for n, fname := range funcNames {
		v := 0.
		fmt.Printf("Being greedy based on %s: \n", fname)
		for _, i := range greedy(items, knapsackSize, funcs[n]) {
			fmt.Println("   ", i)
			v += i.value
		}
		fmt.Printf("Total value: $%.2f\n", v)
	}

}