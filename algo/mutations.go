package algo

func Mutate(population [][]int, mutationRate float64) {

	for _, individual := range population {
		if shouldMutate() {
			individual = inversionMutation(individual, mutationRate)
		} else {
			transposonMutation(individual)
		}
	}
}

// Inversion Mutation algorithm.
// Takes in array of individuals from population,
// calculates splits in the individual array by len of it and mutationRange as end range.
func inversionMutation(individual []int, mutationRate float64) (newIndividual []int) {
	randRange := int(float64(len(individual)) * mutationRate)
	x := RandomInt(0, randRange)
	y := RandomInt(x, randRange)

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

func shouldMutate() bool {
	pick := RandomFloat(0, 1)

	return pick < 0.66
}
