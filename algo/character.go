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

type CharactersWithScoresMatrix struct {
	CharactersCount     int
	CharactersAndScores [][][]int
	BestCharacter       [][]int
}

func GetCharactersWithScoresMatrix(distanceMatrix distances.DistanceMatrix) CharactersWithScoresMatrix {
	characters := getCharactersMatrix(distanceMatrix)
	scores := getScore(distanceMatrix, characters)

	var returnArray [][][]int

	for i := 0; i < len(scores)-1; i++ {
		returnArray = append(returnArray, [][]int{characters.Characters[i], {scores[i]}})
	}

	return CharactersWithScoresMatrix{
		CharactersCount:     len(returnArray),
		CharactersAndScores: returnArray,
	}
}

func GetTournament(charWithScore CharactersWithScoresMatrix) [][][]int {
	var tournament [][][]int
	for i := 0; i < charWithScore.CharactersCount; i++ {
		tournament = append(tournament, tournamentSelection(charWithScore, 3))
	}

	return tournament
}

func getCharactersMatrix(distanceMatrix distances.DistanceMatrix) CharactersMatrix {
	rand.Seed(time.Now().UnixNano())
	var characters [][]int

	for i := 0; i < distanceMatrix.Rows; i++ {
		var tempSlice []int
		for j := 0; j < distanceMatrix.Rows; j++ {
			tempSlice = append(tempSlice, j)
		}

		rand.Shuffle(len(tempSlice), func(i, j int) {
			tempSlice[i], tempSlice[j] = tempSlice[j], tempSlice[i]
		})

		characters = append(characters, tempSlice)
		tempSlice = nil
	}

	return CharactersMatrix{
		CharactersCount: distanceMatrix.Rows,
		Characters:      characters,
	}
}

func getScore(distanceMatrix distances.DistanceMatrix, characterMatrix CharactersMatrix) []int {
	var scores []int

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

func getBestCharacter(charMatrix CharactersMatrix, scores []int) [][]int {
	bestScore := scores[0]
	var bestScoreIndex int

	for i := 0; i < len(scores)-1; i++ {
		if scores[i] < bestScore {
			bestScoreIndex = i
			bestScore = scores[i]
		}
	}

	return [][]int{
		charMatrix.Characters[bestScoreIndex],
		{bestScore},
	}
}

func tournamentSelection(charWithScore CharactersWithScoresMatrix, selectivePressure int) [][]int {
	bestCharacter := [][]int{{0}, {0}}
	selectSlice := selection(charWithScore.CharactersCount)
	shuffle(selectSlice)

	for i := 0; i < selectivePressure; i++ {
		character := charWithScore.CharactersAndScores[selectSlice[i]]
		score := character[1][0]
		if bestCharacter[1][0] == 0 || score < bestCharacter[1][0] {
			bestCharacter = character
		}
	}

	return bestCharacter
}

func SinglePointCrossover(A, B []int) (A_new, B_new []int) {
	x := randIndex(0, len(B))
	A_new, B_new = A[:x], B[:x]
	A_new = append(A_new, B[x:]...)
	B_new = append(B_new, A[x:]...)
	return A_new, B_new
}
