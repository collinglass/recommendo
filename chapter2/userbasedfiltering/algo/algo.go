package algo

import (
	"github.com/collinglass/recommendo/chapter2/userbasedfiltering/data"
)

type SimFunc func(prefs *data.PrefList, user1 int, user2 int) float64

func GetSimilar(prefs *data.PrefList, users map[int]data.User, simFunc SimFunc) map[int]data.User {
	for user1, _ := range users {
		for user2, _ := range users {
			if user1 != user2 {
				u1 := users[user1]
				u2 := users[user2]
				simFunc(prefs, u1.UserId, u2.UserId)
			}
		}
	}
	return users
}
