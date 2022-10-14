package maze

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"path/filepath"

	"context"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"net"
)

type V2Server struct {
}

// CreateSimulation - returns a new Simulation State
func (s *V2Server) CreateSimulation(
	ctx context.Context,
	req *CreateSimulationRequest,
) (*Simulation, error) {
	return NewSim()
}

// Simulate - Request an action on a Simulation
//            Returns the new Simulation State
func (s *V2Server) Simulate(
	ctx context.Context,
	req *SimulationAction,
) (*Simulation, error) {
	x, y, valid := req.Sim.DryMove(req.Action)
	if valid {
		req.Sim.Hero.X = x
		req.Sim.Hero.Y = y
	}
	return req.Sim, nil
}

// GetFeaturesV2 - Given a simulation return v2 features
func (s *V2Server) GetFeaturesV2(
	ctx context.Context,
	req *Simulation,
) (*FeaturesV2, error) {
	ret := &FeaturesV2{}
	ret.Features = getInputV2(req)
	return ret, nil
}

func (s *V2Server) mustEmbedUnimplementedSimulatorServer() {

}

func getInputV2(s *Simulation) []float64 {
	input := []float64{}

	v, err := s.Maze.Get(s.Hero.X, s.Hero.Y-1)
	if err != nil || v == WALL {
		input = append(input, 1)
	} else {
		input = append(input, 0)
	}
	v, err = s.Maze.Get(s.Hero.X, s.Hero.Y+1)
	if err != nil || v == WALL {
		input = append(input, 1)
	} else {
		input = append(input, 0)
	}

	v, err = s.Maze.Get(s.Hero.X-1, s.Hero.Y)
	if err != nil || v == WALL {
		input = append(input, 1)
	} else {
		input = append(input, 0)
	}

	v, err = s.Maze.Get(s.Hero.X+1, s.Hero.Y)
	if err != nil || v == WALL {
		input = append(input, 1)
	} else {
		input = append(input, 0)
	}

	// Normalizing all coordinates using (x - min)/(max - min)
	input = append(
		input,
		float64(s.Hero.X)/float64(50),
		float64(s.Hero.Y)/float64(50),
		float64(s.Maze.Exit.X)/float64(50),
		float64(s.Maze.Exit.Y)/float64(50),
		float64(s.Hero.X-s.Maze.Exit.X)/float64(50), // DX
		float64(s.Hero.Y-s.Maze.Exit.Y)/float64(50), // DY
	)
	return input
}

func distV2(hX, hY, eX, eY int32) float64 {
	return math.Pow(float64(hX-eX), 2) + math.Pow(float64(hY-eY), 2)
}

func getOutputV2(s *Simulation) ([]int32, error) {
	origDist := distV2(s.Hero.X, s.Hero.Y, s.Maze.Exit.X, s.Maze.Exit.Y)
	minDir := ""
	for _, dir := range []string{UP, DOWN, RIGHT, LEFT} {
		x, y, valid := s.DryMove(dir)
		if !valid {
			continue
		}

		newDist := distV2(x, y, s.Maze.Exit.X, s.Maze.Exit.Y)
		if newDist < origDist {
			minDir = dir
			continue
		}
		if newDist == origDist {
			// If both are equal choose one randomly
			if minDir == "" || rand.Intn(2) == 1 {
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

	for _, dir := range []string{UP, DOWN, RIGHT, LEFT} {
		if dir == minDir {
			ret = append(ret, 1)
		} else {
			ret = append(ret, 0)
		}
	}

	return ret, nil
}

func writeInputHeaderV2(fp *os.File) error {
	header := ""
	header += "up_wall,down_wall,left_wall,right_wall,hero_x,hero_y,exit_x,exit_y,dx,dy\n"

	_, err := fp.Write([]byte(header))
	return errors.Wrap(err, "failed write input header")
}

func writeOutputHeaderV2(fp *os.File) error {
	header := "up,down,right,left\n"

	_, err := fp.Write([]byte(header))
	return err
}

func writeOutputV2(path string, output []int32) error {
	// If the path doesnt exist, add the header
	outputPath := filepath.Join(path, "labels.csv")
	fp, err := os.OpenFile(outputPath, os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		fp, err = os.OpenFile(outputPath, os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			return errors.Wrap(err, "failed create label file")
		}

		if err := writeOutputHeaderV2(fp); err != nil {
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

func writeInputV2(path string, input []float64) error {
	// If the path doesnt exist, add the header
	inputPath := filepath.Join(path, "inputs.csv")
	fp, err := os.OpenFile(inputPath, os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		fp, err = os.OpenFile(inputPath, os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			return errors.Wrap(err, "failed create file")
		}

		if err := writeInputHeaderV2(fp); err != nil {
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

func GenerateTrainingDataV2(path string) error {
	cnt := 0
	for j := 0; j < 1000; j++ {
		s, err := NewSim()
		if err != nil {
			return err
		}

		for i := 0; i < 100; i++ {
			if (cnt+1)%10 == 0 {
				fmt.Printf("\r%d/%d", cnt+1, 1000*100)
			}
			input := getInputV2(s)
			output, err := getOutputV2(s)
			if err != nil {
				s.PlaceHero()
				cnt++
				continue
			}
			if err := writeInputV2(path, input); err != nil {
				return errors.Wrap(err, "writing input v1")
			}
			if err := writeOutputV2(path, output); err != nil {
				return errors.Wrap(err, "writing output v1")
			}
			s.PlaceHero()
			cnt++
		}
	}
	return nil
}

// startV2Server - starts the grpc server and never returns
func StartV2Server(port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		return errors.Wrap(err, "failed to listen")
	}

	server := &V2Server{}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	RegisterSimulatorServer(grpcServer, server)
	grpcServer.Serve(lis)
	return nil
}