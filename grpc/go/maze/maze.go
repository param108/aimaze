package maze

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
	"strings"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// NewMaze - creates a new maze with a size 50x50
// This creates an empty maze.
// You must call Create() to generate the maze
func NewMaze() *Maze {
	m := &Maze{}
	m.Size = &Size{
		Width: 50,
		Height: 50,
	}
	m.Exit = &Point{}

	m.Maze = strings.Repeat(" ", 50*50)
	m.DoorsPerWall = 25
	return m
}

// NewFromFile - Read a maze from a file
func NewFromFile(filename string) (*Maze, error) {
	bts, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	m := Maze{}
	err = json.Unmarshal(bts, &m)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

const WALL = "#"
const EMPTY = " "

// Get - get the character at (x,y)
// x and y are 0 indexed
func (m *Maze) Get(x, y int32) (string, error) {
	if x >= m.Size.Width || x < 0 {
		return "", errors.New("x out of bounds")
	}
	if y >= m.Size.Height || y < 0 {
		return "", errors.New("y out of bounds")
	}

	return string(m.Maze[y*m.Size.Width+x]), nil
}

// Set - make the position x,y the provided item
func (m *Maze) Set(x, y int32, item string) error {
	if x >= m.Size.Width || x < 0 {
		return errors.New("x out of bounds")
	}
	if y >= m.Size.Height || y < 0 {
		return errors.New("y out of bounds")
	}

	if y*m.Size.Width+x+1 > int32(len(m.Maze)) {
		m.Maze = m.Maze[:y*m.Size.Width+x] + item
	} else {
		m.Maze = m.Maze[:y*m.Size.Width+x] + item + m.Maze[y*m.Size.Width+x+1:]
	}

	return nil
}

// IsBorder - checks if the point is a border point
// or not.
func (m *Maze) IsBorder(x, y int32) (bool, error) {
	if x >= m.Size.Width || x < 0 {
		return false, errors.New("x out of bounds")
	}
	if y >= m.Size.Height || y < 0 {
		return false, errors.New("y out of bounds")
	}

	if x == 0 || x == (m.Size.Width-1) {
		return true, nil
	}

	if y == 0 || y == (m.Size.Height-1) {
		return true, nil
	}

	return false, nil
}

// Save - Save the maze as json
func (m *Maze) Save(filename string) error {
	bts, err := json.Marshal(m)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, bts, 0777)
	if err != nil {
		return err
	}

	return nil
}

// IsCorner - the point in question a corner ?
func (m *Maze) IsCorner(x, y int32) bool {
	isBorder, err := m.IsBorder(x, y)
	if err != nil {
		return false
	}
	if isBorder {
		if (x == 0 && y == 0) ||
			(x == 0 && y == (m.Size.Height-1)) ||
			(x == (m.Size.Width-1) && y == 0) ||
			(x == (m.Size.Width-1) && y == (m.Size.Height-1)) {
			return true
		}
	}
	return false
}

// createBorder - creates the border of the maze
// leaves one point empty and populates the exit.
func (m *Maze) createBorder() {
	exitFound := false
	for y := int32(0); y < m.Size.Height; y++ {
		for x := int32(0); x < m.Size.Width; x++ {
			if isBorder, _ := m.IsBorder(x, y); isBorder {
				// if we havent found an exit toss a random
				// number and if it is divisible by 2 then
				// lets choose this as the exit.
				if !exitFound && (rand.Intn(10) == 6) && !m.IsCorner(x, y) {
					exitFound = true
					m.Exit.X = x
					m.Exit.Y = y
					m.Set(x, y, EMPTY)
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
	for x := int32(1); x < m.Size.Width-1; x++ {
		// columns 1,3,5,7 etc were empty
		if x%2 != 0 {
			continue
		}

		foundSoFar := int32(0)

		// this is a wall column
		for y := int32(1); y < m.Size.Height-1; y++ {

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

				if y == m.Size.Height-2 {
					w, _ := m.Get(x, m.Size.Height-1)

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
	for y := int32(0); y < m.Size.Height; y++ {
		for x := int32(0); x < m.Size.Width; x++ {
			s, _ := m.Get(x, y)
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
