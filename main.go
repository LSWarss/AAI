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

	charsWithScores := algo.GetCharactersWithScoresMatrix(matrixes[2])
	log.Println("Characters with scores:", charsWithScores.CharactersAndScores)

	tournament := algo.GetTournament(charsWithScores)
	log.Println("Tournament", tournament)

	var characters [][]int
	for i := 0; i < len(tournament); i++ {
		characters = append(characters, tournament[i][0])
	}

	new_A, new_B := algo.SinglePointCrossover(characters[0], characters[1])
	log.Println(new_A, new_B)
}
