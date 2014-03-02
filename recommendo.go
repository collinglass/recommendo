package main

import (
	"fmt"
	"github.com/collinglass/recommendo/algo"
	"github.com/collinglass/recommendo/data"
	"github.com/collinglass/recommendo/reco"
)

func main() {
	// Populate data list
	users := data.Populate()

	userId := 6

	// recommendation list using Pearson
	reco.Recommend(userId, users, algo.Pearson)

	// Print user's recommended list
	fmt.Println("Pearson: ", users[userId].Recommended)

	// recommendation list using Euclidean
	reco.Recommend(userId, users, algo.Euclidean)

	// Print user's recommended list
	fmt.Println("Euclidean: ", users[userId].Recommended)
}
