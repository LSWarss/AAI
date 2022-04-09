package algo

// Inversion Mutation algorithm.
// Takes in array of individuals from population,
// calculates splits in the individual array by len of it and mutationRange as end range.
func InversionMutation(individual []int, mutationRate float32) (newIndividual []int) {
	randRange := int(float32(len(individual)) * mutationRate)
	x := randomInt(0, randRange)
	y := randomInt(x, randRange)

	newIndividual = append(newIndividual, individual[y:]...)
	newIndividual = append(newIndividual, individual[x:y]...)
	newIndividual = append(newIndividual, individual[:x]...)

	return newIndividual
}
