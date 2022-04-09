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
	bestIndividual := [][]int{{0}, {0}}
	selectSlice := CreateSelectionArray(populationWithFitness.PopulationSize)
	Shuffle(selectSlice)

	for i := 0; i < selectivePressure; i++ {
		individual := populationWithFitness.ScoredPopulation[selectSlice[i]]
		score := individual[1][0]
		if bestIndividual[1][0] == 0 || score < bestIndividual[1][0] {
			bestIndividual = individual
		}
	}

	return bestIndividual
}
