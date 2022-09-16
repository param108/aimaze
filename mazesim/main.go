package main


import (
	"github.com/param108/aimaze/maze"
)


func main() {
	m := maze.NewMaze()

	if err := m.Create(); err != nil {
		panic(err)
	}

	m.Print()
}
