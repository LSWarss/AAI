package files

import "io/fs"

func NewDistanceMatrixFromFS(filesystem fs.FS) ([]DistanceMatrix, error) {
	dir, err := fs.ReadDir(filesystem, ".")
	if err != nil {
		return nil, err
	}

	var matrixes []DistanceMatrix
	for _, f := range dir {
		matrix, err := getMatrix(filesystem, f)
		if err != nil {
			return nil, err
		}
		matrixes = append(matrixes, matrix)
	}
	return matrixes, nil
}

func getMatrix(fileSystem fs.FS, f fs.DirEntry) (DistanceMatrix, error) {
	matrixFile, err := fileSystem.Open(f.Name())
	if err != nil {
		return DistanceMatrix{}, err
	}
	defer matrixFile.Close()

	return newDistanceMatrix(matrixFile)
}
