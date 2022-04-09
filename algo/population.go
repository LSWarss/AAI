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

var (
	erasCount = 200
)

func CreatePopulationWithFitnessMatrix(distanceMatrix distances.DistanceMatrix) PopulationWithFitness {
	population := CreatePopulationMatrix(distanceMatrix, erasCount)
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
func CreatePopulationMatrix(distanceMatrix distances.DistanceMatrix, erasCount int) PopulationMatrix {
	rand.Seed(time.Now().UnixNano())
	var population [][]int

	for i := 0; i < erasCount; i++ {
		tempIndividual := createSelectionArray(distanceMatrix.Rows)

		rand.Shuffle(len(tempIndividual), func(i, j int) {
			tempIndividual[i], tempIndividual[j] = tempIndividual[j], tempIndividual[i]
		})

		population = append(population, tempIndividual)
		tempIndividual = nil
	}

	return PopulationMatrix{
		PopulationSize: erasCount,
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
