package main

import (
	"fmt"
	"github.com/collinglass/recommendo/chapter2/userbasedfiltering/algo"
	"github.com/collinglass/recommendo/chapter2/userbasedfiltering/data"
	"github.com/collinglass/recommendo/chapter2/userbasedfiltering/reco"
)

func main() {
	// Populate data list
	_, prefs := data.Populate()

	userId := 6

	// recommendation list using Pearson
	recolist := reco.Recommend(&prefs, userId, algo.Pearson)

	// Print user's recommended list
	fmt.Println("Pearson: ", recolist[userId])

	// recommendation list using Euclidean
	recolist = reco.Recommend(&prefs, userId, algo.Euclidean)

	// Print user's recommended list
	fmt.Println("Euclidean: ", recolist[userId])
}
