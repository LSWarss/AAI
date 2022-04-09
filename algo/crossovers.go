package algo

// Single Point Crossover algorithm.
// Takes two arrays of individuals and then splits both in one place, after this
// switches placement of the splited parts in both arrays.
func SinglePointCrossover(individualA, individualB []int) (newIndividualA, newIndividualB []int) {
	x := randomInt(0, len(individualB))
	newIndividualA, newIndividualB = individualA[:x], individualB[:x]
	newIndividualA = append(newIndividualA, individualB[x:]...)
	newIndividualB = append(newIndividualB, individualA[x:]...)
	return newIndividualA, newIndividualB
}
