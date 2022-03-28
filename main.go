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

	charsWithScores := algo.GetCharactersWithScoresMatrix(matrixes[1])
	log.Println("Characters with scores:", charsWithScores.CharactersAndScores)

	tournament := algo.GetTournament(charsWithScores)
	log.Println("Tournament", tournament)
}
