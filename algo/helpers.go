package algo

import (
	"math/rand"
	"reflect"
	"time"
)

func randIndex(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func selection(count int) []int {
	selection := make([]int, count)
	for i := 0; i < count; i++ {
		selection[i] = i
	}
	return selection
}

func shuffle(slice interface{}) {
	rv := reflect.ValueOf(slice)
	swap := reflect.Swapper(slice)
	length := rv.Len()
	for i := length - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		swap(i, j)
	}
}
