package maze

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"

	"github.com/param108/aimaze/mazesim/spec/grpc/maze"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// NewMaze - creates a new maze with a size 50x50
// This creates an empty maze.
// You must call Create() to generate the maze
func NewMaze() *maze.Maze {
	m := &maze.Maze{}
	m.Size = &maze.Size{
		Width: 50,
		Height: 50,
	}
	m.Exit = &maze.Point{}

	m.Maze = strings.Repeat(" ", 50*50)
	m.DoorsPerWall = 25
	return m
}

// NewFromFile - Read a maze from a file
func NewFromFile(filename string) (*maze.Maze, error) {
	bts, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	m := maze.Maze{}
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
func Get(m *maze.Maze, x, y int32) (string, error) {
	if x >= m.Size.Width || x < 0 {
		return "", errors.New("x out of bounds")
	}
	if y >= m.Size.Height || y < 0 {
		return "", errors.New("y out of bounds")
	}

	return string(m.Maze[y*m.Size.Width+x]), nil
}

// Set - make the position x,y the provided item
func Set(m *maze.Maze, x, y int32, item string) error {
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
func IsBorder(m *maze.Maze, x, y int32) (bool, error) {
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
func Save(m *maze.Maze, filename string) error {
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
func IsCorner(m *maze.Maze, x, y int32) bool {
	isBorder, err := IsBorder(m, x, y)
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
func createBorder(m *maze.Maze) {
	exitIdx := rand.Intn(int(m.Size.Height*2 + m.Size.Width*2) - 4)
	for (exitIdx == 0 || // first corner
		exitIdx == int(m.Size.Width - 1) || // second corner
		exitIdx == int(m.Size.Width + 2*m.Size.Height) - 2 - 1 || // third corner
		exitIdx == int(m.Size.Height*2 + m.Size.Width*2) - 4 - 1) {
		exitIdx = rand.Intn(int(m.Size.Height*2 + m.Size.Width*2) - 4)
	}

	exitFound := false
	for y := int32(0); y < m.Size.Height; y++ {
		for x := int32(0); x < m.Size.Width; x++ {
			if isBorder, _ := IsBorder(m, x, y); isBorder {
				exitPlaced := false
				// Place the exit at the exitIdx'th point on the border
				if !exitFound {
					if exitIdx == 0 {
						exitPlaced = true
						exitFound = true
						m.Exit.X = x
						m.Exit.Y = y
						Set(m, x, y, EMPTY)
					} else {
						exitIdx --
					}
				}

				if !exitPlaced {
					Set(m, x, y, WALL)
				}
			}
		}
	}
}

// createWalls - Basically every alternate column
// 15 out of 50
func createWalls(m *maze.Maze) {
	for x := int32(1); x < m.Size.Width-1; x++ {
		// columns 1,3,5,7 etc were empty
		if x%2 != 0 {
			continue
		}

		foundSoFar := int32(0)

		// this is a wall column
		for y := int32(1); y < m.Size.Height-1; y++ {

			if rand.Intn(10)%2 == 0 && foundSoFar < m.DoorsPerWall {

				if x == 1 {
					w, _ := Get(m, 0, y)

					if w == EMPTY {
						// if we are in the second column
						// and the first column is the exit then
						// we cannot place a wall there otherwize
						// we will block the exit.
						continue
					}
				}

				if x == m.Size.Width-2 {
					w, _ := Get(m, x+1, y)

					if w == EMPTY {
						// if we are in the second last column
						// and the last column is the exit then
						// we cannot place a wall there otherwize
						// we will block the exit.
						continue
					}
				}

				if y == m.Size.Height-2 {
					w, _ := Get(m, x, m.Size.Height-1)

					if w == EMPTY {
						// if we are in the second last row of a column
						// and the last column is the exit then
						// we cannot place a wall there otherwize
						// we will block the exit.
						continue
					}
				}

				if y == 1 {
					w, _ := Get(m, x, 0)

					if w == EMPTY {
						// if we are in the second row of a column
						// and the first row is the exit then
						// we cannot place a wall there otherwize
						// we will block the exit.
						continue
					}
				}

				if y == m.Size.Height-2 {
					w, _ := Get(m, x, m.Size.Height-1)

					if w == EMPTY {
						// if we are in the second last row of a column
						// and the last column is the exit then
						// we cannot place a wall there otherwize
						// we will block the exit.
						continue
					}
				}
				foundSoFar++
				Set(m, x, y, WALL)
			}
		}
	}
}

func Print(m *maze.Maze) {
	for y := int32(0); y < m.Size.Height; y++ {
		for x := int32(0); x < m.Size.Width; x++ {
			s, _ := Get(m, x, y)
			if len(s) == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print(s)
			}
		}
		fmt.Println("")
	}
}

func Create(m *maze.Maze) error {
	// create the border
	createBorder(m)

	createWalls(m)

	return nil
}
