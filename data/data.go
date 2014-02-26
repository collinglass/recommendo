package data

type User struct {
	BookZero  float32
	BookOne   float32
	BookTwo   float32
	BookThree float32
	BookFour  float32
	BookFive  float32
}

func Populate() map[int]User {
	users := make(map[int]User)

	user1 := &User{
		BookZero:  2.5,
		BookOne:   3.5,
		BookTwo:   3.0,
		BookThree: 3.5,
		BookFour:  2.5,
		BookFive:  3.0,
	}
	user2 := &User{
		BookZero:  3.0,
		BookOne:   3.5,
		BookTwo:   1.5,
		BookThree: 5.0,
		BookFive:  3.0,
		BookFour:  3.5,
	}
	user3 := &User{
		BookZero:  2.5,
		BookOne:   3.0,
		BookThree: 3.5,
		BookFive:  4.0,
	}
	user4 := &User{
		BookOne:   3.5,
		BookTwo:   3.0,
		BookFive:  4.5,
		BookThree: 4.0,
		BookFour:  2.5,
	}
	user5 := &User{
		BookZero:  3.0,
		BookOne:   4.0,
		BookTwo:   2.0,
		BookThree: 3.0,
		BookFive:  3.0,
		BookFour:  2.0,
	}
	user6 := &User{
		BookZero:  3.0,
		BookOne:   4.0,
		BookFive:  3.0,
		BookThree: 5.0,
		BookFour:  3.5,
	}
	user7 := &User{
		BookOne:   4.5,
		BookFour:  1.0,
		BookThree: 4.0,
	}
	users[0] = *user1
	users[1] = *user2
	users[2] = *user3
	users[3] = *user4
	users[4] = *user5
	users[5] = *user6
	users[6] = *user7

	return users
}
