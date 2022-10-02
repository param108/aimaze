package main

import (
	"fmt"

	"github.com/param108/aimaze/sim"
)


func main() {
	s, err := sim.NewSim()

	if err != nil {
		panic(err)
	}

	s.Print()

	fmt.Println("Saved")

	err = s.Save("saveit.json")
	if err != nil {
		panic(err)
	}

	s,err = sim.LoadSim("saveit.json")
	if err != nil {
		panic(err)
	}

	s.Print()


}
