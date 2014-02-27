package main

import (
	"fmt"
	"github.com/collinglass/recommendo/algo"
	"github.com/collinglass/recommendo/data"
)

func main() {
	users := data.Populate()

	algo.PearsonUsers(users)
	//algo.EuclideanUsers(users)
	fmt.Println(users[0].Similars)
}
