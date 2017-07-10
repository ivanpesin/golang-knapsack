package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
)

type item struct {
	name   string
	value  float64
	weight float64
}

var knapsackCapacity = -1.
var store = []item{}

func (i item) String() string {
	return fmt.Sprintf(" %-10s $%10.2f %10.2f kg", i.name, i.value, i.weight)
}

// greedy implements greedy knapsack alg
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

// combinations returns all possible combinations of items in store.
// Possible combinations are sent to a channel to avoid large memory consumption
func combinations(items []item, ch chan []item) {
	defer close(ch)

	p := int(math.Pow(2., float64(len(items))))

	for i := 0; i < p; i++ {
		set := []item{}
		for j := 0; j < len(items); j++ {
			if (i>>uint(j))&1 == 1 {
				set = append(set, items[j])
			}
		}
		ch <- set
	}
}

// getSackWeight returns weight of a given set of items
func getSackWeight(set []item) (r float64) {
	for _, i := range set {
		r += i.weight
	}
	return
}

// getSackValue returns value of a given set
func getSackValue(set []item) (r float64) {
	for _, i := range set {
		r += i.value
	}
	return
}

// bestSolution looks through all possible combinations of items
// and selects the one with highest value which is below or eq target weight
func bestSolution(items []item, maxWeight float64) (float64, []item) {
	bestVal := 0.
	bestSack := []item{}

	ch := make(chan []item)
	go combinations(items, ch)

	for sack := range ch {
		if getSackWeight(sack) <= maxWeight {
			v := getSackValue(sack)
			if v > bestVal {
				bestVal = v
				bestSack = sack
			}
		}
	}
	return bestVal, bestSack
}

// readStore reads items and their properties from a file
func readStore(fn string) {
	f, err := os.Open(fn)
	if err != nil {
		fmt.Printf("ERROR: Unable to open file: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Reading store from file: %s\n", fn)
	defer f.Close()
	store = []item{}
	s := bufio.NewScanner(f)
	for s.Scan() {
		l := strings.TrimSpace(s.Text())
		if len(l) > 0 && (l[0] == ';' || l[0] == '#') {
			continue
		}
		fields := strings.Fields(l)
		if knapsackCapacity == -1 && len(fields) == 1 {
			val, err := strconv.ParseFloat(fields[0], 64)
			if err != nil {
				fmt.Printf("ERROR: Unable to parse value in:\n>>>   %v\n", l)
				continue
			}
			knapsackCapacity = val
			continue
		}
		if len(fields) != 3 {
			fmt.Printf("ERROR: Invalid number of fields, must be 3:\n>>>   %v\n", l)
			continue
		}
		val, err := strconv.ParseFloat(fields[1], 64)
		if err != nil {
			fmt.Printf("ERROR: Unable to parse value in:\n>>>   %v\n", l)
			continue
		}
		weight, err := strconv.ParseFloat(fields[2], 64)
		if err != nil {
			fmt.Printf("ERROR: Unable to parse weight in:\n>>>   %v\n", l)
			continue
		}

		//fmt.Printf("D: Appending: %s %f %f\n", fields[0], val, weight)
		store = append(store, item{fields[0], val, weight})
	}

	// check we got knapsack capacity from input file
	if knapsackCapacity < 0 {
		fmt.Printf("ERROR: Knapsack capacity unspecified, probably misformed input file\n")
		os.Exit(1)
	}

	// check we have at least 1 item in store
	if len(store) < 1 {
		fmt.Printf("ERROR: Empty store, probably misformed input file\n")
		os.Exit(1)
	}

}

func main() {

	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <input_file>\n", path.Base(os.Args[0]))
		os.Exit(0)
	}
	readStore(os.Args[1])

	fmt.Println("List of goods in store: ")
	for _, v := range store {
		fmt.Println("   ", v)
	}

	items := make([]item, len(store))
	copy(items, store)

	funcs := make([]func(i, j int) bool, 3)
	funcs[0] = func(i, j int) bool { return items[i].value > items[j].value }
	funcs[1] = func(i, j int) bool { return items[i].weight < items[j].weight }
	funcs[2] = func(i, j int) bool {
		return items[i].value/items[i].weight > items[j].value/items[j].weight
	}

	funcNames := make([]string, 3)
	funcNames[0] = "value"
	funcNames[1] = "weight"
	funcNames[2] = "density"

	for n, fname := range funcNames {
		v := 0.
		fmt.Printf("Being greedy based on %s: \n", fname)
		for _, i := range greedy(items, knapsackCapacity, funcs[n]) {
			fmt.Println("   ", i)
			v += i.value
		}
		fmt.Printf("--- Total value: $%.2f\n", v)
	}

	sort.Slice(items, func(i, j int) bool { return items[i].name < items[j].name })
	fmt.Printf("\nOptimal solution:\n")
	v, sack := bestSolution(items, knapsackCapacity)
	for _, i := range sack {
		fmt.Println("   ", i)
	}
	fmt.Printf("--- Total value: $%.2f\n", v)

}
