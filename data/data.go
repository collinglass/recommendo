package data

type User struct {
	UserId    int
	BookZero  float64
	BookOne   float64
	BookTwo   float64
	BookThree float64
	BookFour  float64
	BookFive  float64
	Similars  map[int]Similar
}

type Similar struct {
	UserId   int
	Distance float64
}

func Populate() map[int]User {
	users := make(map[int]User)

	user0 := &User{
		UserId:    0,
		BookZero:  2.5,
		BookOne:   3.5,
		BookTwo:   3.0,
		BookThree: 3.5,
		BookFour:  2.5,
		BookFive:  3.0,
		Similars:  make(map[int]Similar),
	}
	user1 := &User{
		UserId:    1,
		BookZero:  3.0,
		BookOne:   3.5,
		BookTwo:   1.5,
		BookThree: 5.0,
		BookFive:  3.0,
		BookFour:  3.5,
		Similars:  make(map[int]Similar),
	}
	user2 := &User{
		UserId:    2,
		BookZero:  2.5,
		BookOne:   3.0,
		BookThree: 3.5,
		BookFive:  4.0,
		Similars:  make(map[int]Similar),
	}
	user3 := &User{
		UserId:    3,
		BookOne:   3.5,
		BookTwo:   3.0,
		BookFive:  4.5,
		BookThree: 4.0,
		BookFour:  2.5,
		Similars:  make(map[int]Similar),
	}
	user4 := &User{
		UserId:    4,
		BookZero:  3.0,
		BookOne:   4.0,
		BookTwo:   2.0,
		BookThree: 3.0,
		BookFive:  3.0,
		BookFour:  2.0,
		Similars:  make(map[int]Similar),
	}
	user5 := &User{
		UserId:    5,
		BookZero:  3.0,
		BookOne:   4.0,
		BookFive:  3.0,
		BookThree: 5.0,
		BookFour:  3.5,
		Similars:  make(map[int]Similar),
	}
	user6 := &User{
		UserId:    6,
		BookOne:   4.5,
		BookFour:  1.0,
		BookThree: 4.0,
		Similars:  make(map[int]Similar),
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
