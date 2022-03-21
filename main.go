package main

import (
	"log"
	"os"

	distances "github.com/lswarss/AAI/files"
)

func main() {
	matrix, err := distances.NewDistanceMatrixFromFS(os.DirFS("data"))
	if err != nil {
		log.Fatal(err)
	}

	log.Println(matrix)
}
