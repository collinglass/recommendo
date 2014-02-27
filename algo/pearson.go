package algo

import (
	"github.com/collinglass/recommendo/data"
	"math"
)

func Pearson2(user1 *data.User, user2 *data.User) {
	var sum1 float64
	var sum2 float64
	var sumsq1 float64
	var sumsq2 float64
	var psum float64
	var n int

	for key, value := range user1.Ratings {
		if value != 0 && user2.Ratings[key] != 0 {
			// Sum all ratings
			sum1 += value
			sum2 += user2.Ratings[key]

			// Sum all ratings squared
			sumsq1 += math.Pow(value, 2)
			sumsq2 += math.Pow(user2.Ratings[key], 2)

			// Sum all products
			psum += value * user2.Ratings[key]

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

	id1 := user1.UserId
	id2 := user2.UserId

	user1.Similars[id2] = data.Similar{id2, result}
	user2.Similars[id1] = data.Similar{id1, result}
}

func PearsonUsers(users map[int]data.User) map[int]data.User {
	for user1, _ := range users {
		for user2, _ := range users {
			if user1 != user2 {
				u1 := users[user1]
				u2 := users[user2]
				Pearson2(&u1, &u2)
			}
		}
	}
	return users
}
