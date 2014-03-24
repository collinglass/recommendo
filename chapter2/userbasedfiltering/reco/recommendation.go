package reco

import (
	"github.com/collinglass/recommendo/chapter2/userbasedfiltering/algo"
	"github.com/collinglass/recommendo/chapter2/userbasedfiltering/data"
)

func Recommend(prefpointer *data.PrefList, person int, simFunc algo.SimFunc) data.RecoList {
	prefs := *prefpointer
	reco := make(map[int]float64)
	simSum := make(map[int]float64)

	for other, _ := range prefs {
		if other != person {
			sim := simFunc(&prefs, person, other)
			if sim > 0 {
				for item, _ := range prefs[other] {
					if prefs[person][item] == 0 {
						reco[item] += prefs[other][item] * sim
						simSum[item] += sim
					}
				}
			}
		}
	}

	// Create Normalized list
	rankings := data.NewRecoList()
	rankings[person] = make(map[int]data.Recommendation)
	for item, recomendation := range reco {
		rankings[person][item] = data.Recommendation{item, recomendation / simSum[item]}
	}
	return rankings
}
