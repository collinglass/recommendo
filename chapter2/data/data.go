package data

import (
	"math"
	"math/rand"
)

type User struct {
	UserId int
}

type Similar struct {
	UserId int
	Score  float64
}

type Recommendation struct {
	Book  int
	Score float64
}

type PrefList map[int]map[int]float64

type SimList map[int]map[int]Similar

type RecoList map[int]map[int]Recommendation

func NewPrefList() PrefList {
	return make(map[int]map[int]float64)
}

func NewRecoList() RecoList {
	return make(map[int]map[int]Recommendation)
}

func NewSimList() SimList {
	return make(map[int]map[int]Similar)
}

func round(x float64, prec int) float64 {
	var rounder float64
	pow := math.Pow(10, float64(prec))
	intermed := x * pow
	_, frac := math.Modf(intermed)
	if frac >= 0.5 {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}

	return rounder / pow
}

func Populate() (map[int]User, PrefList) {
	users := make(map[int]User)
	prefs := NewPrefList()
	r := rand.New(rand.NewSource(5))

	numberofusers := 1000
	numberofitems := 100

	i := 0
	j := 0
	for i < numberofusers {
		users[i] = User{
			UserId: i,
		}
		prefs[i] = make(map[int]float64)
		j = 0
		for j < numberofitems {
			if (j*i)%3 == 0 {
				prefs[i][j] = round(r.Float64()*5, 1)
			} else {
				prefs[i][j] = 0
			}
			j += 1
		}

		i += 1
	}

	return users, prefs
}
