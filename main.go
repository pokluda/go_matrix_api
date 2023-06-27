package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	http.Handle("/echo", http.HandlerFunc(apiEcho))
	http.Handle("/inverse", http.HandlerFunc(apiInverse))
	http.Handle("/flatten", http.HandlerFunc(apiFlatten))
	http.Handle("/sum", http.HandlerFunc(apiSum))
	http.Handle("/multiply", http.HandlerFunc(apiMultiply))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panic("ListenAndServe:", err)
	}
}

func apiEcho(writer http.ResponseWriter, request *http.Request) {
	m := parseMatrixFromRequest(writer, request)

	printMatrixToResponse(m, writer)
}

func apiInverse(writer http.ResponseWriter, request *http.Request) {
	m := parseMatrixFromRequest(writer, request)
	inversed := m.inverse()

	printMatrixToResponse(inversed, writer)
}

func apiFlatten(writer http.ResponseWriter, request *http.Request) {
	m := parseMatrixFromRequest(writer, request)
	arr := m.flatten()

	strArr := []string{}
	for _, intVal := range arr {
		strArr = append(strArr, strconv.Itoa(intVal))
	}

	fmt.Fprint(writer, strings.Join(strArr, ","))
}

func apiSum(writer http.ResponseWriter, request *http.Request) {
	m := parseMatrixFromRequest(writer, request)

	fmt.Fprint(writer, strconv.Itoa(m.aggregate(sum, 0)))
}

func apiMultiply(writer http.ResponseWriter, request *http.Request) {
	m := parseMatrixFromRequest(writer, request)

	fmt.Fprint(writer, strconv.Itoa(m.aggregate(mult, 1)))
}
