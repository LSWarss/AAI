package algo

// Makes Tournament selection on given population with fitness and returns the population after the selection.
func MakeTournament(populationWithFitness PopulationWithFitness) [][][]int {
	var tournament [][][]int
	for i := 0; i < populationWithFitness.PopulationSize; i++ {
		tournament = append(tournament, tournamentSelection(populationWithFitness, 3))
	}

	return tournament
}

func tournamentSelection(populationWithFitness PopulationWithFitness, selectivePressure int) [][]int {
	bestCharacter := [][]int{{0}, {0}}
	selectSlice := createSelectionArray(populationWithFitness.PopulationSize)
	shuffle(selectSlice)

	for i := 0; i < selectivePressure; i++ {
		character := populationWithFitness.ScoredPopulation[selectSlice[i]]
		score := character[1][0]
		if bestCharacter[1][0] == 0 || score < bestCharacter[1][0] {
			bestCharacter = character
		}
	}

	return bestCharacter
}
