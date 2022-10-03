package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"path/filepath"

	"github.com/param108/aimaze/maze"
	"github.com/param108/aimaze/sim"
	"github.com/pkg/errors"
)

func getInputV2(s *sim.Sim) []int {
	input := []int{}

	v, err := s.M.Get(s.H.X, s.H.Y - 1)
	if err != nil || v == maze.WALL {
		input = append(input, 1)
	} else {
		input = append(input, 0)
	}
	v, err = s.M.Get(s.H.X, s.H.Y + 1)
	if err != nil || v == maze.WALL {
		input = append(input, 1)
	} else {
		input = append(input, 0)
	}

	v, err = s.M.Get(s.H.X - 1, s.H.Y)
	if err != nil || v == maze.WALL {
		input = append(input, 1)
	} else {
		input = append(input, 0)
	}

	v, err = s.M.Get(s.H.X + 1, s.H.Y)
	if err != nil || v == maze.WALL {
		input = append(input, 1)
	} else {
		input = append(input, 0)
	}

	input = append(
		input,
		s.H.X,
		s.H.Y,
		s.M.E.X,
		s.M.E.Y,
		s.H.X - s.M.E.X, // DX
		s.H.Y - s.M.E.Y, // DY
	)
	return input
}

func distV2(hX, hY, eX, eY int) float64 {
	return math.Pow(float64(hX-eX), 2) + math.Pow(float64(hY-eY), 2)
}

func getOutputV2(s *sim.Sim) ([]int, error) {
	origDist := distV2(s.H.X, s.H.Y, s.M.E.X, s.M.E.Y)
	minDir := ""
	for _, dir := range []string{sim.UP, sim.DOWN, sim.RIGHT, sim.LEFT} {
		x, y, valid := s.DryMove(dir)
		if !valid {
			continue
		}

		newDist := distV2(x, y, s.M.E.X, s.M.E.Y)
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

	ret := []int{}

	for _, dir := range []string{sim.UP, sim.DOWN, sim.RIGHT, sim.LEFT} {
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

func writeOutputV2(path string, output []int) error {
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

func writeInputV2(path string, input []int) error {
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
		dataLine += fmt.Sprintf("%d", v)
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

func generateTrainingDataV2(path string) error {
	cnt := 0
	for j := 0; j < 1000; j++ {
		s, err := sim.NewSim()
		if err != nil {
			return err
		}

		for i := 0; i < 100; i++ {
			if (cnt + 1) %10 == 0 {
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
