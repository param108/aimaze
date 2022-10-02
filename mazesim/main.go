package main


import (
	"github.com/param108/aimaze/maze"
	"fmt"
)


func main() {
	m := maze.NewMaze()

	if err := m.Create(); err != nil {
		panic(err)
	}

	m.Print()

	fmt.Println("Saved")

	err := m.Save("saveit.json")
	if err != nil {
		panic(err)
	}

	m,err = maze.NewFromFile("saveit.json")
	if err != nil {
		panic(err)
	}

	m.Print()


}
