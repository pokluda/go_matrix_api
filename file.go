package main

import (
	"encoding/csv"
	"net/http"
	"strconv"
)

func loadMatrixFromCsvFile(request *http.Request) ([][]string, error) {
	file, _, err := request.FormFile("file")

	if err != nil {
		return nil, err
	}

	defer file.Close()
	strMatrix, err := csv.NewReader(file).ReadAll()

	if err != nil {
		return nil, err
	}

	return strMatrix, nil
}

func convertMatrix(stringMatrix [][]string) (Matrix, error) {
	intMatrix := make([][]int, len(stringMatrix))
	var err error = nil

	for rowIdx, rows := range stringMatrix {
		for _, strVal := range rows {
			intVal, err := strconv.Atoi(strVal)

			if err != nil {
				return intMatrix, err
			}

			intMatrix[rowIdx] = append(intMatrix[rowIdx], intVal)
		}
	}

	return intMatrix, err
}
