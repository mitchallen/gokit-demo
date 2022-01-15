package main

import (
	"fmt"

	"github.com/mitchallen/gokit"
	"github.com/mitchallen/gokit/calc"

	"github.com/mitchallen/coin"
)

func main() {
	fmt.Println(gokit.Info())

	fmt.Println(calc.Add(5, 6))

	fmt.Println(coin.Flip())
}
