package main

import (
	"log"
	"os"

	alog "github.com/lswarss/AAI/algo"
	distances "github.com/lswarss/AAI/files"
)

func main() {
	matrixes, err := distances.NewDistanceMatrixFromFS(os.DirFS("data"))
	if err != nil {
		log.Fatal(err)
	}

	characters := alog.NewCharactersMatrix(matrixes[2])
	log.Println(matrixes)
	log.Println(characters)
}
