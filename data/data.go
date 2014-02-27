package data

type User struct {
	UserId   int
	Ratings  map[int]float64
	Similars map[int]Similar
}

type Similar struct {
	UserId   int
	Distance float64
}

func Populate() map[int]User {
	users := make(map[int]User)

	user0 := &User{
		UserId: 0,
		Ratings: map[int]float64{
			0: 2.5,
			1: 3.5,
			2: 3.0,
			3: 3.5,
			4: 2.5,
			5: 3.0,
		},
		Similars: make(map[int]Similar),
	}

	user1 := &User{
		UserId: 1,
		Ratings: map[int]float64{
			0: 3.0,
			1: 3.5,
			2: 1.5,
			3: 5.0,
			5: 3.0,
			4: 3.5,
		},
		Similars: make(map[int]Similar),
	}
	user2 := &User{
		UserId: 2,
		Ratings: map[int]float64{
			0: 2.5,
			1: 3.0,
			3: 3.5,
			5: 4.0,
		},
		Similars: make(map[int]Similar),
	}
	user3 := &User{
		UserId: 3,
		Ratings: map[int]float64{
			1: 3.5,
			2: 3.0,
			5: 4.5,
			3: 4.0,
			4: 2.5,
		},
		Similars: make(map[int]Similar),
	}
	user4 := &User{
		UserId: 4,
		Ratings: map[int]float64{
			0: 3.0,
			1: 4.0,
			2: 2.0,
			3: 3.0,
			5: 3.0,
			4: 2.0,
		},
		Similars: make(map[int]Similar),
	}
	user5 := &User{
		UserId: 5,
		Ratings: map[int]float64{
			0: 3.0,
			1: 4.0,
			5: 3.0,
			3: 5.0,
			4: 3.5,
		},
		Similars: make(map[int]Similar),
	}
	user6 := &User{
		UserId: 6,
		Ratings: map[int]float64{
			1: 4.5,
			4: 1.0,
			3: 4.0,
		},
		Similars: make(map[int]Similar),
	}
	users[0] = *user0
	users[1] = *user1
	users[2] = *user2
	users[3] = *user3
	users[4] = *user4
	users[5] = *user5
	users[6] = *user6

	return users
}
