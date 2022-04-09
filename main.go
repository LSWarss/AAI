package main

import (
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

	log.Println(matrixes[2])

	charsWithScores := algo.CreatePopulationWithFitnessMatrix(matrixes[2])
	log.Println("Characters with scores:", charsWithScores.ScoredPopulation)

	tournament := algo.MakeTournament(charsWithScores)
	log.Println("Tournament", tournament)

	var characters [][]int
	for i := 0; i < len(tournament); i++ {
		characters = append(characters, tournament[i][0])
	}

	var crossover [][]int
	for j := 0; j < len(characters)-1; j++ {
		new_A, new_B := algo.SinglePointCrossover(characters[j], characters[j+1])
		crossover = append(crossover, new_A)
		crossover = append(crossover, new_B)
	}

	log.Println("After crossover", crossover)

	for i, v := range crossover {
		crossover[i] = algo.InversionMutation(v, 0.3)
	}

	log.Println("After inversion", crossover)

	scores := algo.CountFitness(matrixes[2], crossover)

	log.Println("Scores", scores)
}
