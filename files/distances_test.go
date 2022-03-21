package files_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	distances "github.com/lswarss/AAI/files"
)

type StubFailingFs struct{}

var ErrAlwaysFail = errors.New("oh no, I always fail")

func (s StubFailingFs) Open(name string) (fs.File, error) {
	return nil, ErrAlwaysFail
}

func TestDistances(t *testing.T) {
	const (
		test = `5
0
1 0
4 2 0
7 2 5 0
9 1 7 8 0
`
		test1 = `10
0
1 0
4 2 0
7 2 5 0
9 1 7 8 0
9 1 7 8 1 0
9 1 7 8 4 2 0
9 1 7 8 7 2 5 0
9 1 7 8 9 1 7 8 0
9 1 7 8 9 1 7 8 9 0
`
	)
	t.Run("errors checking", func(t *testing.T) {
		_, err := distances.NewDistanceMatrixFromFS(StubFailingFs{})

		assertError(t, err, ErrAlwaysFail)
	})

	t.Run("load matrix", func(t *testing.T) {
		fs := fstest.MapFS{
			"test.txt":   {Data: []byte(test)},
			"text2.text": {Data: []byte(test1)},
		}

		matrixes, err := distances.NewDistanceMatrixFromFS(fs)

		assertNoError(t, err)
		assertMatrix(t, matrixes[0], distances.DistanceMatrix{
			Rows: 5,
			Matrix: [][]int{{0, 1, 4, 7, 9},
				{1, 0, 2, 2, 1},
				{4, 2, 0, 5, 7},
				{7, 2, 5, 0, 8},
				{9, 1, 7, 8, 0}},
		})
		assertMatrix(t, matrixes[1], distances.DistanceMatrix{
			Rows: 10,
			Matrix: [][]int{{0, 1, 4, 7, 9, 9, 9, 9, 9, 9},
				{1, 0, 2, 2, 1, 1, 1, 1, 1, 1},
				{4, 2, 0, 5, 7, 7, 7, 7, 7, 7},
				{7, 2, 5, 0, 8, 8, 8, 8, 8, 8},
				{9, 1, 7, 8, 0, 1, 4, 7, 9, 9},
				{9, 1, 7, 8, 1, 0, 2, 2, 1, 1},
				{9, 1, 7, 8, 4, 2, 0, 5, 7, 7},
				{9, 1, 7, 8, 7, 2, 5, 0, 8, 8},
				{9, 1, 7, 8, 9, 1, 7, 8, 0, 9},
				{9, 1, 7, 8, 9, 1, 7, 8, 9, 0}},
		})
	})
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("We didn't want an error but we did get one")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("We wanted an error but didn't get one")
	}

	if got != want {
		t.Errorf("got %s but wanted %s", got, want)
	}
}

func assertMatrix(t *testing.T, got distances.DistanceMatrix, want distances.DistanceMatrix) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
