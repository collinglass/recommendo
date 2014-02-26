package eucl

import (
	"github.com/collinglass/recommendo/data"
	"math"
)

func Euclidean2(user1 *data.User, user2 *data.User) {
	diff := make([]float64, 6, 6)

	if user1.BookZero != 0 && user2.BookZero != 0 {
		diff[0] = math.Pow((user1.BookZero + user2.BookZero), 2)
	}
	if user1.BookOne != 0 && user2.BookOne != 0 {
		diff[1] = math.Pow((user1.BookOne + user2.BookOne), 2)
	}
	if user1.BookTwo != 0 && user2.BookTwo != 0 {
		diff[2] = math.Pow((user1.BookTwo + user2.BookTwo), 2)
	}
	if user1.BookThree != 0 && user2.BookThree != 0 {
		diff[3] = math.Pow((user1.BookThree + user2.BookThree), 2)
	}
	if user1.BookFour != 0 && user2.BookFour != 0 {
		diff[4] = math.Pow((user1.BookFour + user2.BookFour), 2)
	}
	if user1.BookFive != 0 && user2.BookFive != 0 {
		diff[5] = math.Pow((user1.BookFive + user2.BookFive), 2)
	}

	var result float64
	for i := 0; i < len(diff); i++ {
		result += diff[i]
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
