package datav3

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"path/filepath"

	"context"

	mazemgr "github.com/param108/aimaze/mazesim/maze"
	"github.com/param108/aimaze/mazesim/spec/grpc/maze"
	"github.com/pkg/errors"
)

type Server struct {
}

// CreateSimulation - returns a new Simulation State
func (s *Server) CreateSimulation(
	ctx context.Context,
	req *maze.CreateSimulationRequest,
) (*maze.Simulation, error) {
	return mazemgr.NewSim()
}

// Simulate - Request an action on a Simulation
//
//	Returns the new Simulation State
func (s *Server) Simulate(
	ctx context.Context,
	req *maze.SimulationAction,
) (*maze.Simulation, error) {
	x, y, valid := mazemgr.DryMove(req.Sim, req.Action)
	if valid {
		req.Sim.Prev.X = req.Sim.Hero.X
		req.Sim.Prev.Y = req.Sim.Hero.Y
		req.Sim.Hero.X = x
		req.Sim.Hero.Y = y
	}
	return req.Sim, nil
}

// GetFeaturesV2 - Given a simulation return v2 features
func (s *Server) GetFeaturesV2(
	ctx context.Context,
	req *maze.Simulation,
) (*maze.FeaturesV2, error) {
	ret := &maze.FeaturesV2{}
	ret.Features = getInput(req)
	return ret, nil
}

func getInput(s *maze.Simulation) []float64 {
	input := []float64{}

	v, err := mazemgr.Get(s.Maze, s.Hero.X, s.Hero.Y-1)
	if err != nil || v == mazemgr.WALL {
		input = append(input, 1)
	} else {
		input = append(input, 0)
	}
	v, err = mazemgr.Get(s.Maze, s.Hero.X, s.Hero.Y+1)
	if err != nil || v == mazemgr.WALL {
		input = append(input, 1)
	} else {
		input = append(input, 0)
	}

	v, err = mazemgr.Get(s.Maze, s.Hero.X-1, s.Hero.Y)
	if err != nil || v == mazemgr.WALL {
		input = append(input, 1)
	} else {
		input = append(input, 0)
	}

	v, err = mazemgr.Get(s.Maze, s.Hero.X+1, s.Hero.Y)
	if err != nil || v == mazemgr.WALL {
		input = append(input, 1)
	} else {
		input = append(input, 0)
	}

	/*for _, p := range s.Maze.Maze {
		if string(p) == mazemgr.WALL {
			input = append(input, float64(1)/2500)
		} else {
			input = append(input, 0)
		}
	}*/

	input = append(
		input,
		float64(s.Hero.X)/float64(s.Maze.Size.Width),
		float64(s.Hero.Y)/float64(s.Maze.Size.Height),
		float64(s.Maze.Exit.X)/float64(s.Maze.Size.Width),
		float64(s.Maze.Exit.Y)/float64(s.Maze.Size.Height),
		float64(s.Prev.X)/float64(s.Maze.Size.Width),
		float64(s.Prev.Y)/float64(s.Maze.Size.Height),
		// Add the crosses for hero, exit and prev now
		float64(s.Hero.X)/float64(s.Maze.Size.Width)*
			float64(s.Hero.Y)/float64(s.Maze.Size.Height),
		float64(s.Maze.Exit.X)/float64(s.Maze.Size.Width)*
			float64(s.Maze.Exit.Y)/float64(s.Maze.Size.Height),
		float64(s.Prev.X)/float64(s.Maze.Size.Width)*
			float64(s.Prev.Y)/float64(s.Maze.Size.Height),
	)
	return input
}

func dist(hX, hY, eX, eY int32) float64 {
	return math.Pow(float64(hX-eX), 2) + math.Pow(float64(hY-eY), 2)
}

func writeFiles(outputPath string, m *maze.Maze, x, y int32, prev_x, prev_y int32, direction string) error {
	err := writeInput(outputPath, getInput(&maze.Simulation{
		Maze: m,
		Hero: &maze.Point{
			X: x,
			Y: y,
		},
		Prev: &maze.Point{
			X: prev_x,
			Y: prev_y,
		},
	}))

	if err != nil {
		return err
	}
	ret := []int32{}

	for _, dir := range []string{mazemgr.UP, mazemgr.DOWN, mazemgr.RIGHT, mazemgr.LEFT} {
		if dir == direction {
			ret = append(ret, 1)
		} else {
			ret = append(ret, 0)
		}
	}

	err = writeOutput(outputPath, ret)
	if err != nil {
		return err
	}
	return nil
}

var foundRecursion map[int32]bool
var totalSteps = 0

// recursion - recurse randomly towards the exit
// if you can move towards
func recursion(x, y int32, prev_x, prev_y int32, m *maze.Maze, outputPath string, depth int) bool {

	if totalSteps > 150000 {
		return false
	}

	// we exit before setting foundRecursion because if it
	// returns here this node has not disqualified itself as
	// a useful node. i.e you could possibly still reach
	// the exit from this node if you came at it from some other
	// route.
	if depth == 150 {
		return false
	}

	if foundRecursion == nil {
		foundRecursion = map[int32]bool{}
	}

	thisKey := x*100 + y
	foundRecursion[thisKey] = true

	// if you start from some other starting point
	// this node is still valid
	// BUT this explodes the complexity so ignoring for now
	// If the training doesnt work, we will uncomment this
	defer func() {
		foundRecursion[thisKey] = false
	}()

	newX := x
	newY := y

	valid := func(X, Y int32) bool {
		s, err := mazemgr.Get(m, X, Y)
		if err != nil {
			return false
		}
		if s == mazemgr.WALL {
			return false
		}
		return true
	}

	alreadySeen := map[string]bool{}

	a := []string{}

	if m.Exit.X > x {
		a = append(a, mazemgr.RIGHT)
		alreadySeen[mazemgr.RIGHT] = true
	} else if m.Exit.X < x {
		a = append(a, mazemgr.LEFT)
		alreadySeen[mazemgr.LEFT] = true
	}

	if m.Exit.Y < y {
		a = append(a, mazemgr.UP)
		alreadySeen[mazemgr.UP] = true
	} else if m.Exit.Y > y {
		a = append(a, mazemgr.DOWN)
		alreadySeen[mazemgr.DOWN] = true
	}

	// Randomize the array of directions as otherwise we bias the
	// neural net to always go up more frequently
	directions := []string{
		mazemgr.UP,
		mazemgr.DOWN,
		mazemgr.RIGHT,
		mazemgr.LEFT,
	}

	rand.Shuffle(len(directions), func(i, j int) { directions[i], directions[j] = directions[j], directions[i] })

	for _, dir := range directions {
		if !alreadySeen[dir] {
			a = append(a, dir)
		}
	}

	if len(a) > 4 {
		panic(fmt.Sprintf("too many %v", a))
	}

	for _, direction := range a {

		switch direction {
		case mazemgr.UP:
			newY = y - 1
		case mazemgr.DOWN:
			newY = y + 1
		case mazemgr.RIGHT:
			newX = x + 1
		case mazemgr.LEFT:
			newX = x - 1
		}

		if valid(newX, newY) {
			key := newX*100 + newY
			if !foundRecursion[key] {
				if newX == m.Exit.X && newY == m.Exit.Y {
					if err := writeFiles(outputPath, m, x, y, prev_x, prev_y, direction); err != nil {
						panic(err)
					}
					return true
				} else {
					totalSteps += 1
					if !recursion(newX, newY, x, y, m, outputPath, depth+1) {
						continue
					} else {
						if err := writeFiles(outputPath, m, x, y, prev_x, prev_y, direction); err != nil {
							panic(err)
						}
						return true
					}
				}
			}
		}
	}

	//should reach here only if no moves are valid
	return false

}

func getOutput(s *maze.Simulation) ([]int32, error) {
	origDist := dist(s.Hero.X, s.Hero.Y, s.Maze.Exit.X, s.Maze.Exit.Y)
	minDir := ""
	for _, dir := range []string{mazemgr.UP, mazemgr.DOWN, mazemgr.RIGHT, mazemgr.LEFT} {
		x, y, valid := mazemgr.DryMove(s, dir)
		if !valid {
			continue
		}

		newDist := dist(x, y, s.Maze.Exit.X, s.Maze.Exit.Y)
		if newDist < origDist {
			minDir = dir
			continue
		}
		if newDist == origDist {
			// If both are equal choose one randomly
			if minDir == "" && rand.Intn(2) == 1 {
				minDir = dir
			}
		}
	}

	// we couldnt choose a direction
	// lets skip this datapoint
	if minDir == "" {
		return nil, errors.New("cant find dir")
	}

	ret := []int32{}

	for _, dir := range []string{mazemgr.UP, mazemgr.DOWN, mazemgr.RIGHT, mazemgr.LEFT} {
		if dir == minDir {
			ret = append(ret, 1)
		} else {
			ret = append(ret, 0)
		}
	}

	return ret, nil
}

func writeInputHeader(fp *os.File) error {
	header := ""
	/*for i := 0; i < 2500; i++ {
		header += fmt.Sprintf("x%d,", i+1)
	}*/
	header += "up,down,left,right,hero_x,hero_y,exit_x,exit_y,prev_x,prev_y,cross_hero,cross_exit,cross_prev\n"

	_, err := fp.Write([]byte(header))
	return errors.Wrap(err, "failed write input header")
}

func writeOutputHeader(fp *os.File) error {
	header := "up,down,right,left\n"

	_, err := fp.Write([]byte(header))
	return err
}

func writeOutput(path string, output []int32) error {
	// If the path doesnt exist, add the header
	outputPath := filepath.Join(path, "labels.csv")
	fp, err := os.OpenFile(outputPath, os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		fp, err = os.OpenFile(outputPath, os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			return errors.Wrap(err, "failed create label file")
		}

		if err := writeOutputHeader(fp); err != nil {
			fp.Close()
			return errors.Wrap(err, "failed write label line")
		}
	}
	defer fp.Close()

	dataLine := ""
	for ix, v := range output {
		dataLine += fmt.Sprintf("%d", v)
		if ix != len(output)-1 {
			dataLine += ","
		}
	}
	dataLine += "\n"
	if _, err := fp.Write([]byte(dataLine)); err != nil {
		return errors.Wrap(err, "failed write label line")
	}
	return nil
}

func writeInput(path string, input []float64) error {
	// If the path doesnt exist, add the header
	inputPath := filepath.Join(path, "inputs.csv")
	fp, err := os.OpenFile(inputPath, os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		fp, err = os.OpenFile(inputPath, os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			return errors.Wrap(err, "failed create file")
		}

		if err := writeInputHeader(fp); err != nil {
			fp.Close()
			return errors.Wrap(err, "failed write input header")
		}
	}
	defer fp.Close()

	dataLine := ""
	for ix, v := range input {
		dataLine += fmt.Sprintf("%f", v)
		if ix != len(input)-1 {
			dataLine += ","
		}
	}
	dataLine += "\n"
	if _, err := fp.Write([]byte(dataLine)); err != nil {
		return errors.Wrap(err, "failed write input line")
	}
	return nil
}

func GenerateTrainingData(path string) error {
	cnt := 0
	errCnt := 0
	for j := 0; j < 1000; j++ {
		s, err := mazemgr.NewSim()
		if err != nil {
			return err
		}

		for i := 0; i < 10; i++ {
			if (cnt+1)%10 == 0 {
				fmt.Printf("\r%d/%d err: %d", cnt+1, 1000*10, errCnt)
			}

			// reset the found counter
			foundRecursion = nil
			totalSteps = 0
			// find the path
			// Initially s.Prev will be the same as s.Hero
			if !recursion(s.Hero.X, s.Hero.Y, s.Prev.X, s.Prev.Y, s.Maze, path, 0) {
				errCnt++
			}

			// run it again sam
			mazemgr.PlaceHero(s)
			cnt++
		}
	}
	return nil
}
