package algo

func Mutate(population [][]int, mutationRate float64) {

	for i, individual := range population {
		if shouldMutate(mutationRate) {
			population[i] = inversionMutation(individual)
		}
	}
}

// Inversion Mutation algorithm.
// Takes in array of individuals from population,
// calculates splits in the individual array by len of it and mutationRange as end range.
func inversionMutation(individual []int) (newIndividual []int) {
	x := RandomInt(0, len(individual))
	y := RandomInt(x, len(individual))

	newIndividual = append(newIndividual, individual[y:]...)
	newIndividual = append(newIndividual, individual[x:y]...)
	newIndividual = append(newIndividual, individual[:x]...)

	return newIndividual
}

func getSwapPoints(lenght int) (first, second int) {
	first = RandomInt(0, lenght)
	for first != second {
		second = RandomInt(0, lenght)
	}

	return first, second
}

func swapGenes(genotype []int, firstSwap, secondSwap int) {
	temp := genotype[firstSwap]
	genotype[firstSwap] = genotype[secondSwap]
	genotype[secondSwap] = temp
}

func transposonMutation(individual []int) {
	first, second := getSwapPoints(len(individual))
	swapGenes(individual, first, second)
}

func shouldMutate(mutationRate float64) bool {
	pick := RandomFloat(0, 1)

	return pick < mutationRate
}
