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

	// recommendation list
	reco.Recommend(userId, users, algo.Pearson)

	// Print user's recommended list
	fmt.Println(users[userId].Recommended)
}
