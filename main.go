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

	var bestOfAll algo.TSPResult = algo.TSPResult{
		BestIndividual: []int{},
		BestScore:      5000000000,
	}

	jobs := make(chan distances.DistanceMatrix, 100)
	results := make(chan algo.TSPResult, 100)

	for i := 0; i < 10; i++ {
		go worker(jobs, results)
	}

	for j := 0; j < 100; j++ {
		jobs <- matrixes[1]
	}
	close(jobs)

	for r := 0; r < 100; r++ {
		result := <-results

		fmt.Println("BI", result.BestIndividual)
		fmt.Println("BS", result.BestScore)

		if result.BestScore < bestOfAll.BestScore {
			bestOfAll = result
		}
	}
	close(results)

	fmt.Println("--------")
	fmt.Println("BOA - BI", bestOfAll.BestIndividual)
	fmt.Println("BOA - BS", bestOfAll.BestScore)
}

func worker(jobs <-chan distances.DistanceMatrix, results chan<- algo.TSPResult) {
	for n := range jobs {
		results <- algo.GeneticAlgorithm(n, 250, 0.05, 100)
	}
}
