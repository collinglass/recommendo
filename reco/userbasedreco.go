package reco

import (
	"fmt"
	"github.com/collinglass/recommendo/algo"
	"github.com/collinglass/recommendo/data"
)

func UserRunner() (map[int]data.Recommendation, map[int]data.Recommendation) {
	// Populate data list
	_, prefs := data.Populate()

	// USER-BASED Filtering
	userId := 2

	// recommendation list using Pearson
	recolist, err := UserBasedRecommend(&prefs, userId, algo.Pearson)
	if err != nil {
		fmt.Println(err)
	}

	// Print user's recommended list
	//fmt.Println("User-based Pearson: ", recolist[userId])

	// recommendation list using Euclidean
	recolist2, err := UserBasedRecommend(&prefs, userId, algo.Euclidean)
	if err != nil {
		fmt.Println(err)
	}

	// Print user's recommended list
	//fmt.Println("User-based Euclidean: ", recolist2[userId])

	return recolist[userId], recolist2[userId]
}

func UserBasedRecommend(prefpointer *data.PrefList, person int, simFunc algo.SimFunc) (data.RecoList, error) {
	prefs := *prefpointer

	/*if _, ok := prefs[person]; ok {
		if !ok {
			return nil, errors.New("Id is out of bounds!")
		}
	}*/
	reco := make(map[int]float64)
	simSum := make(map[int]float64)

	for other, _ := range prefs {
		if other != person {
			sim := simFunc(&prefs, person, other)
			if sim > 0 {
				for item, value := range prefs[other] {
					if prefs[person][item] == 0 {
						reco[item] += value * sim
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
	return rankings, nil
}
