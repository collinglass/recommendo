package main

import (
	"fmt"
	"github.com/collinglass/recommendo/algo"
	"github.com/collinglass/recommendo/data"
	"github.com/collinglass/recommendo/sort"
)

func main() {
	// Populate data list
	users := data.Populate()

	/* func(size of list, data.UserId, map of data.User, algo.SimFunc) */
	fmt.Println(sort.TopMatches(3, 0, users, algo.Pearson))
}
