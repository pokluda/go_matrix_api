package main

import (
	"fmt"
	"net/http"
	"strings"
)

func parseMatrixFromRequest(writer http.ResponseWriter, request *http.Request) Matrix {
	strMatrix, err := loadMatrixFromCsvFile(request)

	if err != nil {
		handleError("Failed to parsed matrix from a CSV file.", err, writer)
		return nil
	}

	m, err := convertMatrix(strMatrix)

	if err != nil {
		handleError("Failed to convert matrix from strings to integers.", err, writer)
		return nil
	}

	err = m.validate()

	if err != nil {
		handleError("Invalid matrix.", err, writer)
		return nil
	}

	return m
}

func printMatrixToResponse(m Matrix, writer http.ResponseWriter) {
	var response string

	for _, rows := range m {
		response = fmt.Sprintf("%s%s\n", response,
			strings.Trim(strings.Join(strings.Fields(fmt.Sprint(rows)), ","), "[]"))
	}

	fmt.Fprint(writer, response)
}

func handleError(errorMessage string, err error, writer http.ResponseWriter) {
	writer.Write([]byte(fmt.Sprintf("%s %s", errorMessage, err.Error())))
}
