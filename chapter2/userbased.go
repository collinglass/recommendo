package main

import (
	"fmt"
	"github.com/collinglass/recommendo/chapter2/algo"
	"github.com/collinglass/recommendo/chapter2/data"
	"github.com/collinglass/recommendo/chapter2/reco"
)

func main() {
	// Populate data list
	_, prefs := data.Populate()

	// USER-BASED Filtering
	userId := 2

	// recommendation list using Pearson
	recolist, err := reco.Recommend(&prefs, userId, algo.Pearson)
	if err != nil {
		fmt.Println(err)
	}

	// Print user's recommended list
	fmt.Println("User-based Pearson: ", recolist[userId])

	// recommendation list using Euclidean
	recolist, err = reco.Recommend(&prefs, userId, algo.Euclidean)
	if err != nil {
		fmt.Println(err)
	}

	// Print user's recommended list
	fmt.Println("User-based Euclidean: ", recolist[userId])

}
