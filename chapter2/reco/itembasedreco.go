package reco

import (
	//"errors"
	"fmt"
	"github.com/collinglass/recommendo/chapter2/algo"
	"github.com/collinglass/recommendo/chapter2/data"
	"github.com/collinglass/recommendo/chapter2/sort"
)

func ItemRecommend(prefpointer *data.PrefList, simpointer *data.SimList, user int) (map[int]data.Recommendation, error) {
	prefs := *prefpointer
	similarities := *simpointer
	userRatings := prefs[user]
	scores := make(map[int]float64)
	totalSim := make(map[int]float64)
	rankings := make(map[int]data.Recommendation)
	for item, rating := range userRatings {
		for item2, similarity := range similarities[item] {
			if userRatings[item2] == 0 {
				scores[item2] = similarity.Score * rating
				totalSim[item2] += similarity.Score
			}
		}
	}
	for item, score := range scores {
		rankings[item] = data.Recommendation{item, (score / totalSim[item])}
	}
	return rankings, nil
}

func GetSimilar(prefpointer *data.PrefList, simFunc algo.SimFunc) data.SimList {
	TransformPrefs(prefpointer)

	prefs := *prefpointer
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

func TransformPrefs(prefpointer *data.PrefList) {
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
	*prefpointer = result
}
