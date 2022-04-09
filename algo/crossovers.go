package algo

// Single Point Crossover algorithm.
// Takes two arrays of individuals and then splits both in one place, after this
// switches placement of the splited parts in both arrays.
func singlePointCrossover(individualA, individualB []int) (newIndividualA, newIndividualB []int) {
	x := randomInt(0, len(individualB))
	newIndividualA, newIndividualB = individualA[:x], individualB[:x]
	newIndividualA = append(newIndividualA, individualB[x:]...)
	newIndividualB = append(newIndividualB, individualA[x:]...)
	return newIndividualA, newIndividualB
}

func MakeCrossover(population [][]int) (crossedOverPop [][]int) {
	for j := 0; j < len(population)-1; j++ {
		newIndividualA, newIndividualB := singlePointCrossover(population[j], population[j+1])
		crossedOverPop = append(crossedOverPop, newIndividualA)
		crossedOverPop = append(crossedOverPop, newIndividualB)
	}

	return crossedOverPop
}
