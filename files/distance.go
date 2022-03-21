package files

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type DistanceMatrix struct {
	Rows   int
	Matrix [][]int
}

func newDistanceMatrix(matrixFile io.Reader) (DistanceMatrix, error) {
	scanner := bufio.NewScanner(matrixFile)
	rowNumber := readRowNumber(scanner)
	var slices [][]int

	if rowNumber == 0 {
		return DistanceMatrix{}, fmt.Errorf("No rows data")
	}

	for i := 0; i < rowNumber; i++ {
		scanner.Scan()                             // scan line on index i
		text := scanner.Text()                     // assign the text to variable
		numbersStrings := strings.Split(text, " ") // split text in to slice of stringss
		var numbers []int
		for _, number := range numbersStrings { // go through slice of strings
			number, _ := strconv.Atoi(number) // convert numberString to int
			numbers = append(numbers, number) // append it to the clean slice of ints
		}
		slices = append(slices, numbers) // append numbers slice to slice of slices
	}

	var reverted []int
	for i := 0; i < rowNumber; i++ {
		for j := 1 + i; j < rowNumber; j++ {
			reverted = append(reverted, slices[j][i])
		}
		slices[i] = append(slices[i], reverted...)
		reverted = nil
	}

	return DistanceMatrix{
		Rows:   rowNumber,
		Matrix: slices,
	}, nil
}

func readRowNumber(scanner *bufio.Scanner) int {
	scanner.Scan() // we scan first line of the file looking for rows count
	rows, err := strconv.Atoi(scanner.Text())

	if err != nil {
		fmt.Println(err)
		return 0
	}
	return rows
}
