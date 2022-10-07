package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/param108/aimaze/grpc/go/maze"
)

func main() {
	err := maze.StartV2Server(9999)
	if err != nil {
		log.Fatal(err)
	}
}

func generateTrainingData(path string) {
	rand.Seed(time.Now().UnixNano())

	err := maze.GenerateTrainingDataV2(path)
	if err != nil {
		log.Printf("%+v\n", err)
	}

}
