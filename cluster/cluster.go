package main

import (
	"fmt"
	"github.com/collinglass/recommendo/algo"
	"github.com/collinglass/recommendo/data"
	"github.com/collinglass/recommendo/reco"
	"math"
	"math/rand"
	"reflect"
)

func kcluster(prefpointer *data.PrefList, k int) map[int]map[int]int {
	prefs := *prefpointer
	r := rand.New(rand.NewSource(99))
	ranges := make(map[int]map[int]float64)
	clusters := make(map[int]map[int]float64)

	for i, pref := range prefs {
		if len(ranges[i]) == 0 {
			ranges[i] = make(map[int]float64)
		}
		for _, t := range pref {
			if t < ranges[i][0] {
				ranges[i][0] = t
			}
			if t > ranges[i][1] {
				ranges[i][1] = t
			}
		}
	}
	i := 0
	for i < k {
		clusters[i] = make(map[int]float64)
		for j, value := range ranges {
			clusters[i][j] = r.Float64()*(value[1]-value[0]) + value[0]
		}
		i++
	}

	lastmatches := make(map[int]map[int]int)
	bestmatches := make(map[int]map[int]int)
	i = 0
	t := 0
	for t < 100 {
		count := 0
		fmt.Println("Iteration ", t)
		for i < k {
			bestmatches[i] = make(map[int]int)
			i++
		}
		for j, pref := range prefs {
			bestmatch := 0
			i = 0
			for i < k {
				d := euclidean(clusters[i], pref)
				if d < euclidean(clusters[bestmatch], pref) {
					bestmatch = i
				}
				i++
			}
			bestmatches[bestmatch][count] = j
			count += 1
		}
		if reflect.DeepEqual(lastmatches, bestmatches) {
			break
		}
		lastmatches = bestmatches
		i = 0
		for i < k {
			avgs := make(map[int]float64)
			if len(bestmatches[i]) > 0 {
				for rowid, _ := range bestmatches[i] {
					m := 0
					for m < len(prefs[rowid]) {
						avgs[m] += prefs[rowid][m]
						m++
					}
				}
				for _, avg := range avgs {
					avg = avg / float64(len(bestmatches[i]))
				}
				clusters[i] = avgs
			}
			i++
		}
		t++
	}

	fmt.Println("Best Matches: ", bestmatches)
	return bestmatches
}

func euclidean(user1 map[int]float64, user2 map[int]float64) float64 {
	diff := make(map[int]float64)

	for key, value := range user1 {
		if value != 0 && user2[key] != 0 {
			diff[key] = math.Pow((value - user2[key]), 2)
		}
	}

	var result float64
	for _, value := range diff {
		result += value
	}
	return 1 / (1 + math.Sqrt(result))
}

func ClusterRunner() map[int]map[int]int {
	// Populate data list
	_, prefs := data.Populate()

	return kcluster(&prefs, 3)
}
