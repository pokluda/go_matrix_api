package main

import (
	"errors"
)

type Matrix [][]int
type AggFn func(Matrix, int) int

var matrixErrorEmpty = errors.New("There's not much to do with an empty matrix. Please, give me some proper input.")
var matrixErrorNotSquare = errors.New("This is a boring matrix API therefore the input matrix must be a square matrix.")

func (m Matrix) inverse() Matrix {
	n := len(m)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			m[i][j], m[j][i] = m[j][i], m[i][j]
		}
	}

	return m
}

func (m Matrix) flatten() []int {
	arr := []int{}

	for _, row := range m {
		arr = append(arr, row...)
	}

	return arr
}

func (m Matrix) aggregate(fn func(int, int) int, initVal int) int {
	aggVal := initVal

	for _, rows := range m {
		for _, intVal := range rows {
			aggVal = fn(aggVal, intVal)
		}
	}

	return aggVal
}

func sum(a, b int) int {
	return a + b
}

func mult(a, b int) int {
	return a * b
}

func (m Matrix) validate() error {
	if len(m) == 0 {
		return matrixErrorEmpty
	}

	if !m.IsSquare() {
		return matrixErrorNotSquare
	}

	return nil
}

func (m Matrix) IsSquare() bool {
	rowsLen := len(m)

	for _, rows := range m {
		colsLen := len(rows)

		if rowsLen != colsLen {
			return false
		}
	}

	return true
}
