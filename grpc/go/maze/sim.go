package maze

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"

	"fmt"

)


func (s *Simulation) PlaceHero() {
	X := rand.Int31n(s.Maze.Size.Width)
	Y := rand.Int31n(s.Maze.Size.Height)

	sym, err := s.Maze.Get(X, Y)
	for err != nil || sym == WALL || (X == s.Maze.Exit.X && Y == s.Maze.Exit.Y) {
		X = rand.Int31n(s.Maze.Size.Width)
		Y = rand.Int31n(s.Maze.Size.Height)

		sym, err = s.Maze.Get(X, Y)
	}

	s.Hero.X = X
	s.Hero.Y = Y
}

// NewSim - Returns a new Sim with random maze and Hero position
func NewSim() (*Simulation, error) {
	m := NewMaze()
	if err := m.Create(); err != nil {
		return nil, err
	}

	sim := &Simulation{
		Maze: m,
	}

	sim.Hero = &Point{}

	sim.PlaceHero()

	return sim, nil
}

const HERO = "H"

// LoadSim - Loads the Sim from json input
func LoadSim(filename string) (*Simulation, error) {
	bts, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	s := &Simulation{}
	if err := json.Unmarshal(bts, s); err != nil {
		return nil, err
	}

	return s, nil
}

// Save - save the sim board to a file as json
func (s *Simulation) Save(filename string) error {
	bts, err := json.Marshal(s)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, bts, 0655)
}

// Print - prints the sim board
func (s *Simulation) Print() {
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

			s, _ := s.Maze.Get(x, y)
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
func (s *Simulation) DryMove(direction string) (int32, int32, bool) {
	newX := s.Hero.X
	newY := s.Hero.Y

	valid := func(X, Y int32) (int32, int32, bool) {
		s, err := s.Maze.Get(X, Y)
		if err != nil {
			return X, Y, false
		}
		if s == WALL {
			return X, Y, false
		}
		return X, Y, true
	}

	fmt.Println(direction)

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
