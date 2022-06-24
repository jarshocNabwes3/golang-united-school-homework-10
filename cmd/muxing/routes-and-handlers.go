package main

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

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
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	body = append([]byte(`I got message:\n`), body...)
	w.Write(body)
}

func badRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func requestPath(w http.ResponseWriter, r *http.Request) {
	param := strings.Split(r.URL.Path, "/")[2]
	// vars := mux.Vars(r)
	body := "Hello, " + param + "!"
	w.Write([]byte(body))
}

func addRoutes(router *mux.Router) {
	router.HandleFunc("/headers", sumAandB).Methods("POST").
		Headers("a", "", "b", "")
	router.HandleFunc("/data", sumAandB).Methods("POST")
	router.HandleFunc("/bad", badRequest).Methods("GET")
	router.HandleFunc("/name/{param}", requestPath).Methods("GET")
}
