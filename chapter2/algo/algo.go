package algo

import (
	"github.com/collinglass/recommendo/chapter2/data"
)

type SimFunc func(prefs *data.PrefList, user1 int, user2 int) float64
