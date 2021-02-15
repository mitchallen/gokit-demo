package main

import (
	"fmt"

	"github.com/mitchallen/gokit"
	"github.com/mitchallen/gokit/calc"
)

func main() {
	fmt.Println(gokit.Info())

	fmt.Println(calc.Add(5, 6))
}
