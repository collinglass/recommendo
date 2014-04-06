package algo

import (
	"github.com/collinglass/recommendo/data"
	"math"
)

func Euclidean(prefpointer *data.PrefList, user1 int, user2 int) float64 {
	diff := make(map[int]float64)
	prefs := *prefpointer

	for key, value := range prefs[user1] {
		if value != 0 && prefs[user2][key] != 0 {
			diff[key] = math.Pow((value - prefs[user2][key]), 2)
		}
	}

	var result float64
	for _, value := range diff {
		result += value
	}
	return 1 / (1 + math.Sqrt(result))
}
