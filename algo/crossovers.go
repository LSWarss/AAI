package algo

import (
	"math/rand"
)

// Single Point Crossover algorithm.
// Takes two arrays of individuals and then splits both in one place, after this
// switches placement of the splited parts in both arrays.
func singlePointCrossover(individualA, individualB []int) (newIndividualA, newIndividualB []int) {
	x := RandomInt(0, len(individualB))
	newIndividualA, newIndividualB = individualA[:x], individualB[:x]
	newIndividualA = append(newIndividualA, individualB[x:]...)
	newIndividualB = append(newIndividualB, individualA[x:]...)
	return newIndividualA, newIndividualB
}

func MakeSinglePointCrossover(population [][]int) (crossedOverPop [][]int) {
	for j := 0; j < len(population)-1; j++ {
		newIndividualA, newIndividualB := singlePointCrossover(population[j], population[j+1])
		crossedOverPop = append(crossedOverPop, newIndividualA)
		crossedOverPop = append(crossedOverPop, newIndividualB)
	}

	return crossedOverPop
}

// PMX performs partially mapped crossover. PMX inherits a random slice of one
// parent. The position of the other values is more random when there is greater
// difference between the parents.
func pmx(individualA, individualB []int) (individualC []int) {
	individualC = make([]int, len(individualB))
	if rand.Float64() < 0.5 {
		individualA, individualB = individualB, individualA
	}
	_, left, right := RandSlice(individualA)

	for i := range individualC {
		individualC[i] = -1
	}
	copy(individualC[left:right], individualA[left:right])

	for i := left; i < right; i++ {
		if Search(individualC, individualB[i]) == -1 {
			j := i
			for left <= j && j < right {
				j = Search(individualB, individualA[j])
			}
			individualC[j] = individualB[i]
		}
	}

	for i := range individualC {
		if individualC[i] == -1 {
			individualC[i] = individualB[i]
		}
	}

	return individualC
}

func MakePMXCrossover(population [][]int) (crossedOverPop [][]int) {
	for j := 0; j < len(population)-1; j++ {
		firstChild := pmx(population[j], population[j+1])
		secondChild := pmx(population[j+1], population[j])
		crossedOverPop = append(crossedOverPop, firstChild)
		crossedOverPop = append(crossedOverPop, secondChild)
	}

	return crossedOverPop
}
