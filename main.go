package main

import (
	"fmt"
	"github.com/wesovilabs.com/go-mock-example/matches"
)

func main() {
	fmt.Print(matches.GetResultForMatch("Madrid", "Barcelona"))
	fmt.Print(matches.GetResultForMatch("Madrid", "Barcelona"))
	fmt.Print(matches.GetResultForMatch("Madrid", "Barcelona"))
	fmt.Print(matches.GetResultForMatch("Madrid", "Barcelona"))
}
