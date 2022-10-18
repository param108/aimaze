package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/param108/aimaze/mazesim/datav3"
	"github.com/param108/aimaze/mazesim/spec/grpc/maze"
)

func main() {
	cmd := os.Args[1]

	switch (cmd) {
	case "serve":
		s := &datav3.Server{}
		err := maze.StartServer(9999, s)
		if err != nil {
			log.Fatal(err)
		}
	case "generate":
		path := os.Args[2]
		generateTrainingData(path)
	}

}

func generateTrainingData(path string) {
	rand.Seed(time.Now().UnixNano())

	err := datav3.GenerateTrainingData(path)
	if err != nil {
		log.Printf("%+v\n", err)
	}

}
