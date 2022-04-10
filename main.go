package main

import (
	"fmt"
	"log"
	"os"

	algo "github.com/lswarss/AAI/algo"
	distances "github.com/lswarss/AAI/files"
)

func main() {
	matrixes, err := distances.NewDistanceMatrixFromFS(os.DirFS("data"))
	if err != nil {
		log.Fatal(err)
	}

	bestIndividual, bestScore := algo.GeneticAlgorithm(matrixes[1], 250, 0.05, 100)

	fmt.Println("BI: ", bestIndividual)
	fmt.Println("BS: ", bestScore)
}
