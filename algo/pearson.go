package algo

import (
	"github.com/collinglass/recommendo/data"
	"math"
)

func Pearson(prefpointer *data.PrefList, user1 int, user2 int) float64 {
	var sum1 float64
	var sum2 float64
	var sumsq1 float64
	var sumsq2 float64
	var psum float64
	var n int
	prefs := *prefpointer

	for key, value := range prefs[user1] {
		if value != 0 && prefs[user2][key] != 0 {
			// Sum all prefs
			sum1 += value
			sum2 += prefs[user2][key]

			// Sum all prefs squared
			sumsq1 += math.Pow(value, 2)
			sumsq2 += math.Pow(prefs[user2][key], 2)

			// Sum all products
			psum += value * prefs[user2][key]

			// Count common
			n++
		}
	}
	var nfloat float64

	// convert int to float for pearson calculation
	nfloat = float64(n)

	// Calculate Pearson Score
	num := psum - (sum1 * sum2 / nfloat)
	den := math.Sqrt((sumsq1 - math.Pow(sum1, 2)/nfloat) * (sumsq2 - math.Pow(sum2, 2)/nfloat))

	var result float64
	if den == 0.0 {
		result = 0.0
	} else {
		result = num / den
	}

	return result
}
