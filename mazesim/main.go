package main

import (
	"math/rand"
	"os"
	"time"
	"log"
)

func main() {
	generateTrainingData(os.Args[1])
}

func generateTrainingData(path string) {
	rand.Seed(time.Now().UnixNano())

	err := generateTrainingDataV1(path)
	if err != nil {
		log.Printf("%+v\n", err)
	}

}
