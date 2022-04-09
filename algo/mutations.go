package algo

// Inversion Mutation algorithm.
// Takes in array of individuals from population,
// calculates splits in the individual array by len of it and mutationRange as end range.
func inversionMutation(individual []int, mutationRate float64) (newIndividual []int) {
	randRange := int(float64(len(individual)) * mutationRate)
	x := randomInt(0, randRange)
	y := randomInt(x, randRange)

	newIndividual = append(newIndividual, individual[y:]...)
	newIndividual = append(newIndividual, individual[x:y]...)
	newIndividual = append(newIndividual, individual[:x]...)

	return newIndividual
}

func MakeInversionMutation(population [][]int, mutationRate float64) (mutatedPopulation [][]int) {
	for _, individual := range population {
		mutatedPopulation = append(mutatedPopulation, inversionMutation(individual, mutationRate))
	}

	return mutatedPopulation
}
