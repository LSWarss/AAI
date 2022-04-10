package algo

import (
	"math/rand"
	"reflect"
	"time"
)

func RandomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func RandomFloat(min, max float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return (rand.Float64() * (max - min)) + min
}

// Creates selection array with length of given len
// and fills it with number from 0 to the lenght.
func CreateSelectionArray(len int) []int {
	selection := make([]int, len)
	for i := 0; i < len; i++ {
		selection[i] = i
	}
	return selection
}

// Shuffles arrays of any type
func Shuffle(slice interface{}) {
	rv := reflect.ValueOf(slice)
	swap := reflect.Swapper(slice)
	length := rv.Len()
	for i := length - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		swap(i, j)
	}
}

// RandSlice returns a random slice of the argument along with the boundaries.
// That is to say:
//     sub == slice[left:right]
func RandSlice(slice []int) (sub []int, left, right int) {
	left = rand.Intn(len(slice))
	right = left
	for right == left {
		right = rand.Intn(len(slice))
	}
	if right < left {
		left, right = right, left
	}
	return slice[left:right], left, right
}

// Search searches an int slice for a particular value and returns the index.
// If the value is not found, Search returns -1.
func Search(slice []int, val int) (idx int) {
	for idx = range slice {
		if slice[idx] == val {
			return idx
		}
	}
	return -1
}
