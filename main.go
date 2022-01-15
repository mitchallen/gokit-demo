package main

import (
	"fmt"

	"github.com/mitchallen/gokit"
	"github.com/mitchallen/gokit/calc"

	"github.com/mitchallen/coin"
)

func demoFlip() {

	// Test Weighted

	a := make([]bool, 100)

	for i, _ := range a {
		a[i] = coin.Flip()
	}

	heads := 0
	tails := 0

	for _, v := range a {
		if v {
			heads++
		} else {
			tails++
		}
	}

	fmt.Printf("[Flip()] Heads: %d Tails: %d \n", heads, tails)
}

func demoWeighted() {

	// Test Weighted

	a := make([]bool, 100)

	var testWeight float32 = 0.75

	for i, _ := range a {
		a[i] = coin.Weighted(testWeight)
	}

	heads := 0
	tails := 0

	for _, v := range a {
		if v {
			heads++
		} else {
			tails++
		}
	}

	fmt.Printf("[Weight(%.2f)] Heads: %d Tails: %d \n", testWeight, heads, tails)
}

func main() {
	fmt.Println(gokit.Info())

	fmt.Println(calc.Add(5, 6))

	fmt.Println(coin.Flip())

	fmt.Println(coin.Weighted(0.75))

	demoFlip()
	demoWeighted()

}
