package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"

	knapsack "github.com/ivanpesin/golang-knapsack"
)

var store = []knapsack.Item{}
var knapsackCapacity = -1.

// readStore reads items and their properties from a file
func readStore(fn string) {
	f, err := os.Open(fn)
	if err != nil {
		fmt.Printf("ERROR: Unable to open file: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Reading store from file: %s\n", fn)
	defer f.Close()
	store = []knapsack.Item{}
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
		store = append(store, knapsack.Item{Name: fields[0], Value: val, Weight: weight})
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

	items := make([]knapsack.Item, len(store))
	copy(items, store)

	funcs := make([]func(i, j int) bool, 3)
	funcs[0] = func(i, j int) bool { return items[i].Value > items[j].Value }
	funcs[1] = func(i, j int) bool { return items[i].Weight < items[j].Weight }
	funcs[2] = func(i, j int) bool {
		return items[i].Value/items[i].Weight > items[j].Value/items[j].Weight
	}

	funcNames := make([]string, 3)
	funcNames[0] = "value"
	funcNames[1] = "weight"
	funcNames[2] = "density"

	for n, fname := range funcNames {
		v := 0.
		fmt.Printf("Being greedy based on %s: \n", fname)
		for _, i := range knapsack.Greedy(items, knapsackCapacity, funcs[n]) {
			fmt.Println("   ", i)
			v += i.Value
		}
		fmt.Printf("--- Total value: $%.2f\n", v)
	}

	sort.Slice(items, func(i, j int) bool { return items[i].Name < items[j].Name })
	fmt.Printf("\nOptimal solution:\n")
	v, sack := knapsack.BestSolution(items, knapsackCapacity)
	for _, i := range sack {
		fmt.Println("   ", i)
	}
	fmt.Printf("--- Total value: $%.2f\n", v)

}
