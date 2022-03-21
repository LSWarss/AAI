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

type Scores []int

func GetScore(distanceMatrix distances.DistanceMatrix, characterMatrix CharactersMatrix) Scores {
	var scores Scores

	for i := 0; i < characterMatrix.CharactersCount; i++ {
		var tempSum int
		for j := 0; j < characterMatrix.CharactersCount; j++ {
			characterIndex := characterMatrix.Characters[i][j]
			tempSum += distanceMatrix.Matrix[i][characterIndex]
		}

		scores = append(scores, tempSum)
		tempSum = 0
	}

	return scores
}

func randIndex(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
