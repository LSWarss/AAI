package algo

import (
	"math/rand"
	"reflect"
	"time"
)

func randomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

// Creates selection array with length of given len
// and fills it with number from 0 to the lenght.
func createSelectionArray(len int) []int {
	selection := make([]int, len)
	for i := 0; i < len; i++ {
		selection[i] = i
	}
	return selection
}

// Shuffles arrays of any type
func shuffle(slice interface{}) {
	rv := reflect.ValueOf(slice)
	swap := reflect.Swapper(slice)
	length := rv.Len()
	for i := length - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		swap(i, j)
	}
}
