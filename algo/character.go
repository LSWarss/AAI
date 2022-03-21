package algo

import (
	"math/rand"
	"time"

	distances "github.com/lswarss/AAI/files"
)

type CharactersMatrix struct {
	CharactersCount int
	Characters      [][]int
}

func NewCharactersMatrix(distanceMatrix distances.DistanceMatrix) CharactersMatrix {
	var characters [][]int

	for i := 0; i < distanceMatrix.Rows; i++ {
		var tempSlice []int
		for j := 0; j < distanceMatrix.Rows; j++ {
			tempSlice = append(tempSlice, randIndex(0, distanceMatrix.Rows))
		}

		characters = append(characters, tempSlice)
		tempSlice = nil
	}

	return CharactersMatrix{
		CharactersCount: distanceMatrix.Rows,
		Characters:      characters,
	}
}

func randIndex(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}
