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

	characters := algo.NewCharactersMatrix(matrixes[2])
	// log.Println(matrixes[1])
	log.Println(characters)

	scores := algo.GetScore(matrixes[2], characters)
	log.Println(scores)

	bestScore := algo.GetBestCharacter(characters, scores)
	log.Println("BEST SCORE:", bestScore)
}
