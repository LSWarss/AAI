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

	characters := algo.NewCharactersMatrix(matrixes[1])
	// log.Println(matrixes[1])
	log.Println(characters)

	scores := algo.GetScore(matrixes[1], characters)
	log.Println(scores)
}
