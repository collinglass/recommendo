package main

import (
	"fmt"
	"github.com/collinglass/recommendo/itembasedfiltering/algo"
	"github.com/collinglass/recommendo/itembasedfiltering/data"
	"github.com/collinglass/recommendo/itembasedfiltering/reco"
)

def transformPrefs(prefs):
      result={}
      for person in prefs:
        for item in prefs[person]:
          result.setdefault(item,{})
          # Flip item and person
          result[item][person]=prefs[person][item]
      return result

func transformPrefs(prefs) {
	result := []
	
}

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
