package reco

import (
	"fmt"
	"github.com/collinglass/recommendo/algo"
	"github.com/collinglass/recommendo/data"
	"github.com/collinglass/recommendo/sort"
)

func ItemRunner() (map[int]data.Recommendation, map[int]data.Recommendation) {
	// Populate data list
	_, prefs := data.Populate()

	// ITEM-BASED Filtering
	userId := 2

	simlist := GetSimilar(&prefs, algo.Pearson)

	// recommendation list using Pearson
	recolist, err := ItemBasedRecommend(&prefs, &simlist, userId)
	if err != nil {
		fmt.Println(err)
	}

	// Print user's recommended list
	//fmt.Println("Item-based Pearson: ", recolist)

	simlist = GetSimilar(&prefs, algo.Euclidean)

	// recommendation list using Euclidean
	recolist2, err := ItemBasedRecommend(&prefs, &simlist, userId)
	if err != nil {
		fmt.Println(err)
	}

	// Print user's recommended list
	//fmt.Println("Item-based Euclidean: ", recolist2)

	return recolist, recolist2
}

func ItemBasedRecommend(prefpointer *data.PrefList, simpointer *data.SimList, user int) (map[int]data.Recommendation, error) {
	prefs := *prefpointer
	similarities := *simpointer
	userRatings := prefs[user]
	scores := make(map[int]float64)
	totalSim := make(map[int]float64)
	rankings := make(map[int]data.Recommendation)
	for item, rating := range userRatings {
		for item2, similarity := range similarities[item] {
			if userRatings[item2] == 0 {
				scores[item2] += similarity.Score * rating
				totalSim[item2] += similarity.Score
			}
		}
	}
	for item, score := range scores {
		rankings[item] = data.Recommendation{item, (score / totalSim[item])}
	}
	return sort.SortMapByValue(rankings), nil
}

func GetSimilar(prefpointer *data.PrefList, simFunc algo.SimFunc) data.SimList {
	prefs := transformPrefs(prefpointer)

	simlist := data.NewSimList()
	c := 0
	for item, _ := range prefs {
		c += 1
		if c%100 == 0 {
			fmt.Println("%d / %d", c, len(prefs))
		}
		if len(simlist[item]) == 0 {
			simlist[item] = make(map[int]data.Similar)
		}
		simlist[item] = sort.TopMatches(&prefs, item, simFunc)
	}
	return simlist
}

func transformPrefs(prefpointer *data.PrefList) data.PrefList {
	prefs := *prefpointer
	result := data.NewPrefList()
	for person, items := range prefs {
		for item, pref := range items {
			if len(result[item]) == 0 {
				result[item] = make(map[int]float64)
			}
			result[item][person] = pref
		}
	}
	return result
}
