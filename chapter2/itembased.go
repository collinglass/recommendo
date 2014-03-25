package main

import (
	"fmt"
	"github.com/collinglass/recommendo/chapter2/algo"
	"github.com/collinglass/recommendo/chapter2/data"
	"github.com/collinglass/recommendo/chapter2/reco"
	"github.com/collinglass/recommendo/chapter2/sort"
)

def calculateSimilarItems(prefs,n=10):
      # Create a dictionary of items showing which other items they
      # are most similar to.
      result={}
      # Invert the preference matrix to be item-centric
      itemPrefs=transformPrefs(prefs)
      c=0
      for item in itemPrefs:
        # Status updates for large datasets
        c+=1
        if c%100==0: print "%d / %d" % (c,len(itemPrefs))
        # Find the most similar items to this one
        scores=topMatches(itemPrefs,item,n=n,similarity=sim_distance)
        result[item]=scores
      return result

func GetSimilar(prefpointer *data.PrefList, n int, simFunc algo.SimFunc) data.SimList {
	prefs = *prefpointer
	simlist = data.NewSimList()
	c := 0
	for item, _ := range prefs {
		c += 1
		if c % 100 == 0 { fmt.Println("%d / %d", c, len(prefs)) }
		simlist[item] = sort.TopMatches(&prefs, item, n, simFunc)
	}
	return simlist
}

func main() {
	// Populate data list
	_, prefs := data.Populate()

	// ITEM-BASED Filtering
	itemId := 2

	reco.TransformPrefs(&prefs)

	// recommendation list using Pearson
	recolist, err := reco.Recommend(&prefs, itemId, algo.Pearson)
	if err != nil {
		fmt.Println(err)
	}

	// Print user's recommended list
	fmt.Println("Item-based Pearson: ", recolist[itemId])

	// recommendation list using Euclidean
	recolist, err = reco.Recommend(&prefs, itemId, algo.Euclidean)
	if err != nil {
		fmt.Println(err)
	}

	// Print user's recommended list
	fmt.Println("Item-based Euclidean: ", recolist[itemId])
}
