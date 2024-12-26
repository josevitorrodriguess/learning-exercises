package main

import "fmt"

func main() {
	x := [][]int{{1, 2}, {3, 4}}
	y := [][]int{{5, 6}, {7, 8}}

	z := sumMatrixes(x, y)

	// z será: {{6, 8}, {10, 12}}
	fmt.Println(z)
}

func sumMatrixes(x [][]int, y [][]int) [][]int {
	z := make([][]int, len(x))
	for i := range z {
		z[i] = make([]int, len(x[i]))
	}

	for i := range x {
		for j := range x[i] {
			z[i][j] = x[i][j] + y[i][j]
		}
	}
	return z
}

func fillMatrix(vec []int, n, m int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			matrix[i][j] = vec[i*m+j]
		}
	}

	return matrix
}

func makeMatrix(size int) [][]int {
	matrix := make([][]int, size)
	for i := range matrix {
		matrix[i] = make([]int, size)
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if j == size-i-1 {
				matrix[i][j] = 1
			} else {
				matrix[i][j] = 0
			}
		}
	}

	return matrix
}
