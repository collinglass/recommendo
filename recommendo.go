package main

import (
	"fmt"
	"github.com/collinglass/recommendo/data"
	"github.com/collinglass/recommendo/eucl"
)

func main() {
	users := data.Populate()

	eucl.EuclideanUsers(users)
	fmt.Println(users[0].Similars)
}
