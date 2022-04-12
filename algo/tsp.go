package algo

import (
	"fmt"

	distances "github.com/lswarss/AAI/files"
)

func createPopulationWithFitness(population [][]int, fitness []int) (scoredPopulation PopulationWithFitness) {
	for i := 0; i < len(fitness); i++ {
		scoredPopulation.ScoredPopulation = append(scoredPopulation.ScoredPopulation, [][]int{population[i], {fitness[i]}})
	}
	scoredPopulation.PopulationSize = len(scoredPopulation.ScoredPopulation)

	return scoredPopulation
}

func getPopulationFromSelection(selection [][][]int) (population [][]int) {
	for i := 0; i < len(selection); i++ {
		population = append(population, selection[i][0])
	}

	return population
}

func nextGeneration(distancesMatrix distances.DistanceMatrix, population [][]int, mutationRate float64) [][]int {
	populationFitness := CountFitness(distancesMatrix, population)
	scoredPopulation := createPopulationWithFitness(population, populationFitness)
	selection := MakeTournament(scoredPopulation)
	populationFromSelection := getPopulationFromSelection(selection)
	pick := RandomFloat(0, 1)

	if pick < 0.75 {
		populationFromSelection = MakePMXCrossover(populationFromSelection)
	}

	Mutate(populationFromSelection, mutationRate)

	return populationFromSelection
}

type TSPResult struct {
	BestIndividual []int
	BestScore      int
}

func GeneticAlgorithm(distancesMatrix distances.DistanceMatrix, populationSize int, mutationRate float64, generations int) TSPResult {
	initialPop := CreatePopulationMatrix(distancesMatrix, populationSize)

	for i := 0; i <= generations; i++ {
		initialPop.Population = nextGeneration(distancesMatrix, initialPop.Population, mutationRate)
		bestIndividual, bestScore := GetBestFitnessAndIndividual(distancesMatrix, initialPop.Population)
		fmt.Println("G: ", i)
		fmt.Println("BI: ", bestIndividual)
		fmt.Println("BS: ", bestScore)
	}

	bestIndividual, bestScore := GetBestFitnessAndIndividual(distancesMatrix, initialPop.Population)

	return TSPResult{
		BestIndividual: bestIndividual,
		BestScore:      bestScore,
	}
}
