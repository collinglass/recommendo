package eucl

import (
	"github.com/collinglass/recommendo/data"
	"math"
)

func Euclidean2(user1 *data.User, user2 *data.User) {
	diff := make(map[int]float64)

	for key, _ := range user1.Ratings {
		if user1.Ratings[key] != 0 && user2.Ratings[key] != 0 {
			diff[key] = math.Pow((user1.Ratings[key] + user2.Ratings[key]), 2)
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

func EuclideanUsers(users map[int]data.User) map[int]data.User {
	for user1, _ := range users {
		for user2, _ := range users {
			if user1 != user2 {
				u1 := users[user1]
				u2 := users[user2]
				Euclidean2(&u1, &u2)
			}
		}
	}
	return users
}
