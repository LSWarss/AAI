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

	// jobs := make(chan distances.DistanceMatrix, 5)
	// results := make(chan algo.TSPResult, 5)

	// for i := 0; i < 5; i++ {
	// 	go worker(jobs, results)
	// }

	// for j := 0; j < 5; j++ {
	// 	jobs <- matrixes[1]
	// }
	// close(jobs)

	// for r := 0; r < 5; r++ {
	// 	result := <-results

	// 	fmt.Println("BI", result.BestIndividual)
	// 	fmt.Println("BS", result.BestScore)

	// 	if result.BestScore < bestOfAll.BestScore {
	// 		bestOfAll = result
	// 	}
	// }
	// close(results)
	bestOfAll := algo.GeneticAlgorithm(matrixes[4], 250, 0.05, 500000)
	fmt.Println("--------")
	fmt.Println("BOA - BI", bestOfAll.BestIndividual)
	fmt.Println("BOA - BS", bestOfAll.BestScore)
}

func worker(jobs <-chan distances.DistanceMatrix, results chan<- algo.TSPResult) {
	for n := range jobs {
		results <- algo.GeneticAlgorithm(n, 250, 0.05, 500000)
	}
}
