package data

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

func Populate() (map[int]User, PrefList) {
	users := make(map[int]User)
	prefs := NewPrefList()

	users[0] = User{
		UserId: 0,
	}
	prefs[users[0].UserId] = make(map[int]float64)
	prefs[users[0].UserId][0] = 2.5
	prefs[users[0].UserId][1] = 3.5
	prefs[users[0].UserId][2] = 3.0
	prefs[users[0].UserId][3] = 3.5
	prefs[users[0].UserId][4] = 2.5
	prefs[users[0].UserId][5] = 3.0

	users[1] = User{
		UserId: 1,
	}
	prefs[users[1].UserId] = make(map[int]float64)
	prefs[users[1].UserId][0] = 3.0
	prefs[users[1].UserId][1] = 3.5
	prefs[users[1].UserId][2] = 1.5
	prefs[users[1].UserId][3] = 5.0
	prefs[users[1].UserId][4] = 3.5
	prefs[users[1].UserId][5] = 3.0

	users[2] = User{
		UserId: 2,
	}
	prefs[users[2].UserId] = make(map[int]float64)
	prefs[users[2].UserId][0] = 2.5
	prefs[users[2].UserId][1] = 3.0
	prefs[users[2].UserId][3] = 3.5
	prefs[users[2].UserId][5] = 4.0

	users[3] = User{
		UserId: 3,
	}
	prefs[users[3].UserId] = make(map[int]float64)
	prefs[users[3].UserId][1] = 3.5
	prefs[users[3].UserId][2] = 3.0
	prefs[users[3].UserId][3] = 4.0
	prefs[users[3].UserId][4] = 2.5
	prefs[users[3].UserId][5] = 4.5

	users[4] = User{
		UserId: 4,
	}
	prefs[users[4].UserId] = make(map[int]float64)
	prefs[users[4].UserId][0] = 3.0
	prefs[users[4].UserId][1] = 4.0
	prefs[users[4].UserId][2] = 2.0
	prefs[users[4].UserId][3] = 3.0
	prefs[users[4].UserId][4] = 2.0
	prefs[users[4].UserId][5] = 3.0

	users[5] = User{
		UserId: 5,
	}
	prefs[users[5].UserId] = make(map[int]float64)
	prefs[users[5].UserId][0] = 3.0
	prefs[users[5].UserId][1] = 4.0
	prefs[users[5].UserId][3] = 5.0
	prefs[users[5].UserId][4] = 3.5
	prefs[users[5].UserId][5] = 3.0

	users[6] = User{
		UserId: 6,
	}
	prefs[users[6].UserId] = make(map[int]float64)
	prefs[users[6].UserId][1] = 4.5
	prefs[users[6].UserId][3] = 4.0
	prefs[users[6].UserId][4] = 1.0

	return users, prefs
}
