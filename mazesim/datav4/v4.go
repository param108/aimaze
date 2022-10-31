package datav4

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"

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
	req.Sim.Step += 1
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

	for _, p := range s.Maze.Maze {
		if string(p) == mazemgr.WALL {
			input = append(input, float64(1))
		} else {
			input = append(input, 0)
		}
	}

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

	dx := float64(s.Hero.X)/float64(s.Maze.Size.Width) - float64(s.Maze.Exit.X)/float64(s.Maze.Size.Width)
	dy := float64(s.Hero.Y)/float64(s.Maze.Size.Height) - float64(s.Maze.Exit.Y)/float64(s.Maze.Size.Height)

	hx := float64(s.Hero.X) / float64(s.Maze.Size.Width)
	hy := float64(s.Hero.Y) / float64(s.Maze.Size.Height)

	ex := float64(s.Maze.Exit.X) / float64(s.Maze.Size.Width)
	ey := float64(s.Maze.Exit.Y) / float64(s.Maze.Size.Height)

	px := float64(s.Prev.X) / float64(s.Maze.Size.Width)
	py := float64(s.Prev.Y) / float64(s.Maze.Size.Height)

	dh := (float64(s.Hero.X) + float64(s.Hero.Y)*float64(s.Maze.Size.Width)) / 2500
	de := (float64(s.Maze.Exit.X) + float64(s.Maze.Exit.Y)*float64(s.Maze.Size.Width)) / 2500
	dp := (float64(s.Prev.X) + float64(s.Prev.Y)*float64(s.Maze.Size.Width)) / 2500

	input = append(
		input,
		hx,
		hy,
		ex,
		ey,
		px,
		py,
		// Add the crosses for hero, exit and prev now
		dh,
		de,
		dp,
		dx,
		dy,
		dx*dy)
	return input
}

func dist(hX, hY, eX, eY int32) float64 {
	return math.Pow(float64(hX-eX), 2) + math.Pow(float64(hY-eY), 2)
}

//////////////////////////////////////////////////////////////////////////////////////////////////
// func writeFiles(																			    //
// 	outputPath string,																		    //
// 	m *maze.Maze,																			    //
// 	x, y int32,																				    //
// 	prev_x, prev_y int32,																	    //
// 	direction string,																		    //
// 	depth int32,																			    //
// ) error {																				    //
// 	err := writeInput(outputPath, getInput(&maze.Simulation{								    //
// 		Maze: m,																			    //
// 		Hero: &maze.Point{																	    //
// 			X: x,																			    //
// 			Y: y,																			    //
// 		},																					    //
// 		Prev: &maze.Point{																	    //
// 			X: prev_x,																		    //
// 			Y: prev_y,																		    //
// 		},																					    //
// 		Step: depth,																		    //
// 	}))																						    //
// 																							    //
// 	if err != nil {																			    //
// 		return err																			    //
// 	}																						    //
// 	ret := []int32{}																		    //
// 																							    //
// 	for _, dir := range []string{mazemgr.UP, mazemgr.DOWN, mazemgr.RIGHT, mazemgr.LEFT} {	    //
// 		if dir == direction {																    //
// 			ret = append(ret, 1)															    //
// 		} else {																			    //
// 			ret = append(ret, 0)															    //
// 		}																					    //
// 	}																						    //
// 																							    //
// 	err = writeOutput(outputPath, ret)														    //
// 	if err != nil {																			    //
// 		return err																			    //
// 	}																						    //
// 	return nil																				    //
// }																						    //
//////////////////////////////////////////////////////////////////////////////////////////////////

type recurse struct {
	foundRecursion map[int32]bool
	totalSteps     int32
}

// bfs - tracks bfs progress
type bfs struct {
	totalSteps int32
	seen       map[int32]bool
	work       []*bfsWork
	sim        *maze.Simulation
	mx         sync.Mutex
}

// newBFS
func newBFS(sim *maze.Simulation) *bfs {
	return &bfs{
		seen: map[int32]bool{
			sim.Hero.X + (sim.Maze.Size.Width*sim.Hero.Y): true,
		},
		work: []*bfsWork{{points: []*maze.Point{{X: sim.Hero.X, Y: sim.Hero.Y}}}},
		sim:  sim,
	}
}

// checkUpdateSeen - return true if point is already seen.
// if not seen set seen to true and return false
func (b *bfs) checkUpdateSeen(x, y int32) bool {
	b.mx.Lock()
	defer b.mx.Unlock()
	key := x + (y*b.sim.Maze.Size.Width)
	if b.seen[key] {
		return true
	}

	b.seen[key] = true

	return false
}

// addWork - add work to the bfs work queue
func (b *bfs) addWork(w *bfsWork) {
	b.mx.Lock()
	defer b.mx.Unlock()
	b.work = append(b.work, w)
}

// bfsWork - the work that is being tried
type bfsWork struct {
	// points - list of points traversed so far
	// current point is at the end
	points []*maze.Point

	// direction - direction to try now
	directions []string
}

func validPoint(m *maze.Maze, X, Y int32) bool {
	s, err := mazemgr.Get(m, X, Y)
	if err != nil {
		return false
	}
	if s == mazemgr.WALL {
		return false
	}
	return true
}

func (b *bfs) run() (*bfsWork, bool) {
	var wg sync.WaitGroup
	doneChan := make(chan *bfsWork, 2)
	success := false
	for {
		workList := b.work
		if len(workList) == 0 {
			close(doneChan)
			break
		} else {
			if debugRun {
				fmt.Println("work:", len(workList))
			}
		}

		b.work = []*bfsWork{}
		for i:=0; i < len(workList); i++ {
			b.totalSteps += 1
			wg.Add(1)

			go func(work *bfsWork) {
				defer wg.Done()
				ret, done := b.doWork(work)
				if done {
					doneChan <- ret
					success = true
				}
			}(workList[i])

		}

		wg.Wait()

		if b.totalSteps > 150000 {
			if debugRun {
				fmt.Println("Too many steps")
			}
			close(doneChan)
			break
		}

		if success {
			close(doneChan)
			break
		}
	}

	w := <-doneChan
	if w != nil {
		return w, true
	}

	return w, false
}

var debugRun = false

func (b *bfs) doWork(w *bfsWork) (*bfsWork, bool) {

	if len(w.points) > 500 {
		if debugRun {
			fmt.Println("Too many points", len(w.points))
		}
		return nil, false
	}

	// Randomize the array of directions as otherwise we bias the
	// neural net to always go up more frequently
	directions := []string{
		mazemgr.UP,
		mazemgr.DOWN,
		mazemgr.RIGHT,
		mazemgr.LEFT,
	}

	//rand.Shuffle(len(directions), func(i, j int) { directions[i], directions[j] = directions[j], directions[i] })

	var (
		x, y int32
	)

	x = w.points[len(w.points)-1].X
	y = w.points[len(w.points)-1].Y

	for _, direction := range directions {
		newX := x
		newY := y

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

		if validPoint(b.sim.Maze, newX, newY) {
			m := b.sim.Maze

			if !b.checkUpdateSeen(newX, newY) {
				newDirections := make([]string, len(w.directions))
				newPoints := make([]*maze.Point, len(w.points))

				copy(newDirections, w.directions)
				copy(newPoints, w.points)

				newDirections = append(newDirections, direction)

				p := &maze.Point{X: newX, Y: newY}

				newPoints = append(newPoints, p)
				p = nil
				if newX == m.Exit.X && newY == m.Exit.Y {
					// Found the exit return the last direction turned.
					return &bfsWork{
						directions: newDirections,
						points:     newPoints,
					}, true
				} else {
					if debugRun {
						fmt.Println("adding", newX, newY)
					}

					b.addWork(&bfsWork{
						directions: newDirections,
						points:     newPoints,
					})
				}
			} else {
				if debugRun {
					fmt.Println("seen", newX, newY)
				}
			}
		} else {
			if debugRun {
					fmt.Println("invalid", newX, newY)
			}
		}
	}

	return nil, false
}

// recursion - recurse randomly towards the exit
// if you can move towards
func (r *recurse) recursion(x, y int32, prev_x, prev_y int32, m *maze.Maze, outputPath string, depth int32) ([]*maze.Point, bool) {

	if r.totalSteps > 150000 {
		return nil, false
	}

	// we exit before setting foundRecursion because if it
	// returns here this node has not disqualified itself as
	// a useful node. i.e you could possibly still reach
	// the exit from this node if you came at it from some other
	// route.
	if depth == 150 {
		return nil, false
	}

	if r.foundRecursion == nil {
		r.foundRecursion = map[int32]bool{}
	}

	thisKey := x*100 + y
	r.foundRecursion[thisKey] = true

	// if you start from some other starting point
	// this node is still valid
	// BUT this explodes the complexity so ignoring for now
	// If the training doesnt work, we will uncomment this
	defer func() {
		r.foundRecursion[thisKey] = false
	}()

	newX := x
	newY := y

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

		if validPoint(m, newX, newY) {
			key := newX*100 + newY
			if !r.foundRecursion[key] {
				if newX == m.Exit.X && newY == m.Exit.Y {
					// Found the exit return the last direction turned.
					return []*maze.Point{{X: m.Exit.X, Y: m.Exit.Y}}, true
				} else {
					r.totalSteps += 1
					pointList, success := r.recursion(newX, newY, x, y, m, outputPath, depth+1)
					if !success {
						continue
					} else {
						// Append the direction taken (which led to success) to
						// the beginning of the returned list of directions.
						// This way the returned array is in the order of directions
						// taken.
						return append([]*maze.Point{{X: newX, Y: newY}}, pointList...), true
					}
				}
			}
		}
	}

	//should reach here only if no moves are valid
	return nil, false

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
	for i := 0; i < 2500; i++ {
		header += fmt.Sprintf("x%d,", i+1)
	}
	header += "up,down,left,right,hero_x,hero_y,exit_x,exit_y,prev_x,prev_y,cross_hero,cross_exit,cross_prev,dx,dy,cross_delta\n"

	_, err := fp.Write([]byte(header))
	return errors.Wrap(err, "failed write input header")
}

func writeOutputHeader(fp *os.File) error {
	header := ""
	for i := 0; i < 2500; i++ {
		header += fmt.Sprintf("x_%04d,", i)
	}

	header = header[:len(header)-1] + "\n"

	_, err := fp.Write([]byte(header))
	return err
}

var globalOutputFp *os.File
var globalInputFp *os.File

func writeOutput(path string, output []int32) ([]byte, error) {
	if globalOutputFp == nil {
		writemx.Lock()
		defer writemx.Unlock()

		if globalOutputFp == nil {
			// If the path doesnt exist, add the header
			outputPath := filepath.Join(path, "labels.csv")
			fp, err := os.OpenFile(outputPath, os.O_APPEND|os.O_RDWR, 0644)
			if err != nil {
				fp, err = os.OpenFile(outputPath, os.O_CREATE|os.O_RDWR, 0644)
				if err != nil {
					return nil, errors.Wrap(err, "failed create label file")
				}

				if err := writeOutputHeader(fp); err != nil {
					fp.Close()
					return nil, errors.Wrap(err, "failed write label line")
				}
				globalOutputFp = fp

			}
		}
	}

	dataLine := ""
	for ix, v := range output {
		dataLine += fmt.Sprintf("%d", v)
		if ix != len(output)-1 {
			dataLine += ","
		}
	}
	dataLine += "\n"
	return []byte(dataLine), nil
}

func writeInput(path string, input []float64) ([]byte, error) {
	if globalInputFp == nil {
		writemx.Lock()
		defer writemx.Unlock()

		if globalInputFp == nil {
			// If the path doesnt exist, add the header
			inputPath := filepath.Join(path, "inputs.csv")
			fp, err := os.OpenFile(inputPath, os.O_APPEND|os.O_RDWR, 0644)
			if err != nil {
				fp, err = os.OpenFile(inputPath, os.O_CREATE|os.O_RDWR, 0644)
				if err != nil {
					return nil, errors.Wrap(err, "failed create file")
				}

				if err := writeInputHeader(fp); err != nil {
					fp.Close()
					return nil, errors.Wrap(err, "failed write input header")
				}
				globalInputFp = fp
			}
		}
	}

	dataLine := ""
	for ix, v := range input {
		dataLine += fmt.Sprintf("%f", v)
		if ix != len(input)-1 {
			dataLine += ","
		}
	}
	dataLine += "\n"
	return []byte(dataLine), nil
}

func convertToInts(directions []string) []int32 {
	ret := []int32{}

	for _, direction := range directions {
		for _, dir := range []string{mazemgr.UP, mazemgr.DOWN, mazemgr.RIGHT, mazemgr.LEFT} {
			if dir == direction {
				ret = append(ret, 1)
			} else {
				ret = append(ret, 0)
			}
		}
	}

	for i := len(directions); i < 150; i++ {
		ret = append(ret, 0, 0, 0, 0) // add noops for the rest of the moves
	}

	return ret
}

func convertPointsToInts(s *maze.Simulation, points []*maze.Point) []int32 {
	ret := make([]int32, 2500)
	for _, pt := range points {
		ret[int(pt.X+s.Maze.Size.Width*pt.Y)] = 1
	}

	return ret
}

var writemx sync.Mutex

type Data struct {
	in  []byte
	out []byte
}

func GenerateTrainingData(pathString string) error {
	return GenerateTrainingDataBFS(pathString)
}

func GenerateTrainingDataBFS(pathString string) error {
	cnt := atomic.Int32{}
	errCnt := atomic.Int32{}

	data := &Data{}
	wgWriter := sync.WaitGroup{}

	dataInCh := make(chan *Data, 30)
	wgWriter.Add(1)
	go func(path string) {
		defer wgWriter.Done()
		for {
			select {
			case d := <-dataInCh:
				if d == nil {
					return
				}
				writemx.Lock()
				globalInputFp.Write(d.in)
				globalOutputFp.Write(d.out)
				writemx.Unlock()
			}
		}
	}(pathString)

	OuterSize := 30000
	InnerSize := 10
	for i := 0; i < OuterSize; i++ {

		s, err := mazemgr.NewSim()
		if err != nil {
			panic(err)
		}

		for j := 0; j < InnerSize; j++ {
			cnt.Add(1)
			if !debugRun {
				cntVal := cnt.Load()
				if (cntVal+1)%10 == 0 {
					fmt.Printf("\r%d/%d err: %d", cntVal+1, OuterSize*InnerSize, errCnt.Load())
				}
			}

			mazemgr.PlaceHero(s)
			d := newBFS(s)
			if debugRun {
				fmt.Println("New BFS")
			}
			w, done := d.run()
			if done {

				inp, err := writeInput(pathString, getInput(s))
				if err != nil {
					panic(err)
				}
				data.in = append(data.in, inp...)
				out, err := writeOutput(pathString, convertPointsToInts(s, w.points))
				if err != nil {
					panic(err)
				}
				data.out = append(data.out, out...)

				if len(data.in) > 1000000 {
					dataInCh <- data
					data = &Data{}
				}

			} else {
				errCnt.Add(1)
				if debugRun {
					mazemgr.PrintSimSeen(s, d.seen)
					return nil
				}
				return nil
			}
		}
	}

	dataInCh <- data

	close(dataInCh)
	wgWriter.Wait()

	if globalOutputFp != nil {
		globalOutputFp.Close()
	}

	if globalInputFp != nil {
		globalInputFp.Close()
	}

	return nil
}

func GenerateTrainingDataDepthFirst(pathString string) error {
	cnt := atomic.Int32{}
	errCnt := atomic.Int32{}
	wg := sync.WaitGroup{}
	wgWriter := sync.WaitGroup{}

	dataInCh := make(chan *Data, 30)
	wgWriter.Add(1)
	go func(path string) {
		defer wgWriter.Done()
		for {
			select {
			case d := <-dataInCh:
				if d == nil {
					return
				}
				writemx.Lock()
				globalInputFp.Write(d.in)
				globalOutputFp.Write(d.out)
				writemx.Unlock()
			}
		}
	}(pathString)

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(path string) {
			data := &Data{}
			defer wg.Done()
			for j := 0; j < 6250; j++ {
				s, err := mazemgr.NewSim()
				if err != nil {
					panic(err)
				}

				for i := 0; i < 10; i++ {
					cntVal := cnt.Load()
					if (cntVal+1)%10 == 0 {
						fmt.Printf("\r%d/%d err: %d", cntVal+1, 4*6250*10, errCnt.Load())
					}

					r := recurse{}
					// find the path
					// Initially s.Prev will be the same as s.Hero
					if pointList, success := r.recursion(s.Hero.X, s.Hero.Y, s.Prev.X, s.Prev.Y, s.Maze, path, 0); success {
						inp, err := writeInput(path, getInput(s))
						if err != nil {
							panic(err)
						}
						data.in = append(data.in, inp...)
						out, err := writeOutput(path, convertPointsToInts(s, pointList))
						if err != nil {
							panic(err)
						}
						data.out = append(data.out, out...)

						if len(data.in) > 1000000 {
							dataInCh <- data
							data = &Data{}
						}

					} else {
						errCnt.Add(1)
					}

					// run it again sam
					mazemgr.PlaceHero(s)
					cnt.Add(1)
				}
			}

			if len(data.out) > 0 {
				dataInCh <- data
			}

		}(pathString)
	}

	wg.Wait()
	close(dataInCh)
	wgWriter.Wait()

	if globalOutputFp != nil {
		globalOutputFp.Close()
	}

	if globalInputFp != nil {
		globalInputFp.Close()
	}

	return nil
}
