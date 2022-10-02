package sim

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"

	"fmt"

	"github.com/param108/aimaze/maze"
)

type Sim struct {
	M *maze.Maze
	H maze.Point // Hero
}

func (s *Sim) PlaceHero() {
	X := rand.Intn(s.M.S.Width)
	Y := rand.Intn(s.M.S.Height)

	sym, err := s.M.Get(X, Y)
	for err != nil || sym == maze.WALL || (X == s.M.E.X && Y == s.M.E.Y) {
		X = rand.Intn(s.M.S.Width)
		Y = rand.Intn(s.M.S.Height)

		sym, err = s.M.Get(X, Y)
	}

	s.H.X = X
	s.H.Y = Y
}

// NewSim - Returns a new Sim with random maze and Hero position
func NewSim() (*Sim, error) {
	m := maze.NewMaze()
	if err := m.Create(); err != nil {
		return nil, err
	}

	sim := &Sim{
		M: m,
	}

	sim.PlaceHero()

	return sim, nil
}

const HERO = "H"

// LoadSim - Loads the Sim from json input
func LoadSim(filename string) (*Sim, error) {
	bts, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	s := &Sim{}
	if err := json.Unmarshal(bts, s); err != nil {
		return nil, err
	}

	return s, nil
}

// Save - save the sim board to a file as json
func (s *Sim) Save(filename string) error {
	bts, err := json.Marshal(s)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, bts, 0655)
}

// Print - prints the sim board
func (s *Sim) Print() {
	for y := 0; y < s.M.S.Height; y++ {
		for x := 0; x < s.M.S.Width; x++ {
			if x == s.M.E.X && y == s.M.E.Y {
				fmt.Print("E")
				continue
			}

			if x == s.H.X && y == s.H.Y {
				fmt.Print("H")
				continue
			}

			s, _ := s.M.Get(x, y)
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
func (s *Sim) DryMove(direction string) (int, int, bool) {
	newX := s.H.X
	newY := s.H.Y

	valid := func(X, Y int) (int, int, bool) {
		s, err := s.M.Get(X, Y)
		if err != nil {
			return X, Y, false
		}
		if s == maze.WALL {
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
		newY = newX + 1
		return valid(newX, newY)

	case LEFT:
		newY = newX - 1
		return valid(newX, newY)
	}

	return newX, newY, false
}
