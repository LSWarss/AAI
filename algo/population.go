package algo

import (
	"math/rand"
	"time"

	distances "github.com/lswarss/AAI/files"
)

type PopulationMatrix struct {
	PopulationSize int
	Population     [][]int
}

type PopulationWithFitness struct {
	PopulationSize   int
	ScoredPopulation [][][]int
	BestIdividual    [][]int
}

func CreatePopulationWithFitnessMatrix(distanceMatrix distances.DistanceMatrix, popSize int) PopulationWithFitness {
	population := CreatePopulationMatrix(distanceMatrix, popSize)
	populationFitness := CountFitness(distanceMatrix, population.Population)

	var returnArray [][][]int

	for i := 0; i < len(populationFitness)-1; i++ {
		returnArray = append(returnArray, [][]int{population.Population[i], {populationFitness[i]}})
	}

	return PopulationWithFitness{
		PopulationSize:   len(returnArray),
		ScoredPopulation: returnArray,
	}
}

// Creates population matrix by going through distanceMatrix
// for given erasCount and randomly picking distances for individual.
func CreatePopulationMatrix(distanceMatrix distances.DistanceMatrix, popSize int) PopulationMatrix {
	rand.Seed(time.Now().UnixNano())
	var population [][]int

	for i := 0; i < popSize; i++ {
		tempIndividual := createSelectionArray(distanceMatrix.Rows)

		rand.Shuffle(len(tempIndividual), func(i, j int) {
			tempIndividual[i], tempIndividual[j] = tempIndividual[j], tempIndividual[i]
		})

		population = append(population, tempIndividual)
		tempIndividual = nil
	}

	return PopulationMatrix{
		PopulationSize: popSize,
		Population:     population,
	}
}

// Counts fitness for given population agains given Distance Matrix
func CountFitness(distanceMatrix distances.DistanceMatrix, population [][]int) (fitness []int) {
	for _, individual := range population {
		lastIndex := 0
		var tempSum int

		for i := 0; i < distanceMatrix.Rows; i++ {
			distance := individual[i]
			tempSum += distanceMatrix.Matrix[lastIndex][distance]
			lastIndex = distance
		}
		tempSum += distanceMatrix.Matrix[lastIndex][individual[0]]

		fitness = append(fitness, tempSum)
		lastIndex = 0
		tempSum = 0
	}

	return fitness
}

func getBestIndividualIndex(fitness []int) int {
	bestFitness := fitness[0]
	var bestScoreIndex int

	for i := 0; i < len(fitness)-1; i++ {
		if fitness[i] < bestFitness {
			bestScoreIndex = i
			bestFitness = fitness[i]
		}
	}

	return bestScoreIndex
}

func GetBestFitnessAndIndividual(distanceMatrix distances.DistanceMatrix, population [][]int) (bestIndividual []int, bestFitness int) {
	fitness := CountFitness(distanceMatrix, population)
	bestIndex := getBestIndividualIndex(fitness)
	bestIndividual = population[bestIndex]
	bestFitness = fitness[bestIndex]

	return bestIndividual, bestFitness
}
