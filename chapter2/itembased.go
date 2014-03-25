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
	_, prefs2 := data.Populate()

	// ITEM-BASED Filtering
	//itemId := 2
	userId := 2

	simlist := reco.GetSimilar(&prefs, algo.Pearson)

	// recommendation list using Pearson
	recolist, err := reco.ItemBasedRecommend(&prefs, &simlist, userId)
	if err != nil {
		fmt.Println(err)
	}

	// Print user's recommended list
	fmt.Println("Item-based Pearson: ", recolist)

	simlist = reco.GetSimilar(&prefs2, algo.Euclidean)

	// recommendation list using Euclidean
	recolist, err = reco.ItemBasedRecommend(&prefs2, &simlist, userId)
	if err != nil {
		fmt.Println(err)
	}

	// Print user's recommended list
	fmt.Println("Item-based Euclidean: ", recolist)
}
