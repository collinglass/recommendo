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

	// USER-BASED Filtering
	userId := 6

	// recommendation list using Pearson
	recolist := reco.Recommend(&prefs, userId, algo.Pearson)

	// Print user's recommended list
	fmt.Println("User-based Pearson: ", recolist[userId])

	// recommendation list using Euclidean
	recolist = reco.Recommend(&prefs, userId, algo.Euclidean)

	// Print user's recommended list
	fmt.Println("User-based Euclidean: ", recolist[userId])

	// ITEM-BASED Filtering
	itemId := 2

	reco.TransformPrefs(&prefs)

	// recommendation list using Pearson
	recolist := reco.Recommend(&prefs, itemId, algo.Pearson)

	// Print user's recommended list
	fmt.Println("Item-based Pearson: ", recolist[itemId])

	// recommendation list using Euclidean
	recolist = reco.Recommend(&prefs, itemId, algo.Euclidean)

	// Print user's recommended list
	fmt.Println("Item-based Euclidean: ", recolist[itemId])
}
