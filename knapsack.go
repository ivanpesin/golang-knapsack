package knapsack

import (
	"fmt"
	"math"
	"sort"
)

type Item struct {
	Name   string
	Value  float64
	Weight float64
}

func (i Item) String() string {
	return fmt.Sprintf(" %-10s $%10.2f %10.2f kg", i.Name, i.Value, i.Weight)
}

// Greedy implements greedy knapsack alg
func Greedy(items []Item, maxWeight float64, metric func(i, j int) bool) (r []Item) {

	sort.Slice(items, metric)
	//fmt.Println("D: ", items)
	w := 0.

	for _, i := range items {
		if w+i.Weight <= maxWeight {
			r = append(r, i)
			w += i.Weight
		}
	}

	return

}

// combinations returns all possible combinations of items in store.
// Possible combinations are sent to a channel to avoid large memory consumption
func combinations(items []Item, ch chan []Item) {
	defer close(ch)

	p := int(math.Pow(2., float64(len(items))))

	for i := 0; i < p; i++ {
		set := []Item{}
		for j := 0; j < len(items); j++ {
			if (i>>uint(j))&1 == 1 {
				set = append(set, items[j])
			}
		}
		ch <- set
	}
}

// getSackWeight returns weight of a given set of items
func getSackWeight(set []Item) (r float64) {
	for _, i := range set {
		r += i.Weight
	}
	return
}

// getSackValue returns value of a given set
func getSackValue(set []Item) (r float64) {
	for _, i := range set {
		r += i.Value
	}
	return
}

// BestSolution looks through all possible combinations of items
// and selects the one with highest value which is below or eq target weight
func BestSolution(items []Item, maxWeight float64) (float64, []Item) {
	bestVal := 0.
	bestSack := []Item{}

	ch := make(chan []Item)
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
