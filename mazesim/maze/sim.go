package maze

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"

	"fmt"

	"github.com/param108/aimaze/mazesim/spec/grpc/maze"
)


func PlaceHero(s * maze.Simulation) {
	X := rand.Int31n(s.Maze.Size.Width)
	Y := rand.Int31n(s.Maze.Size.Height)

	sym, err := Get(s.Maze, X, Y)
	for err != nil || sym == WALL || (X == s.Maze.Exit.X && Y == s.Maze.Exit.Y) {
		X = rand.Int31n(s.Maze.Size.Width)
		Y = rand.Int31n(s.Maze.Size.Height)

		sym, err = Get(s.Maze, X, Y)
	}

	s.Hero.X = X
	s.Hero.Y = Y

	s.Prev.X = X
	s.Prev.Y = Y

}

// NewSim - Returns a new Sim with random maze and Hero position
func NewSim() (*maze.Simulation, error) {
	m := NewMaze()
	if err := Create(m); err != nil {
		return nil, err
	}

	sim := &maze.Simulation{
		Maze: m,
	}

	sim.Hero = &maze.Point{}

	sim.Prev = &maze.Point{}

	PlaceHero(sim)

	return sim, nil
}

const HERO = "H"

// LoadSim - Loads the Sim from json input
func LoadSim(filename string) (*maze.Simulation, error) {
	bts, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	s := &maze.Simulation{}
	if err := json.Unmarshal(bts, s); err != nil {
		return nil, err
	}

	return s, nil
}

// Save - save the sim board to a file as json
func SaveSim(s *maze.Simulation, filename string) error {
	bts, err := json.Marshal(s)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, bts, 0655)
}

// Print - prints the sim board
func PrintSim(s *maze.Simulation) {
	for y := int32(0); y < s.Maze.Size.Height; y++ {
		for x := int32(0); x < s.Maze.Size.Width; x++ {
			if x == s.Maze.Exit.X && y == s.Maze.Exit.Y {
				fmt.Print("E")
				continue
			}

			if x == s.Hero.X && y == s.Hero.Y {
				fmt.Print("H")
				continue
			}

			s, _ := Get(s.Maze, x, y)
			if len(s) == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print(s)
			}
		}
		fmt.Println("")
	}
}

const (
	UP    = "up"
	DOWN  = "down"
	RIGHT = "right"
	LEFT  = "left"
)

// DryMove - Try and Move the Hero and return the new position
// and if the Move is valid.
// Return values: X, Y, valid
func DryMove(s *maze.Simulation, direction string) (int32, int32, bool) {
	newX := s.Hero.X
	newY := s.Hero.Y

	valid := func(X, Y int32) (int32, int32, bool) {
		s, err := Get(s.Maze, X, Y)
		if err != nil {
			return X, Y, false
		}
		if s == WALL {
			return X, Y, false
		}
		return X, Y, true
	}

	switch direction {
	case UP:
		newY = newY - 1
		return valid(newX, newY)
	case DOWN:
		newY = newY + 1
		return valid(newX, newY)

	case RIGHT:
		newX = newX + 1
		return valid(newX, newY)

	case LEFT:
		newX = newX - 1
		return valid(newX, newY)
	}

	return newX, newY, false
}
