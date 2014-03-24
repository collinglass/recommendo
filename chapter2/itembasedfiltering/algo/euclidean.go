package algo

import (
	"github.com/collinglass/recommendo/itembasedfiltering/data"
	"math"
)

func Euclidean(user1 *data.User, user2 *data.User) {
	diff := make(map[int]float64)

	for key, _ := range user1.Ratings {
		if user1.Ratings[key] != 0 && user2.Ratings[key] != 0 {
			diff[key] = math.Pow((user1.Ratings[key] - user2.Ratings[key]), 2)
		}
	}

	var result float64
	for _, value := range diff {
		result += value
	}
	result = 1 / (1 + math.Sqrt(result))

	id1 := user1.UserId
	id2 := user2.UserId

	user1.Similars[id2] = data.Similar{id2, result}
	user2.Similars[id1] = data.Similar{id1, result}
}
