package algo

import (
	"github.com/collinglass/recommendo/data"
)

type SimFunc func(user1 *data.User, user2 *data.User)

func GetSimilar(users map[int]data.User, simFunc SimFunc) map[int]data.User {
	for user1, _ := range users {
		for user2, _ := range users {
			if user1 != user2 {
				u1 := users[user1]
				u2 := users[user2]
				simFunc(&u1, &u2)
			}
		}
	}
	return users
}
