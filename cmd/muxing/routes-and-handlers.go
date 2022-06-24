package main

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func sumAandB(w http.ResponseWriter, r *http.Request) {
	a, _ := strconv.Atoi(r.Header["A"][0])
	b, _ := strconv.Atoi(r.Header["B"][0])
	c := strconv.Itoa(a + b)
	w.Header()["a+b"] = []string{c}
	w.WriteHeader(http.StatusOK)
}

func bodyMessage(w http.ResponseWriter, r *http.Request) {
	a, _ := strconv.Atoi(r.Header["A"][0])
	b, _ := strconv.Atoi(r.Header["B"][0])
	c := strconv.Itoa(a + b)
	w.Header()["a+b"] = []string{c}
	w.WriteHeader(http.StatusOK)
}

func addRoutes(router *mux.Router) {
	router.HandleFunc("/headers", sumAandB).Methods("POST").
		Headers("a", "", "b", "")
}
