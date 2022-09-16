package maze

import (
	"errors"
	"math/rand"
	"fmt"
)

type Size struct {
	Width  int
	Height int
}

type Point struct {
	X int
	Y int
}

type Maze struct {
	// S - the size of the maze
	S Size
	// M - the maze as raw bytes on the screen
	M []string // The maze
	// E - the position of the exit
	E Point
	// DoorsPerWall - the number of gaps in wall
	DoorsPerWall int
}

// NewMaze - creates a new maze with a size 50x50
// This creates an empty maze.
// You must call Create() to generate the maze
func NewMaze() *Maze {
	return &Maze{
		S: Size{Width: 50, Height: 50},
		M: make([]string, 2500),
		DoorsPerWall: 25,
	}
}

const WALL = "#"
const EMPTY = " "

// Get - get the character at (x,y)
// x and y are 0 indexed
func (m *Maze) Get(x, y int) (string, error) {
	if x >= m.S.Width || x < 0 {
		return "", errors.New("x out of bounds")
	}
	if y >= m.S.Height || y < 0 {
		return "", errors.New("y out of bounds")
	}

	return m.M[y*m.S.Width+x], nil
}

// Set - make the position x,y the provided item
func (m *Maze) Set(x, y int, item string) error {
	if x >= m.S.Width || x < 0 {
		return errors.New("x out of bounds")
	}
	if y >= m.S.Height || y < 0 {
		return errors.New("y out of bounds")
	}

	m.M[y*m.S.Width+x] = item
	return nil
}

// IsBorder - checks if the point is a border point
// or not.
func (m *Maze) IsBorder(x, y int) (bool, error) {
	if x >= m.S.Width || x < 0 {
		return false, errors.New("x out of bounds")
	}
	if y >= m.S.Height || y < 0 {
		return false, errors.New("y out of bounds")
	}

	if x == 0 || x == (m.S.Width-1) {
		return true, nil
	}

	if y == 0 || y == (m.S.Height-1) {
		return true, nil
	}

	return false, nil
}

// createBorder - creates the border of the maze
// leaves one point empty and populates the exit.
func (m *Maze) createBorder() {
	exitFound := false
	for y := 0; y < m.S.Height; y++ {
		for x := 0; x < m.S.Width; x++ {
			if isBorder, _ := m.IsBorder(x, y); isBorder {
				// if we havent found an exit toss a random
				// number and if it is divisible by 2 then
				// lets choose this as the exit.
				if !exitFound && rand.Intn(10)%2 == 0 {
					exitFound = true
					m.E.X = x
					m.E.Y = y
				} else {
					m.Set(x, y, WALL)
				}
			}
		}
	}
}

// createWalls - Basically every alternate column
// 15 out of 50
func (m *Maze) createWalls() {
	for x := 1; x < m.S.Width - 1; x++ {
		// columns 1,3,5,7 etc were empty
		if x % 2 != 0 {
			continue
		}

		foundSoFar := 0

		// this is a wall column
		for y := 1; y < m.S.Height - 1; y++ {

			if rand.Intn(10)%2 == 0 && foundSoFar < m.DoorsPerWall {
				if y == 1 {
					w, _ := m.Get(x, 0)

					if w == EMPTY {
						// if we are in the second row of a column
						// and the first column is the exit then
						// we cannot place a wall there otherwize
						// we will block the exit.
						continue
					}
				}

				if y == m.S.Height - 2 {
					w, _ := m.Get(x, m.S.Height - 1)

					if w == EMPTY {
						// if we are in the second last row of a column
						// and the last column is the exit then
						// we cannot place a wall there otherwize
						// we will block the exit.
						continue
					}
				}
				foundSoFar++
				m.Set(x, y, WALL)
			}
		}
	}
}

func (m *Maze) Print() {
	for y := 0; y < m.S.Height; y++ {
		for x := 0; x < m.S.Width; x++ {
			s,_ := m.Get(x, y)
			if len(s) == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print(s)
			}
		}
		fmt.Println("")
	}
}

func (m *Maze) Create() error {
	// create the border
	m.createBorder()

	m.createWalls()

	return nil
}