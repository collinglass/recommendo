package main

import (
	"fmt"
	"github.com/collinglass/recommendo/data"
)

func main() {
	users := data.Populate()
	fmt.Println(users)
}
